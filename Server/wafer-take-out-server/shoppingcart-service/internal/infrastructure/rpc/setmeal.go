package rpc

import (
	"context"

	setmealpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/proto/setmeal"
)

type SetMealService struct {
	client setmealpb.SetmealServiceClient
}

func NewSetMealService(client setmealpb.SetmealServiceClient) *SetMealService {
	return &SetMealService{client: client}
}

func (svc *SetMealService) FindDetailById(ctx context.Context, id int64) (string, string, float64, error) {

	req := setmealpb.IdMessage{Id: id}
	resp, err := svc.client.FindDetailById(ctx, &req)
	if err != nil {
		return "", "", 0, err
	}
	return resp.Image, resp.Name, float64(resp.Price), nil
}
