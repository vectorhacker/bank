package events

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Model implements the Event interface
type Model struct {
	EventAggregateID uuid.UUID `json:"aggregate_id"`
	EventID          uuid.UUID `json:"id"`
	EventAt          time.Time `json:"timestamp"`
}

// AggregateID implements the Event interface. It returns the
// aggregate id for the event
func (m Model) AggregateID() uuid.UUID {
	return m.EventAggregateID
}

// ID implements the Event interface. It returns the ID of an event
func (m Model) ID() uuid.UUID {
	return m.EventID
}

func (m Model) At() time.Time {
	return m.EventAt
}

// Event represents an Event in the System
type Event interface {
	AggregateID() uuid.UUID
	ID() uuid.UUID
	At() time.Time
}

// Events is an array of event
type Events []Event

func (e Events) Len() int {
	return len(e)
}

func (e Events) Swap(a, b int) {
	e[a], e[b] = e[b], e[a]
}

func (e Events) Less(a, b int) bool {
	return e[a].At().After(e[b].At())
}
