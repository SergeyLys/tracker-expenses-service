package budget

type BudgetPeriodService struct {
	repository *PostgresBudgetRepository
}

func NewBudgetPeriodService(repository *PostgresBudgetRepository) *BudgetPeriodService {
	return &BudgetPeriodService{
		repository: repository,
	}
}

func (service *BudgetPeriodService) GetAllBudgets() ([]BudgetPeriod, error) {
	return service.repository.GetAllBudgets()
}
