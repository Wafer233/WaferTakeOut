package domian

import "context"

type ShopRepository interface {
	Update(context.Context, int) error
	Find(context.Context) (int, error)
}
