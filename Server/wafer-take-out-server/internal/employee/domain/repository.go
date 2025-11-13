package domain

import "context"

type EmployeeRepo interface {
	GetByUsername(context.Context, string) (*Employee, error)
	Insert(context.Context, *Employee) error
	GetByUsernamePaged(context.Context, string, int, int) (int64, []Employee, error)
	UpdateStatusByID(context.Context, *Employee) error
	GetById(context.Context, int64) (*Employee, error)
	UpdateById(context.Context, *Employee) error
}
