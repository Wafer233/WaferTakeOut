package domain

import "context"

type ShoppingCartRepository interface {
	Find(context.Context, int64, int64, int64) (*ShoppingCart, error)
	UpdateNumber(context.Context, int64, int) error
	Create(ctx context.Context, cart *ShoppingCart) error
}
