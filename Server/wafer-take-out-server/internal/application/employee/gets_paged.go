package employeeApp

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/times"
	"github.com/jinzhu/copier"
)

type PageDTO struct {
	Name     string `form:"name"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

type PageVO struct {
	Total     int64      `json:"total"`
	Employees []Employee `json:"records"`
}

type Employee struct {
	Id         int64          `json:"id"`
	Name       string         `json:"name"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Phone      string         `json:"phone"`
	Sex        string         `json:"sex"`
	IDNumber   string         `json:"idNumber"`
	Status     int            `json:"status"`
	CreateTime times.JSONTime `json:"createTime"`
	UpdateTime times.JSONTime `json:"updateTime"`
	CreateUser int64          `json:"createUser"`
	UpdateUser int64          `json:"updateUser"`
}

func (svc *EmployeeService) PageQuery(ctx context.Context, dto *PageDTO) (*PageVO, error) {

	total, employees, err := svc.repo.GetByUsernamePaged(ctx, dto.Name, dto.Page, dto.PageSize)

	if err != nil {
		return nil, err
	}

	var employeesVO []Employee

	err = copier.Copy(&employeesVO, &employees)
	if err != nil {
		return nil, err
	}

	return &PageVO{
		Total:     total,
		Employees: employeesVO,
	}, nil
}
