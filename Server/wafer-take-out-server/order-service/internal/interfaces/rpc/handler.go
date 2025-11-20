package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/application"
	orderpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/proto"
	"github.com/jinzhu/copier"
)

type OrderHandler struct {
	svc *application.OrderService
	orderpb.UnimplementedOrderServiceServer
}

func NewOrderHandler(svc *application.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) Submit(ctx context.Context,
	req *orderpb.SubmitRequest) (*orderpb.SubmitResponse, error) {

	dto := application.SubmitDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.Submit(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}
	resp := &orderpb.SubmitResponse{}
	_ = copier.Copy(&resp, &vo)
	return resp, nil

}

func (h *OrderHandler) Payment(ctx context.Context,
	req *orderpb.PaymentRequest) (*orderpb.PaymentResponse, error) {

	dto := application.PaymentDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.Payment(ctx, &dto)
	if err != nil {
		return nil, err
	}
	resp := &orderpb.PaymentResponse{}
	_ = copier.Copy(&resp, &vo)
	return resp, nil
}

func (h *OrderHandler) Page(ctx context.Context,
	req *orderpb.UserPageRequest) (*orderpb.UserPageResponse, error) {

	dto := application.UserPageDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.Page(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}
	resp := &orderpb.UserPageResponse{}
	_ = copier.Copy(&resp, &vo)
	return resp, nil
}

func (h *OrderHandler) FindPageAdmin(ctx context.Context,
	req *orderpb.AdminPageRequest) (*orderpb.AdminPageResponse, error) {

	dto := application.AdminPageDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.FindPageAdmin(ctx, &dto)
	if err != nil {
		return nil, err
	}
	resp := &orderpb.AdminPageResponse{}
	_ = copier.Copy(&resp, &vo)
	return resp, nil
}

func (h *OrderHandler) GetOrder(ctx context.Context,
	req *orderpb.IdMessage) (*orderpb.UserOrderResponse, error) {

	vo, err := h.svc.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := &orderpb.UserOrderResponse{}
	_ = copier.Copy(&resp, &vo)
	return resp, nil
}

func (h *OrderHandler) UserCancel(ctx context.Context,
	req *orderpb.IdMessage) (*orderpb.EmptyMessage, error) {

	err := h.svc.UserCancel(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &orderpb.EmptyMessage{}, nil
}

func (h *OrderHandler) CreateSame(ctx context.Context,
	req *orderpb.CreateSameRequest) (*orderpb.EmptyMessage, error) {

	err := h.svc.CreateSame(ctx, req.OrderId, req.CurId)
	if err != nil {
		return nil, err
	}
	return &orderpb.EmptyMessage{}, nil
}

func (h *OrderHandler) Confirm(ctx context.Context,
	req *orderpb.IdMessage) (*orderpb.EmptyMessage, error) {
	err := h.svc.Confirm(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &orderpb.EmptyMessage{}, nil
}

func (h *OrderHandler) Rejection(ctx context.Context,
	req *orderpb.RejectionRequest) (*orderpb.EmptyMessage, error) {

	dto := application.RejectionDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Rejection(ctx, &dto)
	if err != nil {
		return nil, err
	}
	return &orderpb.EmptyMessage{}, nil
}

func (h *OrderHandler) Cancel(ctx context.Context,
	req *orderpb.CancelRequest) (*orderpb.EmptyMessage, error) {

	dto := application.CancelDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Cancel(ctx, &dto)
	if err != nil {
		return nil, err
	}
	return &orderpb.EmptyMessage{}, nil
}

func (h *OrderHandler) Delivery(ctx context.Context,
	req *orderpb.IdMessage) (*orderpb.EmptyMessage, error) {

	err := h.svc.Delivery(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &orderpb.EmptyMessage{}, nil

}

func (h *OrderHandler) Complete(ctx context.Context,
	req *orderpb.IdMessage) (*orderpb.EmptyMessage, error) {

	err := h.svc.Complete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &orderpb.EmptyMessage{}, nil
}

func (h *OrderHandler) GetStatistics(ctx context.Context, req *orderpb.EmptyMessage) (*orderpb.StatisticsResponse, error) {

	vo, err := h.svc.GetStatistics(ctx)
	if err != nil {
		return nil, err
	}
	resp := &orderpb.StatisticsResponse{}
	_ = copier.Copy(&resp, &vo)
	return resp, nil
}
