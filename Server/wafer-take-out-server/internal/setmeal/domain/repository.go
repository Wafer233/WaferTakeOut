package domain

import "context"

type SetMealRepository interface {
	FindPage(context.Context, int64, string, int, int, int) ([]*SetMeal, int64, error)
	Create(context.Context, *SetMeal, []*SetMealDish) error
	UpdateStatus(context.Context, *SetMeal) error
	Delete(context.Context, []int64) error
	FindById(context.Context, int64) (*SetMeal, []*SetMealDish, error)
	Update(context.Context, *SetMeal, []*SetMealDish) error
}
