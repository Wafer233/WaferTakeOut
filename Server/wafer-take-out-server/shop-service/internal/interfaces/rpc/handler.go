package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shop-service/internal/application"
	shoppb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shop-service/proto"
)

type ShopHandler struct {
	svc *application.ShopService
	shoppb.UnimplementedShopServiceServer
}

func NewShopHandler(svc *application.ShopService) *ShopHandler {
	return &ShopHandler{svc: svc}
}

func (h *ShopHandler) UpdateStatus(ctx context.Context,
	req *shoppb.StatusMessage) (*shoppb.EmptyMessage, error) {

	err := h.svc.UpdateStatus(ctx, int(req.Status))
	if err != nil {
		return nil, err
	}
	return &shoppb.EmptyMessage{}, nil
}

func (h *ShopHandler) FindStatus(ctx context.Context,
	req *shoppb.EmptyMessage) (*shoppb.StatusMessage, error) {

	sta, err := h.svc.FindStatus(ctx)
	if err != nil {
		return nil, err
	}
	resp := shoppb.StatusMessage{
		Status: int32(sta),
	}
	return &resp, nil
}
