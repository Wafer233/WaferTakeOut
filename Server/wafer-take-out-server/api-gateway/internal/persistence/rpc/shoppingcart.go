package rpc

import (
	"context"

	shoppingcartApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/shoppingcart"
	shoppingcartpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/shoppingcart"
	"github.com/jinzhu/copier"
)

type ShoppingCartService struct {
	client shoppingcartpb.ShoppingCartServiceClient
}

func NewShoppingCartService(client shoppingcartpb.ShoppingCartServiceClient) *ShoppingCartService {
	return &ShoppingCartService{
		client: client,
	}
}

func (svc *ShoppingCartService) Add(ctx context.Context, dto *shoppingcartApp.CartDTO, curId int64) error {

	req := shoppingcartpb.CartRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.Add(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ShoppingCartService) FindByUserId(ctx context.Context,
	userId int64) ([]shoppingcartApp.RecordVO, error) {

	req := &shoppingcartpb.IdMessage{
		Id: userId,
	}

	resp, err := svc.client.FindByUserId(ctx, req)

	if err != nil {
		return nil, err
	}

	records := resp.Records
	vos := make([]shoppingcartApp.RecordVO, len(records))
	_ = copier.Copy(&vos, &records)
	return vos, nil
}

func (svc *ShoppingCartService) Sub(ctx context.Context, dto *shoppingcartApp.CartDTO, curId int64) error {

	req := shoppingcartpb.CartRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.Sub(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ShoppingCartService) Delete(ctx context.Context, curId int64) error {

	_, err := svc.client.Delete(ctx, &shoppingcartpb.IdMessage{Id: curId})
	return err

}
