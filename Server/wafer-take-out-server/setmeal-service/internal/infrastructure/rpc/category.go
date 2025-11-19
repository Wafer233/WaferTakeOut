package rpc

import (
	"context"

	categorypb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/proto/category"
)

type CategoryService struct {
	client categorypb.CategoryServiceClient
}

func NewCategoryService(client categorypb.CategoryServiceClient) *CategoryService {
	return &CategoryService{client: client}
}

func (svc *CategoryService) FindNameById(ctx context.Context, id int64) (string, error) {

	req := categorypb.IdRequest{Id: id}
	resp, err := svc.client.FindNameById(ctx, &req)
	if err != nil {
		return "", err
	}
	return resp.Name, nil
}
