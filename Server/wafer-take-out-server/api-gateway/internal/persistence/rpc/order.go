package rpc

import (
	"context"

	orderApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/order"
	orderpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/order"
	"github.com/jinzhu/copier"
)

type OrderService struct {
	client orderpb.OrderServiceClient
}

func NewOrderService(client orderpb.OrderServiceClient) *OrderService {
	return &OrderService{client: client}
}

func (svc *OrderService) Submit(ctx context.Context,
	dto *orderApp.SubmitDTO, curId int64) (orderApp.SubmitVO, error) {

	req := &orderpb.SubmitRequest{}
	_ = copier.Copy(req, dto)
	req.CurId = curId
	resp, err := svc.client.Submit(ctx, req)
	if err != nil {
		return orderApp.SubmitVO{}, err
	}
	vo := orderApp.SubmitVO{}
	err = copier.Copy(&vo, resp)
	return vo, err
}

func (svc *OrderService) Payment(ctx context.Context,
	dto *orderApp.PaymentDTO) (orderApp.PaymentVO, error) {

	req := &orderpb.PaymentRequest{}
	_ = copier.Copy(req, dto)
	resp, err := svc.client.Payment(ctx, req)
	if err != nil {
		return orderApp.PaymentVO{}, err
	}
	vo := orderApp.PaymentVO{}
	err = copier.Copy(&vo, resp)
	return vo, err
}

func (svc *OrderService) Page(ctx context.Context, dto *orderApp.UserPageDTO,
	userId int64) (orderApp.UserPageVO, error) {
	req := &orderpb.UserPageRequest{}
	_ = copier.Copy(req, dto)
	req.CurId = userId
	resp, err := svc.client.Page(ctx, req)
	if err != nil {
		return orderApp.UserPageVO{}, err
	}
	vo := orderApp.UserPageVO{}
	_ = copier.Copy(&vo, resp)
	return vo, nil
}

func (svc *OrderService) FindPageAdmin(ctx context.Context,
	dto *orderApp.AdminPageDTO) (orderApp.AdminPageVO, error) {

	req := &orderpb.AdminPageRequest{}
	_ = copier.Copy(req, dto)
	resp, err := svc.client.FindPageAdmin(ctx, req)
	if err != nil {
		return orderApp.AdminPageVO{}, err
	}
	vo := orderApp.AdminPageVO{}
	err = copier.Copy(&vo, resp)
	return vo, err
}

func (svc *OrderService) GetOrder(ctx context.Context, orderId int64) (orderApp.UserOrderVO, error) {

	req := orderpb.IdMessage{Id: orderId}
	resp, err := svc.client.GetOrder(ctx, &req)
	if err != nil {
		return orderApp.UserOrderVO{}, err
	}
	vo := orderApp.UserOrderVO{}
	err = copier.Copy(&vo, resp)
	return vo, err
}

func (svc *OrderService) UserCancel(ctx context.Context, orderId int64) error {

	req := orderpb.IdMessage{Id: orderId}
	_, err := svc.client.UserCancel(ctx, &req)
	return err

}

func (svc *OrderService) CreateSame(ctx context.Context, orderId int64, curID int64) error {

	req := orderpb.CreateSameRequest{
		OrderId: orderId,
		CurId:   curID,
	}
	_, err := svc.client.CreateSame(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *OrderService) Confirm(ctx context.Context, orderId int64) error {
	req := orderpb.IdMessage{Id: orderId}
	_, err := svc.client.Confirm(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *OrderService) Rejection(ctx context.Context, dto *orderApp.RejectionDTO) error {
	req := &orderpb.RejectionRequest{}
	_ = copier.Copy(req, dto)
	_, err := svc.client.Rejection(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *OrderService) Cancel(ctx context.Context, dto *orderApp.CancelDTO) error {
	req := &orderpb.CancelRequest{}
	_ = copier.Copy(req, dto)
	_, err := svc.client.Cancel(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *OrderService) Delivery(ctx context.Context, orderId int64) error {
	req := orderpb.IdMessage{Id: orderId}
	_, err := svc.client.Delivery(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *OrderService) Complete(ctx context.Context, orderId int64) error {
	req := orderpb.IdMessage{Id: orderId}
	_, err := svc.client.Complete(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *OrderService) GetStatistics(ctx context.Context) (orderApp.StatisticsVO, error) {
	req := orderpb.EmptyMessage{}
	resp, err := svc.client.GetStatistics(ctx, &req)
	if err != nil {
		return orderApp.StatisticsVO{}, err
	}
	vo := orderApp.StatisticsVO{}
	err = copier.Copy(&vo, resp)
	return vo, err
}
