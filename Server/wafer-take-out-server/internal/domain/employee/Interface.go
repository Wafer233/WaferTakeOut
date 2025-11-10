package employee

import "context"

type EmployeeRepo interface {
	GetByUsername(context.Context, string) (*Employee, error)
	Insert(context.Context, *Employee) error
	GetByUsernamePaged(ctx context.Context, name string, page int, pageSize int) (int64, []Employee, error)
	UpdateStatusByID(ctx context.Context, status int, id int64) error
}
