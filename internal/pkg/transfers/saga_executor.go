package transfers

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/vectorhacker/bank/internal/pkg/events"
	pb "github.com/vectorhacker/bank/pb/accounts"
)

var (
	// ErrTransferBegun refers to when a transfer has already begun execution
	ErrTransferBegun = errors.New("transfer has already begun")
	// ErrInvalidState occurs when an invalid state is reachd
	ErrInvalidState = errors.New("invalid state")
)

// Executor is a Saga Execution Engine
type Executor struct {
	db         *gorm.DB
	dispatcher events.Dispatcher
	accounts   pb.AccountsCommandClient
}

func NewExecutor(
	db *gorm.DB,
	dispatcher events.Dispatcher,
	accounts pb.AccountsCommandClient,
) *Executor {
	return &Executor{
		db:         db,
		dispatcher: dispatcher,
		accounts:   accounts,
	}
}

// On implements the events.Handler interface
func (e Executor) On(event events.Event) error {

	log.Printf("%v", event)

	db := e.db.Begin()

	transfer := Transfer{
		accounts: e.accounts,
	}

	db = db.Where(Transfer{ID: event.AggregateID(), State: Uninitialized}).FirstOrCreate(&transfer)

	events, err := transfer.Handle(event)
	if err != nil {
		db = db.Rollback()
		if db.Error != nil {
			return db.Error
		}
		return err
	}

	db = db.Save(transfer)

	if err := db.Commit().Error; err != nil {
		return err
	}

	return e.dispatcher.Dispatch(events...)
}
