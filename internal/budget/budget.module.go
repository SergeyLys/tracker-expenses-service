package budget

import (
	"net/http"
)

func RegisterBudgetModule(mux *http.ServeMux) {
	service := NewBudgetPeriodService()
	controller := NewBudgetPeriodController(service)

	mux.HandleFunc("GET /budgets", controller.GetAllBudgets)
}
