package application

import (
	employeeInfra "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/infrastructure"
)

type EmployeeService struct {
	repo *employeeInfra.EmployeeRepository
}

func NewEmployeeService(repo *employeeInfra.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}
