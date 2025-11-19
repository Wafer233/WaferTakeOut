package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/internal/application"
	shoppingcartpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/proto"
	"github.com/jinzhu/copier"
)

type ShoppingCartHandler struct {
	svc *application.ShoppingCartService
	shoppingcartpb.UnimplementedShoppingCartServiceServer
}

func NewShoppingCartHandler(svc *application.ShoppingCartService) *ShoppingCartHandler {
	return &ShoppingCartHandler{svc: svc}
}

func (h *ShoppingCartHandler) Add(ctx context.Context,
	req *shoppingcartpb.CartRequest) (*shoppingcartpb.EmptyMessage, error) {

	dto := application.CartDTO{}
	_ = copier.Copy(&dto, req)

	err := h.svc.Add(ctx, &dto, req.CurId)

	if err != nil {
		return nil, err
	}
	return &shoppingcartpb.EmptyMessage{}, nil
}

func (h *ShoppingCartHandler) FindByUserId(ctx context.Context,
	req *shoppingcartpb.IdMessage) (*shoppingcartpb.CartResponse, error) {

	vo, err := h.svc.FindByUserId(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	records := make([]*shoppingcartpb.CartMessage, len(vo))
	_ = copier.Copy(&records, &vo)
	resp := &shoppingcartpb.CartResponse{Records: records}
	return resp, nil
}

func (h *ShoppingCartHandler) Sub(ctx context.Context,
	req *shoppingcartpb.CartRequest) (*shoppingcartpb.EmptyMessage, error) {

	dto := application.CartDTO{}
	_ = copier.Copy(&dto, req)

	err := h.svc.Sub(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}

	return &shoppingcartpb.EmptyMessage{}, nil
}

func (h *ShoppingCartHandler) Delete(ctx context.Context,
	req *shoppingcartpb.IdMessage) (*shoppingcartpb.EmptyMessage, error) {

	err := h.svc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &shoppingcartpb.EmptyMessage{}, nil
}
