package domain

import "context"

type OrderRepository interface {
	Create(context.Context, *Order) error
	CreateDetail(context.Context, []*OrderDetail) error
	FindByNumber(context.Context, string) (*Order, error)
	UpdateStatus(context.Context, *Order) error
}
