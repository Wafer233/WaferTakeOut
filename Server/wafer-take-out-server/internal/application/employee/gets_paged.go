package employeeApp

import (
	"context"
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
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Sex        string `json:"sex"`
	IDNumber   string `json:"idNumber"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	CreateUser int64  `json:"createUser"`
	UpdateUser int64  `json:"updateUser"`
}

func (svc *EmployeeService) PageQuery(ctx context.Context, dto *PageDTO) (*PageVO, error) {

	total, employees, err := svc.repo.GetByUsernamePaged(ctx, dto.Name, dto.Page, dto.PageSize)

	if err != nil {
		return nil, err
	}

	empVO := make([]Employee, len(employees))
	for i, e := range employees {
		empVO[i] = Employee{
			Id:         e.Id,
			Name:       e.Name,
			Username:   e.Username,
			Password:   e.Password,
			Phone:      e.Phone,
			Sex:        e.Sex,
			IDNumber:   e.IDNumber,
			Status:     e.Status,
			CreateTime: e.CreateTime.Format("2006-01-02 15:04"),
			UpdateTime: e.UpdateTime.Format("2006-01-02 15:04"),
			CreateUser: e.CreateUser,
			UpdateUser: e.UpdateUser,
		}
	}
	return &PageVO{
		Total:     total,
		Employees: empVO,
	}, nil
}
