package domain

import "context"

type DishRepository interface {
	//UpdateById(context.Context, *Dish) error
	DeletesById(context.Context, []int64) error
	Insert(context.Context, *Dish) error
	GetById(context.Context, int64) (*Dish, error)
	GetsByCategoryId(context.Context, int64) ([]*Dish, error)
	GetsPaged(context.Context, string, int64, int, int, int) ([]*Dish, int64, error)
	UpdateStatusById(context.Context, *Dish) error
	UpdateById(context.Context, *Dish) error
}
