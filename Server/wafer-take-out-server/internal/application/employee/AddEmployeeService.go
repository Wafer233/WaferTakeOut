package employeeApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/employee"
)

func (svc *EmployeeService) AddEmployee(ctx context.Context, dto *AddEmployeeDTO, id int64) error {

	entity := &employee.Employee{
		IDNumber: dto.IDNumber,
		Name:     dto.Name,
		Phone:    dto.Phone,
		Sex:      dto.Sex,
		Username: dto.UserName,
	}
	entity.Password = "123456"
	entity.Status = 1
	entity.CreateTime = time.Now()
	entity.UpdateTime = time.Now()

	entity.CreateUser = id
	entity.UpdateUser = id

	err := svc.repo.Insert(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}
