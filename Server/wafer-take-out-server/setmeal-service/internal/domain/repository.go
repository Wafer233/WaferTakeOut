package domain

import "context"

type SetMealRepository interface {
	FindPage(context.Context, int64, string, int, int, int) ([]*SetMeal, int64, error)
	Create(context.Context, *SetMeal, []*SetMealDish) error
	UpdateStatus(context.Context, *SetMeal) error
	Delete(context.Context, []int64) error
	FindById(context.Context, int64) (*SetMeal, []*SetMealDish, error)
	Update(context.Context, *SetMeal, []*SetMealDish) error
	FindByCategoryId(context.Context, int64) ([]*SetMeal, error)
	FindDishById(context.Context, int64) ([]*SetMealDish, error)

	//shoppingcart
	FindDetailById(context.Context, int64) (string, string, float64, error)
}
