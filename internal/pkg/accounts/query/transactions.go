package query

import (
	"github.com/jinzhu/gorm"
	"github.com/vectorhacker/bank/internal/pkg/events"
	"github.com/vectorhacker/bank/internal/pkg/events/domain"
)

type Transactions struct {
	db *gorm.DB
}

// On implements the events.Handler interface. It handles
// Account events to update the write model
func (h *Transactions) On(event events.Event) error {

	db := h.db.Begin()
	account := Account{}
	db = db.Where("id = ?", event.AggregateID().String()).First(&account)

	switch ev := event.(type) {
	case *domain.AccountCredited:

		transaction := Transaction{
			ID:      ev.TransactionID,
			Amount:  ev.Amount,
			Account: account,
			Type:    Credit,
		}

		if !db.NewRecord(transaction) {
			db = db.Delete(&transaction)
		}

		db = db.Create(&transaction)

	case *domain.AccountDebited:
		transaction := Transaction{
			ID:      ev.TransactionID,
			Amount:  ev.Amount,
			Account: account,
			Type:    Debit,
		}

		if !db.NewRecord(transaction) {
			db = db.Delete(&transaction)
		}

		db = db.Create(&transaction)
	}

	if db.Error != nil {
		db = db.Rollback()
		return db.Error
	}

	db = db.Commit()

	return db.Error
}
