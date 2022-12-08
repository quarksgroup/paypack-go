package paypack

type Event struct {
	ID        string      `json:"event_id"`
	Data      Transaction `json:"data"`
	Kind      string      `json:"event_kind"`
	CreatedAt string      `json:"created_at"`
}

type EventList struct {
	Kind         *string `json:"kind,omitempty"`
	From         *string `json:"from,omitempty"`
	To           *string `json:"to,omitempty"`
	EventKind    *string `json:"event-kind,omitempty"`
	Offset       uint64  `json:"offset,omitempty"`
	Limit        uint64  `json:"limit,omitempty"`
	Total        uint64  `json:"total,omitempty"`
	Transactions []Event `json:"transactions,omitempty"`
}
