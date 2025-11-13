package domian

import "context"

type ShopCache interface {
	Set(context.Context, int) error
	Get(context.Context) (int, error)
}
