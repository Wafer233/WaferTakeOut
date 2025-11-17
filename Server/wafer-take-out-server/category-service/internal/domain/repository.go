package domain

import (
	"context"
)

type CategoryRepository interface {
	Create(context.Context, *Category) error
	FindPage(context.Context, string, int, int, int) ([]*Category, int64, error)
	Update(context.Context, *Category) error
	UpdateStatus(context.Context, *Category) error
	Delete(context.Context, int64) error
	FindByType(context.Context, int) ([]*Category, error)
	FindNameById(context.Context, int64) (string, error)
}
