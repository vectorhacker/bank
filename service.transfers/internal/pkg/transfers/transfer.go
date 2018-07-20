package transfers

import (
	"context"
	"log"
	"time"

	"github.com/satori/go.uuid"
	"github.com/vectorhacker/bank/core/events"
	pb "github.com/vectorhacker/bank/service.accounts/pb"
	td "github.com/vectorhacker/bank/service.transfers/pkg/events"
)

// State represents a transaction state
type State string

// State enum
const (
	Uninitialized State = ""
	Begun               = "begun"
	Transfering         = "transfering"
	Completed           = "completed"
	Failed              = "failed"
)

// Transfer representation
type Transfer struct {
	ID          uuid.UUID `gorm:"primary_key"`
	FromAccount uuid.UUID
	ToAccount   uuid.UUID
	State       State
	Amount      int64
	Description string
	accounts    pb.AccountsCommandClient `gorm:"-"`
}

// Handle receives an event to handle, issues commands and produces events
func (t *Transfer) Handle(event events.Event) (events.Events, error) {

	switch ev := event.(type) {
	case *td.TransferBegun:
		log.Println("Began transfer", t.ID.String())
		if t.State != Uninitialized {
			log.Println("expected uninitialized state", t.State)
			return nil, ErrInvalidState
		}

		t.State = Begun
		t.FromAccount = ev.FromAccount
		t.ToAccount = ev.ToAccount
		t.Amount = ev.Amount
		t.Description = ev.Description

		return events.Events{
			td.TransferDebitAccountBegun{
				Model: events.Model{
					EventAggregateID: t.ID,
					EventID:          uuid.Must(uuid.NewV4()),
					EventAt:          time.Now(),
				},
			},
		}, nil

	case *td.TransferDebitAccountBegun:
		log.Println("Debiting transfer", t.ID.String())
		if t.State != Begun {
			return nil, ErrInvalidState
		}

		t.State = Transfering

		ctx := context.Background()

		response, err := t.accounts.DebitAccount(ctx, &pb.DebitAccountRequest{
			ID:            t.FromAccount.String(),
			CorrelationID: t.ID.String(),
			Amount:        t.Amount,
			Timestamp:     time.Now().Unix(),
			Description:   t.Description,
		})
		if err != nil {
			log.Println(err)
			return events.Events{
				&td.TransferDebitFailed{
					Model: events.Model{
						EventAggregateID: t.ID,
						EventID:          uuid.Must(uuid.NewV4()),
						EventAt:          time.Now(),
					},
				},
			}, nil
		}

		return events.Events{
			&td.TransferDebitCompleted{
				Model: events.Model{
					EventAggregateID: t.ID,
					EventID:          uuid.Must(uuid.NewV4()),
					EventAt:          time.Now(),
				},
				TransactionID: response.TransactionID,
			},
		}, nil

	case *td.TransferDebitFailed:
		log.Println("Debit failed", t.ID.String())
		if t.State != Transfering {
			return nil, ErrInvalidState
		}

		t.State = Failed

	case *td.TransferDebitCompleted:
		log.Println("Debited transfer", t.ID.String())
		if t.State != Transfering {
			return nil, ErrInvalidState
		}

		return events.Events{
			&td.TransferCreditAccountBegun{
				Model: events.Model{
					EventAggregateID: t.ID,
					EventID:          uuid.Must(uuid.NewV4()),
					EventAt:          time.Now(),
				},
			},
		}, nil

	case *td.TransferCreditAccountBegun:
		log.Println("Crediting transfer", t.ID.String())
		if t.State != Transfering {
			return nil, ErrInvalidState
		}

		ctx := context.Background()

		response, err := t.accounts.CreditAccount(ctx, &pb.CreditAccountRequest{
			ID:            t.ToAccount.String(),
			CorrelationID: t.ID.String(),
			Amount:        t.Amount,
			Timestamp:     time.Now().Unix(),
			Description:   t.Description,
		})
		if err != nil {
			log.Println(err)
			return events.Events{
				&td.TransferCreditFailed{
					Model: events.Model{
						EventAggregateID: t.ID,
						EventID:          uuid.Must(uuid.NewV4()),
						EventAt:          time.Now(),
					},
				},
			}, nil
		}

		return events.Events{
			&td.TransferCreditCompleted{
				Model: events.Model{
					EventAggregateID: t.ID,
					EventID:          uuid.Must(uuid.NewV4()),
					EventAt:          time.Now(),
				},
				TransactionID: response.TransactionID,
			},
		}, nil

	case *td.TransferCreditCompleted:
		log.Println("Credited transfer", t.ID.String())
		if t.State != Transfering {
			return nil, ErrInvalidState
		}

		return events.Events{
			&td.TransferCompleted{
				Model: events.Model{
					EventAggregateID: t.ID,
					EventID:          uuid.Must(uuid.NewV4()),
					EventAt:          time.Now(),
				},
			},
		}, nil

	case *td.TransferCreditFailed:
		log.Println("Credited transfer failed", t.ID.String())
		if t.State != Transfering {
			return nil, ErrInvalidState
		}

		ctx := context.Background()
		for {
			_, err := t.accounts.CreditAccount(ctx, &pb.CreditAccountRequest{
				ID:            t.FromAccount.String(),
				CorrelationID: t.ID.String(),
				Amount:        t.Amount,
				Timestamp:     time.Now().Unix(),
				Description:   "correction: " + t.Description,
			})
			if err != nil {
				continue
			}

			return events.Events{
				&td.TransferFailed{
					Model: events.Model{
						EventAggregateID: t.ID,
						EventID:          uuid.Must(uuid.NewV4()),
						EventAt:          time.Now(),
					},
				},
			}, nil
		}

	case *td.TransferCompleted:
		log.Println("Completed transfer", t.ID.String())
		if t.State != Transfering {
			return nil, ErrInvalidState
		}

		t.State = Completed
	}

	return nil, nil
}
