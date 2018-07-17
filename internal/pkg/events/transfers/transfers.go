package transfers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/vectorhacker/bank/internal/pkg/events"
)

// TransferBegun event
type TransferBegun struct {
	events.Model
	FromAccount uuid.UUID
	ToAccount   uuid.UUID
	Description string
	Amount      int64
}

// TransferCreditAccountBegun event
type TransferCreditAccountBegun struct {
	events.Model
}

type TransferCreditCompleted struct {
	events.Model
	TransactionID string
}

// TransferDebitAccountBegun event
type TransferDebitAccountBegun struct {
	events.Model
}

type TransferDebitCompleted struct {
	events.Model
	TransactionID string
}

// TransferDebitFailed event
type TransferDebitFailed struct {
	events.Model
}

// TransferCreditFailed event
type TransferCreditFailed struct {
	events.Model
}

// TransferCompleted event
type TransferCompleted struct {
	events.Model
}

type TransferFailed struct {
	events.Model
}
