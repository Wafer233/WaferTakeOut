package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/domain"
)

type AddEmployeeDTO struct {
	ID       int64  `json:"id"`
	IDNumber string `json:"idNumber"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	UserName string `json:"username"`
}

func (svc *EmployeeService) AddEmployee(ctx context.Context, dto *AddEmployeeDTO, id int64) error {

	entity := &domain.Employee{
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
