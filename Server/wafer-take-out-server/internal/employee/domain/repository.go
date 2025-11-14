package domain

import "context"

type EmployeeRepo interface {
	FindByUsername(context.Context, string) (*Employee, error)
	Create(context.Context, *Employee) error
	FindPage(context.Context, string, int, int) (int64, []Employee, error)
	UpdateStatus(context.Context, *Employee) error
	FindById(context.Context, int64) (*Employee, error)
	Update(context.Context, *Employee) error
	UpdatePassword(context.Context, int64, string, string) error
}
