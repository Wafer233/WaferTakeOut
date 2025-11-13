package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/domain"
	"github.com/jinzhu/copier"
)

func (svc *EmployeeService) UpdateEmployee(ctx context.Context,
	dto *AddEmployeeDTO, updateId int64) error {

	entity := &domain.Employee{}

	err := copier.Copy(&entity, &dto)
	if err != nil {
		return err
	}
	entity.UpdateUser = updateId
	entity.UpdateTime = time.Now()
	err = svc.repo.UpdateById(ctx, entity)
	return err
}
