package domain

import "context"

type CategoryRepository interface {
	Insert(context.Context, *Category) error
	GetsByPaged(context.Context, string, int, int, int) ([]*Category, int64, error)
	UpdateById(context.Context, *Category) error
	UpdateStatusById(context.Context, *Category) error
	DeleteById(context.Context, int64) error
	GetsByType(context.Context, int) ([]*Category, error)
	GetById(context.Context, int64) (*Category, error)
}
