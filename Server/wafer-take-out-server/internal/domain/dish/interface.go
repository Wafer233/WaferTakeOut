package dish

import "context"

type DishRepository interface {
	GetsPaged(context.Context, string, int64, int, int, int) ([]*Dish, int64, error)
	DeletesById(ctx context.Context, ids []int64) error
	Insert(context.Context, *Dish) error
}
