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
	FindByIds(context.Context, []int64) (map[int64]string, map[int64]string, error)
	FindByCategoryIdFlavor(context.Context, int64) ([]*Dish, map[int64][]*Flavor, error)
	FindDescriptionById(context.Context, int64) (string, string, error)

	//shoppingcart 用的
	FindDetailById(context.Context, int64) (string, string, float64, error)
}
