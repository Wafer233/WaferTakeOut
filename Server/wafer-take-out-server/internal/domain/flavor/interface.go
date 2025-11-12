package flavor

import "context"

type FlavorRepository interface {
	Inserts(context.Context, []*Flavor) error
	GetsByDishId(context.Context, int64) ([]*Flavor, error)
	UpdatesByDishId(context.Context, []*Flavor) error
}
