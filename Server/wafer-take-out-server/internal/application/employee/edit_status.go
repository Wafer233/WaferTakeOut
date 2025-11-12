package employeeApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/employee"
)

type StatusFlipsDTO struct {
	ID int64 `form:"id"`
}

func (svc *EmployeeService) StatusFlips(ctx context.Context,
	status int, id int64, curId int64) error {

	emp := &employee.Employee{
		Id:         id,
		Status:     status,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}

	err := svc.repo.UpdateStatusByID(ctx, emp)
	return err
}
