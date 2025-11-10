package employeeApp

import (
	"context"

	"github.com/jinzhu/copier"
)

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
