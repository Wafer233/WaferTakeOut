package flavor

import "context"

type FlavorRepository interface {
	Insert(context.Context, []*Flavor) error
	GetsByDishId(context.Context, int64) ([]*Flavor, error)
}
