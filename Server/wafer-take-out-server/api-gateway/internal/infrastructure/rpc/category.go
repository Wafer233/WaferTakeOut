package rpc

import (
	"context"

	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/category"
	categorypb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/category"
	"github.com/jinzhu/copier"
)

type CategoryService struct {
	client categorypb.CategoryServiceClient
}

func NewCategoryService(client categorypb.CategoryServiceClient) *CategoryService {
	return &CategoryService{client: client}
}

func (svc *CategoryService) Create(ctx context.Context, dto *categoryApp.CategoryDTO,
	curId int64) error {
	req := categorypb.CategoryRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.Create(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *CategoryService) Delete(ctx context.Context, id int64) error {
	req := categorypb.IdRequest{Id: id}
	_, err := svc.client.Delete(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *CategoryService) Update(ctx context.Context,
	dto *categoryApp.CategoryDTO, curId int64) error {

	req := categorypb.CategoryRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.Update(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}
func (svc *CategoryService) UpdateStatus(ctx context.Context, id int64, status int, curId int64) error {
	dto := categorypb.UpdateStatusRequest{
		CurId:  curId,
		Id:     id,
		Status: int32(status),
	}

	_, err := svc.client.UpdateStatus(ctx, &dto)
	if err != nil {
		return err
	}
	return nil
}

func (svc *CategoryService) FindPage(ctx context.Context,
	dto *categoryApp.PageDTO) (categoryApp.PageVO, error) {

	req := categorypb.PageRequest{}
	_ = copier.Copy(&req, dto)
	resp, err := svc.client.FindPage(ctx, &req)
	if err != nil {
		return categoryApp.PageVO{}, err
	}
	vo := categoryApp.PageVO{}
	err = copier.Copy(&vo, resp)
	if err != nil {
		return categoryApp.PageVO{}, err
	}
	return vo, nil
}

func (svc *CategoryService) FindByType(ctx context.Context, curType int) ([]categoryApp.Record, error) {

	req := categorypb.FindTypeRequest{
		CurType: int32(curType),
	}
	resp, err := svc.client.FindByType(ctx, &req)
	if err != nil {
		return nil, err
	}
	records := resp.Records
	vo := make([]categoryApp.Record, len(records))

	_ = copier.Copy(&vo, &records)
	return vo, nil
}
