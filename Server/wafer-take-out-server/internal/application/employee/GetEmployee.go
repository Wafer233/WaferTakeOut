package employeeApp

import (
	"context"

	"github.com/jinzhu/copier"
)

func (svc *EmployeeService) GetEmployee(ctx context.Context, id int64) (*Employee, error) {

	employee, err := svc.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	var employeeVO Employee
	err = copier.Copy(&employeeVO, employee)
	if err != nil {
		return nil, err
	}
	return &employeeVO, nil
}
