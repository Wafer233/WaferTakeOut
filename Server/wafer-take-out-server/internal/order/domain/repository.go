package domain

import "context"

type OrderRepository interface {
	Create(context.Context, *Order) error
	CreateDetail(context.Context, []*OrderDetail) error
	FindByNumber(context.Context, string) (*Order, error)
	UpdateStatus(context.Context, *Order) error
	FindPage(context.Context, int, int, int64, int) ([]*Order, int64, error)
	FindDetailByOrderIds(context.Context, []int64) (map[int64][]*OrderDetail, error)
	FindById(context.Context, int64) (*Order, error)
	FindDetailByOrderId(context.Context, int64) ([]*OrderDetail, error)
}
