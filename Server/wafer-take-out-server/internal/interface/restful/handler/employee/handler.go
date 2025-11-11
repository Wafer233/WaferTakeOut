package employeeHandler

import employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"

type EmployeeHandler struct {
	svc *employeeApp.EmployeeService
}

func NewEmployeeHandler(svc *employeeApp.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		svc: svc,
	}
}
