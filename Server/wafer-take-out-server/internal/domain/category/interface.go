package category

import "context"

type CategoryRepo interface {
	Insert(context.Context, *Category) error
	GetsByPaged(context.Context, string, int, int, int) ([]*Category, int64, error)
	UpdateById(context.Context, *Category) error
	UpdateStatusById(context.Context, int64, int) error
	DeleteById(context.Context, int64) error
}
