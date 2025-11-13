package application

import (
	"context"
)

func (svc *EmployeeService) GetEmployee(ctx context.Context, id int64) (*Employee, error) {

	employee, err := svc.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	employeeVO := Employee{
		Id:         employee.Id,
		Name:       employee.Name,
		Username:   employee.Username,
		Password:   employee.Password,
		Phone:      employee.Phone,
		Sex:        employee.Sex,
		IDNumber:   employee.IDNumber,
		Status:     employee.Status,
		CreateTime: employee.UpdateTime.Format("2006-01-02 15:04"),
		UpdateTime: employee.UpdateTime.Format("2006-01-02 15:04"),
		CreateUser: employee.CreateUser,
		UpdateUser: employee.UpdateUser,
	}

	return &employeeVO, nil
}
