package worker

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/crxfoz/outbox"
)

// TODO: add more test-cases
// TODO: add integration tests

type OutboxRecord struct {
	ID   int
	Data []byte
}

func (or OutboxRecord) GetID() int {
	return or.ID
}

type EventRecord struct {
	Data string
}

func TestWorker_Start(t *testing.T) {
	ctx, cancelFn := context.WithCancel(context.Background())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	items := [...]OutboxRecord{
		{
			ID:   1,
			Data: []byte("id-1"),
		},
		{
			ID:   2,
			Data: []byte("id-2"),
		},
		{
			ID:   3,
			Data: []byte("id-3"),
		},
		{
			ID:   4,
			Data: []byte("id-4"),
		},
		{
			ID:   5,
			Data: []byte("id-5"),
		},
	}

	itemsSlice := items[:]
	doneCh := make(chan struct{}, 1)

	eventMapper := NewMockEventMapper[int, OutboxRecord, EventRecord](ctrl)
	eventMapper.EXPECT().ToEvent(gomock.Any()).DoAndReturn(func(arg OutboxRecord) (EventRecord, error) {
		return EventRecord{
			Data: string(arg.Data),
		}, nil
	}).Times(5)

	tx := NewMockTransaction(ctrl)
	tx.EXPECT().Rollback().DoAndReturn(func() error {
		t.Log("rollback")
		return nil
	}).Times(1)

	tx.EXPECT().Commit().DoAndReturn(func() error {
		t.Log("commit")
		return nil
	}).Times(5)

	repo := NewMockRepository[int, OutboxRecord, *MockTransaction](ctrl)

	repo.EXPECT().GetTx(gomock.Any()).DoAndReturn(func(_ context.Context) (*MockTransaction, error) {
		t.Log("got new tx")
		return tx, nil
	}).MinTimes(1)

	repo.EXPECT().RecordWithLock(
		gomock.Any(),
		gomock.Any(),
		gomock.Eq(uint(1)),
		gomock.Eq(false),
		gomock.Eq(outbox.StatusPrepared),
	).DoAndReturn(func(
		_ context.Context,
		_ *MockTransaction,
		limit uint,
		skipLocked bool,
		status outbox.Status,
	) ([]OutboxRecord, error) {
		if len(itemsSlice) == 0 {
			doneCh <- struct{}{}
			return nil, outbox.ErrNoMoreEntries
		}

		curr := itemsSlice[0]
		itemsSlice = itemsSlice[1:]
		return []OutboxRecord{curr}, nil
	}).
		MinTimes(5).
		MaxTimes(10)

	repo.EXPECT().
		UpdateStatus(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Eq(outbox.StatusDone)).
		DoAndReturn(func(_ context.Context, _ *MockTransaction, id int, status outbox.Status) error {
			t.Log("updating status:", id)
			return nil
		}).Times(5)

	repo.EXPECT().
		RecordsByStatus(gomock.Any(), gomock.Eq(outbox.StatusPrepared)).
		DoAndReturn(func(ctx context.Context, status outbox.Status) (uint, error) {
			x := uint(len(itemsSlice))
			if x == 0 {
				doneCh <- struct{}{}
			}
			return x, nil
		}).AnyTimes()

	handler := NewMockHandler[EventRecord](ctrl)

	processIndx := 0
	handler.EXPECT().
		Process(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, record EventRecord) error {
			assert.Equal(t, string(items[processIndx].Data), record.Data)
			processIndx += 1
			return nil
		}).Times(5)

	cfg := NewConfigOrdered()
	cfg.PollingDelay = time.Second * 5

	worker := NewWorker[int, OutboxRecord, EventRecord, *MockTransaction](cfg, handler, eventMapper, repo)

	go func() {
		err := worker.Start(ctx)
		assert.True(t, errors.Is(err, context.Canceled))
	}()

	<-doneCh
	cancelFn()
}
