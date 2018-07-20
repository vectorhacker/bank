package domain

import (
	"github.com/satori/go.uuid"
	"github.com/vectorhacker/bank/core/events"
)

// AccountCreated event
type AccountCreated struct {
	events.Model
	StartingBalance int64
}

// AccountCredited event
type AccountCredited struct {
	events.Model
	TransactionID uuid.UUID
	CorrelationID string
	Description   string
	Amount        int64
}

// AccountDebited event
type AccountDebited struct {
	events.Model
	TransactionID uuid.UUID
	CorrelationID string
	Description   string
	Amount        int64
}

type AccountClosed struct {
	events.Model
}
