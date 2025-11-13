package domian

import "context"

type ShopRepository interface {
	Set(context.Context, int) error
	Get(context.Context) (int, error)
}
