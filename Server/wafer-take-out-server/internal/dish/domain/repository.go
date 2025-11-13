package domain

import "context"

type DishRepository interface {
	Delete(context.Context, []int64) error
	Create(context.Context, *Dish, []*Flavor) error
	FindById(context.Context, int64) (*Dish, []*Flavor, error)
	FindByCategoryId(context.Context, int64) ([]*Dish, error)
	FindPage(context.Context, string, int64, int, int, int) ([]*Dish, int64, error)
	UpdateStatus(context.Context, *Dish) error
	Update(context.Context, *Dish, []*Flavor) error
}
