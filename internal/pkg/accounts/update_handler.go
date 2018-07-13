package accounts

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/vectorhacker/bank/internal/pkg/events"
	domain "github.com/vectorhacker/bank/internal/pkg/events/accounts"
)

// UpdateHandler implements the event.Handler interface. It
// updates the state of the accounts write model
type UpdateHandler struct {
	db *gorm.DB
}

func NewUpdateHandler(db *gorm.DB) *UpdateHandler {
	return &UpdateHandler{
		db: db,
	}
}

// On implements the events.Handler interface. It handles
// Account events to update the write model
func (h *UpdateHandler) On(event events.Event) error {

	log.Println("handling event", event.ID())

	db := h.db.Begin()
	switch ev := event.(type) {
	case *domain.AccountCreated:
		account := Account{
			ID:      ev.AggregateID(),
			Balance: ev.StartingBalance,
		}
		if !db.NewRecord(account) {
			db = db.Delete(&account)
		}
		db = db.Create(&account)

	case *domain.AccountCredited:
		account := Account{}
		db = db.Where(Account{ID: ev.AggregateID()}).First(&account)

		db = db.Model(&account).Updates(Account{
			Balance: account.Balance + ev.Amount,
		})

	case *domain.AccountDebited:
		account := Account{}
		db = db.Where(Account{ID: ev.AggregateID()}).First(&account)

		newBalance := account.Balance - ev.Amount
		db = db.Model(&account).Update("Balance", newBalance)
	}

	if db.Error != nil {
		db = db.Rollback()
		return db.Error
	}

	db = db.Commit()

	return db.Error
}
