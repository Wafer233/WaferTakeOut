package setmeal_dish

import "context"

type SetMealDishRepository interface {
	Inserts(context.Context, []*SetMealDish) error
	GetsBySetMealId(context.Context, int64) ([]*SetMealDish, error)
	UpdatesBySetMealId(context.Context, []*SetMealDish) error
}
