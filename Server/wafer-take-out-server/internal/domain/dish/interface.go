package dish

import "context"

type DishRepository interface {
	//UpdateById(context.Context, int64, *Dish) error
	//Insert(context.Context, *Dish) error
	//GetById(context.Context, int64) (*Dish, error)
	//GetsByCategoryId(context.Context, int64) ([]*Dish, error)
	GetsPaged(context.Context, string, int64, int, int, int) ([]*Dish, int64, error)
	DeletesById(ctx context.Context, ids []int64) error
	//UpdateStatus(context.Context, int64, int) error
}
