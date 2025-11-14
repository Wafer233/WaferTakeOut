package domain

import "context"

type ShoppingCartRepository interface {
	Create(ctx context.Context, cart *ShoppingCart) error
}
