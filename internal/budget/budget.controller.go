package budget

import (
	"encoding/json"
	"net/http"
)

type BudgetPeriodController struct {
	service *BudgetPeriodService
}

type GetAllBudgetsResponse struct {
	budgets []BudgetPeriod `json:"budgets"`
}

func NewBudgetPeriodController(service *BudgetPeriodService) *BudgetPeriodController {
	return &BudgetPeriodController{
		service: service,
	}
}

func (repository *BudgetPeriodController) GetAllBudgets(w http.ResponseWriter, r *http.Request) {
	budgets := repository.service.GetAllBudgets()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetAllBudgetsResponse{
		budgets: budgets,
	})
}
