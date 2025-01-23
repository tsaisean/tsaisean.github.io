package employee

import repo "iHR/db/repositories"

type EmployeeHandler struct {
	repo repo.EmployeeRepository
}

func NewEmployeeHandler(r repo.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{repo: r}
}
