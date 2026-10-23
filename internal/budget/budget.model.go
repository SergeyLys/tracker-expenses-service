package budget

import (
	"time"
)

type BudgetPeriod struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	StartsAt  time.Time `json:"startsAt"`
	EndsAt    time.Time `json:"endsAt"`
	CreatedAt time.Time `json:"createdAt"`
}
