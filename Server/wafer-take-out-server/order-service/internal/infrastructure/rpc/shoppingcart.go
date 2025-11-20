package rpc

import (
	"context"

	shoppingcartpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/proto/shoppingcart"
	"github.com/jinzhu/copier"
)

type ShoppingCartService struct {
	client shoppingcartpb.ShoppingCartServiceClient
}

func NewShoppingCartService(client shoppingcartpb.ShoppingCartServiceClient) *ShoppingCartService {
	return &ShoppingCartService{client: client}
}

func (svc *ShoppingCartService) Add(ctx context.Context, dto *CartDTO, curId int64) error {
	req := shoppingcartpb.CartRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId

	_, err := svc.client.Add(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ShoppingCartService) FindByUserId(ctx context.Context, userId int64) ([]CartVO, error) {

	req := shoppingcartpb.IdMessage{Id: userId}
	resp, err := svc.client.FindByUserId(ctx, &req)
	if err != nil {
		return nil, err
	}
	messages := resp.Records

	vos := make([]CartVO, len(messages))
	_ = copier.Copy(&vos, &messages)
	return vos, nil
}

func (svc *ShoppingCartService) Delete(ctx context.Context, curId int64) error {
	_, err := svc.client.Delete(ctx, &shoppingcartpb.IdMessage{Id: curId})
	return err
}
