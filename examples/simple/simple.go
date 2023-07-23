package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/crxfoz/outbox"
	"github.com/crxfoz/outbox/worker"
)

// OutboxEntity is an entity stored in outbox table
type OutboxEntity struct {
	ID   int
	Data []byte
}

func (oe OutboxEntity) GetID() int {
	return oe.ID
}

// RecordEntity is an entity that Process consumes
type RecordEntity struct {
	Data string
}

type Handler struct {
	unprocessed map[int]OutboxEntity // ignore ordering as it's just an example
}

func (h *Handler) ToEvent(item OutboxEntity) (RecordEntity, error) {
	return RecordEntity{Data: string(item.Data)}, nil
}

func (h *Handler) RecordsByStatus(_ context.Context, _ outbox.Status) (uint, error) {
	count := uint(len(h.unprocessed))
	fmt.Println("getting RecordsByStatus:", count)
	return count, nil
}

func (h *Handler) Process(_ context.Context, record RecordEntity) error {
	fmt.Printf("processing: %+v\n", record)
	time.Sleep(time.Second * 5)

	if rand.Intn(2) == 1 {
		err := fmt.Errorf("kafka producer isn't available")
		fmt.Println("got an error on processing:", err.Error())
		return outbox.NewRetryableError(err)
	}

	return nil
}

func (h *Handler) Commit() error {
	fmt.Println("commit")
	return nil
}

func (h *Handler) Rollback() error {
	fmt.Println("rollback")
	return nil
}

func (h *Handler) RecordWithLock(_ context.Context, _ *Handler, limit uint, _ bool, _ outbox.Status) ([]OutboxEntity, error) {
	fmt.Println("getting records with lock")

	if len(h.unprocessed) == 0 {
		return nil, outbox.ErrNoMoreEntries
	}

	out := make([]OutboxEntity, 0, limit)
	for _, v := range h.unprocessed {
		out = append(out, OutboxEntity{
			ID:   v.ID,
			Data: v.Data,
		})
	}
	return out, nil
}

func (h *Handler) UpdateStatus(_ context.Context, _ *Handler, id int, _ outbox.Status) error {
	fmt.Println("updating status:", id)
	delete(h.unprocessed, id)
	return nil
}

func (h *Handler) GetTx(_ context.Context) (*Handler, error) {
	return h, nil
}

func main() {
	handler := &Handler{unprocessed: map[int]OutboxEntity{
		1: {
			ID:   1,
			Data: []byte("id-1"),
		},
		2: {
			ID:   2,
			Data: []byte("id-2"),
		},
		3: {
			ID:   3,
			Data: []byte("id-3"),
		},
	}}

	cfg := worker.NewConfigOrdered()

	wrk := worker.NewWorker[int, OutboxEntity, RecordEntity, *Handler](cfg, handler, handler, handler)

	err := wrk.Start(context.Background())
	if err != nil {
		fmt.Printf("worker stopped with error: %s", err.Error())
	}
}
