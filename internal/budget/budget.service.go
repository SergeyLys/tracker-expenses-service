package budget

import (
	"fmt"
	"time"
)

type BudgetPeriodService struct {
}

func NewBudgetPeriodService() *BudgetPeriodService {
	return &BudgetPeriodService{}
}

func (repository *BudgetPeriodService) GetAllBudgets() []BudgetPeriod {
	budgets := make([]BudgetPeriod, 0)

	for i := range 5 {
		startsAt := time.Date(2026, time.October, 1, 0, 0, 0, 0, time.UTC)
		endsAt := startsAt.AddDate(0, 1, 0)

		budgets = append(budgets, BudgetPeriod{
			ID:        fmt.Sprintf("bp-%03d", i+1),
			Title:     startsAt.Format("January 2006"),
			StartsAt:  startsAt,
			EndsAt:    endsAt,
			CreatedAt: time.Now(),
		})
	}

	return budgets
}
