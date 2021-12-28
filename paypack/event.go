package paypack

import (
	"context"
	"time"
)

type Event struct {
	ID        string
	Data      EventData
	Kind      string
	CreatedAt string
}

type EventData struct {
	Ref         string     `json:"ref"`
	Status      string     `json:"status"`
	Amount      float64    `json:"amount"`
	Client      string     `json:"client"`
	Kind        string     `json:"kind"`
	Fee         float64    `json:"fee"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ProcessedAt *time.Time `json:"processed_at,omitempty"`
}

type EventList struct {
	Kind      *string
	From      *string
	To        *string
	EventKind *string
	Offset    uint64
	Limit     uint64
	Total     uint64
	Events    []Event
}

type EventService interface {
	// List returns a collection of events that match a list of filters query params
	List(ctx context.Context, options ...Option) (*EventList, error)
}
