package budget

import (
	"database/sql"
	"net/http"
)

func RegisterBudgetModule(mux *http.ServeMux, db *sql.DB) {
	repository := NewPostgresBudgetRepository(db)
	service := NewBudgetPeriodService(repository)
	controller := NewBudgetPeriodController(service)

	mux.HandleFunc("GET /budgets", controller.GetAllBudgets)
}
