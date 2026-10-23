package budget

import (
	"encoding/json"
	"net/http"
)

type BudgetPeriodController struct {
	service *BudgetPeriodService
}

type GetAllBudgetsResponse struct {
	Budgets []BudgetPeriod `json:"budgets"`
}

func NewBudgetPeriodController(service *BudgetPeriodService) *BudgetPeriodController {
	return &BudgetPeriodController{
		service: service,
	}
}

func (repository *BudgetPeriodController) GetAllBudgets(w http.ResponseWriter, r *http.Request) {
	err, budgets := repository.service.GetAllBudgets()

	if err != nil {
		w.WriteHeader(http.StatusOK)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetAllBudgetsResponse{
		Budgets: budgets,
	})
}
