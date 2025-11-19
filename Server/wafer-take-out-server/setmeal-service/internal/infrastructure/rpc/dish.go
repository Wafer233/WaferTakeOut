package rpc

import (
	"context"

	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/proto/dish"
)

type DishService struct {
	client dishpb.DishServiceClient
}

func NewDishService(client dishpb.DishServiceClient) *DishService {
	return &DishService{client: client}
}

func (svc *DishService) FindDescriptionById(ctx context.Context,
	id int64) (string, string, error) {
	req := dishpb.IDMessage{Id: id}
	resp, err := svc.client.FindDescriptionById(ctx, &req)
	if err != nil {
		return "", "", err
	}

	return resp.Description, resp.Image, nil
}
