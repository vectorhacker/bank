package query

import uuid "github.com/satori/go.uuid"

//TransactionType enum
type TransactionType string

//TransactionType enum
const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

// Account represents a bank account
type Account struct {
	ID      uuid.UUID `gorm:"primary_key"`
	Balance int64
}

// Transaction represents a bank transaction
type Transaction struct {
	ID      uuid.UUID `gorm:"primary_key"`
	Account Account
	Amount  int64
	Type    TransactionType
}
