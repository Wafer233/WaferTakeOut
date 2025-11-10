package employeeApp

import (
	employeeInfra "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/employee"
)

type EmployeeService struct {
	repo *employeeInfra.EmployeeRepository
}

func NewEmployeeService(repo *employeeInfra.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}
