package worker

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/crxfoz/outbox"
)

//go:generate mockgen -destination=./mocks_test.go -package=worker -source=./worker.go -typed true

type Worker[ID any, OE Record[ID], RE any, TX Transaction] struct {
	cfg         Config
	handler     Handler[RE]
	mapper      EventMapper[ID, OE, RE]
	repo        Repository[ID, OE, TX]
	logger      outbox.Logger
	ctxProvider func(context.Context) context.Context
	// TODO: add callbacks for errors and other key steps to make it possible
	//  for users to build metrics for outbox for example
}

func NewWorker[ID any, OE Record[ID], RE any, TX Transaction](
	cfg Config,
	handler Handler[RE],
	mapper EventMapper[ID, OE, RE],
	repo Repository[ID, OE, TX], options ...Option) *Worker[ID, OE, RE, TX] {

	worker := &Worker[ID, OE, RE, TX]{
		cfg:     cfg,
		handler: handler,
		mapper:  mapper,
		repo:    repo,
	}

	defaultOptions := Options{}

	for _, opt := range options {
		opt(&defaultOptions)
	}

	// not nice, but taking *Worker in Option type requires to define the same generics which is even worse
	worker.logger = defaultOptions.logger
	worker.ctxProvider = defaultOptions.ctxProvider

	return worker
}

// Record is an interface for the record that is stored in DB in the outbox table
type Record[T any] interface {
	// GetID returns the primary key (usually int or uuid) of record stored in the Outbox table
	GetID() T
}

// Transaction represents a transaction type in DB
// It can be sql.Tx, sqlx.Tx, gorm.Tx or something custom for your application
type Transaction interface {
	Commit() error
	Rollback() error
}

// Repository is an abstraction to interact with Outbox table in DB
type Repository[ID any, O Record[ID], TX Transaction] interface {
	// GetTx returns DB's Transaction in order to receive events with Lock,
	// process them, and update the status of it in one transaction.
	GetTx(ctx context.Context) (TX, error)
	// RecordsByStatus returns the amount of awaiting events in Outbox table.
	// Awaiting events is the events with status = outbox.StatusPrepared
	RecordsByStatus(ctx context.Context, status outbox.Status) (uint, error)
	// RecordWithLock returns records from Outbox table.
	//
	// It is a user's responsibility to use FOR UPDATE inside the implementation.
	// A user should implement: LIMIT using limit value, if skipLocked is true SKIP LOCKED
	// suffix, should be added to FOR UPDATE and filter records by status.
	// Note that it is recommended to use ORDER BY id asc (or timestamp) in implementation
	// if processing order is important for your application.
	RecordWithLock(ctx context.Context, tx TX, limit uint, skipLocked bool, status outbox.Status) ([]O, error)
	// UpdateStatus updates status of record with id
	UpdateStatus(ctx context.Context, tx TX, id ID, status outbox.Status) error
}

// EventMapper is an interface that allows user to cast O (Entity from Outbox table) to R (Entity expected by Handler)
type EventMapper[ID any, O Record[ID], R any] interface {
	// ToEvent casts O to R
	// It may return outbox.RetryableError (depends on User's implementation).
	// If in returns an error, and it's not outbox.RetryableError then outbox's entity will be marked as failed
	// and won't be retried.
	// If it returns outbox.RetryableError then the entity will be retried until it succeeds.
	//
	// Note that if skipLocked isn't used, and it returns outbox.RetryableError it might lead
	// to congestion of awaiting events (if service using outbox is running in a single instance it
	// will lead to forever retrying even if skipLocked is used).
	ToEvent(O) (R, error)
}

// Handler is an abstraction that Process event (e.g. Send it to Kafka)
type Handler[R any] interface {
	// Process processes event.
	//
	// It is recommended to return outbox.RetryableError in your Process implementation
	// to retry it instead of just failing the outbox's entity if action is failed
	// (e.g. Publish event to Kafka) in your code
	Process(context.Context, R) error
}

func (w *Worker[ID, OE, RE, TX]) withTransaction(ctx context.Context, fn func(TX) error) error {
	tx, err := w.repo.GetTx(ctx)
	if err != nil {
		return fmt.Errorf("unable to get transaction: %w", err)
	}

	// try to Rollback if callback return an error
	err = fn(tx)
	if err != nil {
		// if Rollback failed -> aggregates both errors: original and reason caused rollback to fail
		if errLoc := tx.Rollback(); errLoc != nil {
			return errors.Join(
				fmt.Errorf("unable to rollback tx: %w", errLoc),
				err,
			)
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("unable to commit: %w", err)
	}

	return nil
}

func (w *Worker[ID, OE, RE, TX]) process(ctx context.Context, tx TX, item OE) error {
	out, err := w.mapper.ToEvent(item)
	if err != nil {
		return fmt.Errorf("unable to cast OE to RE: %w", err)
	}

	err = w.handler.Process(ctx, out)
	if err != nil {
		return fmt.Errorf("unable to process record: %w", err)
	}

	err = w.repo.UpdateStatus(ctx, tx, item.GetID(), outbox.StatusDone)
	if err != nil {
		return fmt.Errorf("unable to process status: %w", err)
	}

	return nil
}

func (w *Worker[ID, OE, RE, TX]) provideCtx(ctx context.Context) context.Context {
	if w.ctxProvider == nil {
		return ctx
	}

	return w.ctxProvider(ctx)
}

func (w *Worker[ID, OE, RE, TX]) start(ctx context.Context) error {
	count, err := w.repo.RecordsByStatus(w.provideCtx(ctx), outbox.StatusPrepared)
	if err != nil {
		return fmt.Errorf("unable to get amount of tasks: %w", err)
	}

	if w.logger != nil {
		w.logger.Info("got count of awaiting events", "count", count)
	}

	if count == 0 {
		return nil
	}

	for {
		err = w.withTransaction(w.provideCtx(ctx), func(tx TX) error {
			items, err := w.repo.RecordWithLock(
				w.provideCtx(ctx),
				tx,
				w.cfg.BatchSize,
				w.cfg.SkipLocked,
				outbox.StatusPrepared,
			)
			if err != nil {
				return fmt.Errorf("unable to get record to be sent: %w", err)
			}

			for i := 0; i < len(items); i++ {
				// Check ctx each iteration in order to interrupt processing of batch.
				// Otherwise, stopping may take significant time that may break graceful-shutdown
				// process (e.g. k8s won't wait forever when killing pod)
				select {
				case <-ctx.Done():
					if w.logger != nil {
						w.logger.Info("batch processing has been interrupted")
					}
					break
				default:
				}

				err = w.process(w.provideCtx(ctx), tx, items[i])
				if err != nil {
					//  TODO: refactor it to reduce nesting
					// mark task as failed if errors isn't a RetryableError
					var errRet outbox.RetryableError
					if !errors.As(err, &errRet) {
						if err := w.repo.UpdateStatus(
							w.provideCtx(ctx),
							tx,
							items[i].GetID(),
							outbox.StatusFailed,
						); err != nil {
							return fmt.Errorf("unable to set status for failed task: %w", err)
						}
					}
				}
			}
			return nil
		})

		if err != nil {
			if errors.Is(err, outbox.ErrNoMoreEntries) {
				if w.logger != nil {
					w.logger.Debug("no more entries to process")
				}
				break
			}

			return err
		}
	}

	return nil
}

func (w *Worker[ID, OE, RE, TX]) Start(ctx context.Context) error {
	if w.logger != nil {
		w.logger.Info("starting worker")
	}

	tk := time.NewTicker(w.cfg.PollingDelay)
	defer tk.Stop()

	for {
		select {
		case <-tk.C:
			if err := w.start(ctx); err != nil && w.logger != nil {
				w.logger.Error("processing failed", "error", err)
			}
		case <-ctx.Done():
			if w.logger != nil {
				w.logger.Info("stopping worker")
			}
			return ctx.Err()
		}
	}
}
