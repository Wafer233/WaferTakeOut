package setmeal

import "context"

type SetMealRepository interface {
	GetsPaged(context.Context, int64, string, int, int, int) ([]*SetMeal, int64, error)
	Insert(context.Context, *SetMeal) error
	UpdateStatusById(context.Context, *SetMeal) error
	DeletesByIds(context.Context, []int64) error
	GetById(context.Context, int64) (*SetMeal, error)
}
