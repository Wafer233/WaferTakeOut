package setmeal_dish

import "context"

type SetMealDishRepository interface {
	Inserts(context.Context, []*SetMealDish) error
}
