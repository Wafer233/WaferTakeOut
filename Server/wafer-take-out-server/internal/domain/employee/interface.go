package employee

import "context"

type EmployeeRepo interface {
	GetByUsername(context.Context, string) (*Employee, error)
	Insert(context.Context, *Employee) error
}
