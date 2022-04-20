package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/quarksgroup/paypack-go/paypack"
)

type eventService struct {
	http *wrapper
}

// List implements the paypack list events with the given query slice parameters
func (s *eventService) List(ctx context.Context, options ...paypack.Option) (*paypack.EventList, error) {

	var params string
	if len(options) > 0 {
		for _, option := range options {
			params += fmt.Sprintf("%s&", option)
		}
	}

	endpoint := fmt.Sprintf("events/transactions?%s", params)

	out := new(listEventResponse)

	_, err := s.http.do(ctx, "GET", endpoint, nil, out, nil)

	res := &paypack.EventList{
		Kind:         out.Kind,
		Offset:       out.Offset,
		EventKind:    out.EventKind,
		Limit:        out.Limit,
		Total:        out.Total,
		Transactions: make([]paypack.Event, 0),
	}

	for _, event := range out.Transactions {

		transaction := new(paypack.Transaction)

		if err := json.Unmarshal([]byte(event.Data), transaction); err != nil {
			return nil, err
		}

		resp := paypack.Event{
			ID:        event.ID,
			Data:      *transaction,
			Kind:      event.Kind,
			CreatedAt: event.CreatedAt,
		}

		res.Transactions = append(res.Transactions, resp)
	}
	return res, err

}

var _ (paypack.EventService) = (*eventService)(nil)
