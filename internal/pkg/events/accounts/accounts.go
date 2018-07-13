package accounts

import (
	"github.com/satori/go.uuid"
	"github.com/vectorhacker/bank/internal/pkg/events"
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
	Description   string
	Amount        int64
}

// AccountDebited event
type AccountDebited struct {
	events.Model
	TransactionID uuid.UUID
	Description   string
	Amount        int64
}

type AccountClosed struct {
	events.Model
}

type AccountDebitFailed struct {
	AccountDebited
	Reason string
}

type AccountCreditFailed struct {
	AccountCredited
	Reason string
}
