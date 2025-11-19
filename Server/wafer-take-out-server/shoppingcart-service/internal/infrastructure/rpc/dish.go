package rpc

import (
	"context"

	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/proto/dish"
)

type DishService struct {
	client dishpb.DishServiceClient
}

func NewDishService(client dishpb.DishServiceClient) *DishService {
	return &DishService{client: client}
}

func (svc *DishService) FindDetailById(ctx context.Context, id int64) (string, string, float64, error) {

	req := dishpb.IDMessage{Id: id}
	resp, err := svc.client.FindDetailById(ctx, &req)
	if err != nil {
		return "", "", 0, err
	}
	return resp.Image, resp.Name, float64(resp.Price), nil
}
