package query

import (
	"github.com/jinzhu/gorm"
	"github.com/vectorhacker/bank/internal/pkg/events"
	domain "github.com/vectorhacker/bank/internal/pkg/events/accounts"
)

type Accounts struct {
	db *gorm.DB
}

// On implements the events.Handler interface. It handles
// Account events to update the write model
func (h *Accounts) On(event events.Event) error {

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
		db.Create(&account)

	case *domain.AccountCredited:
		account := Account{}
		db = db.Where("id = ?", ev.AggregateID()).First(&account)

		db = db.Model(&account).Updates(Account{
			Balance: account.Balance + ev.Amount,
		})

	case *domain.AccountDebited:
		account := Account{}
		db = db.Where("id = ?", ev.AggregateID()).First(&account)

		db = db.Model(&account).Updates(Account{
			Balance: account.Balance - ev.Amount,
		})
	}

	if db.Error != nil {
		db = db.Rollback()
		return db.Error
	}

	db.Commit()

	return db.Error
}
