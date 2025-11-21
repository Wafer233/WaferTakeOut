package rpc

import (
	"context"

	shoppb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/shop"
)

type ShopService struct {
	client shoppb.ShopServiceClient
}

func NewShopService(client shoppb.ShopServiceClient) *ShopService {
	return &ShopService{client: client}
}

func (svc *ShopService) UpdateStatus(ctx context.Context, status int) error {

	req := shoppb.StatusMessage{Status: int32(status)}
	_, err := svc.client.UpdateStatus(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ShopService) FindStatus(ctx context.Context) (int, error) {

	req := shoppb.EmptyMessage{}
	resp, err := svc.client.FindStatus(ctx, &req)
	if err != nil {
		return 0, err
	}
	return int(resp.Status), nil
}
