package events

import (
	uuid "github.com/satori/go.uuid"
)

// Model implements the Event interface
type Model struct {
	EventAggregateID uuid.UUID `json:"aggregate_id"`
	EventID          uuid.UUID `json:"id"`
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

// Event represents an Event in the System
type Event interface {
	AggregateID() uuid.UUID
	ID() uuid.UUID
}
