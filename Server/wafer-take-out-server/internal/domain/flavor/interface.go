package flavor

import "context"

type FlavorRepository interface {
	Insert(context.Context, []*Flavor) error
}
