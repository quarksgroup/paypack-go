package api

import (
	"context"
	"fmt"

	"github.com/quarksgroup/paypack-go/paypack"
)

type eventService struct {
	http *wrapper
}

// List implements the paypack list events with the given query slice parameters
func (s *eventService) List(ctx context.Context, options ...string) (*paypack.EventList, error) {

	var params string

	if len(options) > 0 {
		for _, option := range options {
			params += fmt.Sprintf("%s&", option)
		}
	}

	endpoint := fmt.Sprintf("events/transactions?%s", params)

	out := new(listEventResponse)

	_, err := s.http.do(ctx, "GET", endpoint, nil, out)

	res := &paypack.EventList{
		Kind:      out.Kind,
		Offset:    out.Offset,
		EventKind: out.EventKind,
		Limit:     out.Limit,
		Total:     out.Total,
		Events:    make([]paypack.Event, 0),
	}
	for _, event := range out.Transactions {

		transaction := &paypack.EventData{
			Ref:         event.Data.Ref,
			Amount:      event.Data.Amount,
			Kind:        event.Data.Kind,
			Status:      event.Data.Status,
			Fee:         event.Data.Fee,
			Client:      event.Data.Client,
			CreatedAt:   event.Data.CreatedAt,
			ProcessedAt: event.Data.ProcessedAt,
		}

		res.Events = append(res.Events, paypack.Event{
			ID:        event.ID,
			Kind:      event.Kind,
			Data:      *transaction,
			CreatedAt: event.CreatedAt,
		})
	}
	return res, err

}

var _ (paypack.EventService) = (*eventService)(nil)
