package domain

import "context"

type ShoppingCartRepository interface {
	Find(context.Context, int64, int64, int64) ([]*ShoppingCart, error)
	UpdateNumber(context.Context, int64, int) error
	Create(context.Context, *ShoppingCart) error
	Delete(context.Context, int64) error
}
