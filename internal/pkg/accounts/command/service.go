package command

import (
	"context"
	"errors"
	"log"

	"github.com/satori/go.uuid"
	"github.com/vectorhacker/bank/internal/pkg/accounts"
	domain "github.com/vectorhacker/bank/internal/pkg/events/accounts"

	e "github.com/vectorhacker/bank/internal/pkg/events"

	"github.com/jinzhu/gorm"
	pb "github.com/vectorhacker/bank/pb/accounts"
)

// Errors
var (
	ErrNoSufficientFunds = errors.New("funds not sufficient to complete transaction")
	ErrInvalidAmount     = errors.New("invalid amount requested")
)

// Service implements the proto.AccountsCommandServer interface.
type Service struct {
	db         *gorm.DB
	dispatcher *e.Dispatcher
}

func New(db *gorm.DB, dispatcher *e.Dispatcher) *Service {
	return &Service{
		db:         db,
		dispatcher: dispatcher,
	}
}

// CreateAccount implements the proto.AccountsCommandServer interface.
// It creates a new bank account and returns the id
func (s *Service) CreateAccount(
	ctx context.Context,
	r *pb.CreateAccountRequest,
) (*pb.CreateAccountResponse, error) {

	events := []e.Event{}
	created := domain.AccountCreated{
		Model: e.Model{
			EventAggregateID: uuid.Must(uuid.NewV4()),
			EventID:          uuid.Must(uuid.NewV4()),
		},
	}
	events = append(events, created)

	if r.InitialDeposit > 0 {
		credit := domain.AccountCredited{
			Model: e.Model{
				EventAggregateID: created.AggregateID(),
				EventID:          uuid.Must(uuid.NewV4()),
			},
			Amount:        r.InitialDeposit,
			TransactionID: uuid.Must(uuid.NewV4()),
			Description:   "Initial Deposit",
		}

		events = append(events, credit)
	}

	// commit events
	if err := s.dispatcher.Dispatch(events...); err != nil {
		return nil, err
	}

	return &pb.CreateAccountResponse{
		ID: created.AggregateID().String(),
	}, nil
}

// DebitAccount implements the proto.AccountsCommandServer interface.
// It debits an account
func (s *Service) DebitAccount(
	ctx context.Context,
	r *pb.DebitAccountRequest,
) (*pb.DebitAccountResponse, error) {
	id, err := uuid.FromString(r.ID)
	if err != nil {
		return nil, err
	}

	account := accounts.Account{}
	db := s.db.Where(&accounts.Account{ID: id}).First(&account)
	if db.Error != nil {
		return nil, db.Error
	}

	futureBalance := account.Balance - r.Amount
	log.Println("future balance", futureBalance)
	if futureBalance < 0 {
		return nil, ErrNoSufficientFunds
	}

	events := []e.Event{}
	debit := domain.AccountDebited{
		Model: e.Model{
			EventAggregateID: id,
			EventID:          uuid.Must(uuid.NewV4()),
		},
		TransactionID: uuid.Must(uuid.NewV4()),
		Amount:        r.Amount,
		Description:   r.Description,
	}

	events = append(events, debit)
	err = s.dispatcher.Dispatch(events...)
	if err != nil {
		return nil, err
	}

	return &pb.DebitAccountResponse{
		TransactionID: debit.TransactionID.String(),
	}, nil
}

// CreditAccount implements the proto.AccountsCommandServer interface.
func (s *Service) CreditAccount(
	ctx context.Context,
	r *pb.CreditAccountRequest,
) (*pb.CreditAccountResponse, error) {

	id, err := uuid.FromString(r.ID)
	if err != nil {
		return nil, err
	}

	account := accounts.Account{}
	db := s.db.Where(&accounts.Account{ID: id}).First(&account)
	if db.Error != nil {
		return nil, db.Error
	}

	if r.Amount <= 0 {
		return nil, ErrInvalidAmount
	}

	events := []e.Event{}
	credit := domain.AccountCredited{
		Model: e.Model{
			EventAggregateID: id,
			EventID:          uuid.Must(uuid.NewV4()),
		},
		TransactionID: uuid.Must(uuid.NewV4()),
		Amount:        r.Amount,
		Description:   r.Description,
	}

	events = append(events, credit)
	err = s.dispatcher.Dispatch(events...)
	if err != nil {
		return nil, err
	}

	return &pb.CreditAccountResponse{
		TransactionID: credit.TransactionID.String(),
	}, nil
}
