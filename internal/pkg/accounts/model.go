package accounts

import (
	"github.com/satori/go.uuid"
)

type Account struct {
	ID      uuid.UUID `gorm:"primary_key"`
	Balance int64     `gorm:"NOT NULL"`
}
