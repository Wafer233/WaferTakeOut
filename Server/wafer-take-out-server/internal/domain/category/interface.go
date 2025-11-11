package category

import "context"

type CategoryRepo interface {
	Insert(context.Context, *Category) error
	GetsByNamePaged(context.Context, string, int, int) ([]*Category, int64, error)
}
