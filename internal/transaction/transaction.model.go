package transaction

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "income"
	TransactionTypeExpense TransactionType = "expense"
)

type Transaction struct {
	ID                 string          `json:"id"`
	Amount             decimal.Decimal `json:"amount"`
	Type               TransactionType `json:"type"`
	CategoryID         string          `json:"categoryId,omitempty"`
	PeriodID           string          `json:"periodId,omitempty"`
	Date               time.Time       `json:"date"`
	Note               string          `json:"note,omitempty"`
	RecurringPaymentID string          `json:"recurringPaymentId,omitempty"`
	CreatedAt          time.Time       `json:"createdAt"`
}
