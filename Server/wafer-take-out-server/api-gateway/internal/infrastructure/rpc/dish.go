package rpc

import (
	"context"

	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/dish"
	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/dish"
	"github.com/jinzhu/copier"
)

type DishService struct {
	client dishpb.DishServiceClient
}

func NewDishService(client dishpb.DishServiceClient) *DishService {
	return &DishService{client: client}
}

func (svc *DishService) Create(ctx context.Context,
	dto *dishApp.DishDTO, curId int64) error {

	req := dishpb.DishRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.Create(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *DishService) Update(ctx context.Context,
	dto *dishApp.DishDTO, curID int64) error {
	req := dishpb.DishRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curID
	_, err := svc.client.Update(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *DishService) FindPage(ctx context.Context,
	dto *dishApp.PageDTO) (dishApp.PageVO, error) {

	req := dishpb.PageRequest{}
	_ = copier.Copy(&req, dto)
	resp, err := svc.client.FindByPage(ctx, &req)
	if err != nil {
		return dishApp.PageVO{}, err
	}

	vo := dishApp.PageVO{}
	_ = copier.Copy(&vo, &resp)
	return vo, nil
}

func (svc *DishService) FindByCategoryId(ctx context.Context,
	cid int64) ([]dishApp.Record, error) {

	req := dishpb.IDMessage{Id: cid}
	resp, err := svc.client.FindByCategoryId(ctx, &req)
	if err != nil {
		return nil, err
	}
	records := resp.Records
	vos := make([]dishApp.Record, len(records))

	_ = copier.Copy(&vos, &records)
	return vos, nil
}

func (svc *DishService) FindById(ctx context.Context,
	id int64) (dishApp.DishVO, error) {

	resp, err := svc.client.FindById(ctx, &dishpb.IDMessage{Id: id})
	if err != nil {
		return dishApp.DishVO{}, err
	}

	vo := dishApp.DishVO{}
	_ = copier.Copy(&vo, &resp)
	return vo, nil

}

func (svc *DishService) UpdateStatus(ctx context.Context,
	id int64, status int, curId int64) error {

	req := dishpb.UpdateRequest{
		CurId:  curId,
		Id:     id,
		Status: int32(status),
	}

	_, err := svc.client.UpdateStatus(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *DishService) Delete(ctx context.Context, idArr []int64) error {

	req := dishpb.IdsMessage{Ids: idArr}
	_, err := svc.client.Delete(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *DishService) FindByCategoryIdFlavor(ctx context.Context,
	categoryId int64) ([]dishApp.DishVO, error) {

	req := dishpb.IDMessage{Id: categoryId}
	resp, err := svc.client.FindByCategoryIdFlavor(ctx, &req)
	if err != nil {
		return nil, err
	}
	dishes := resp.Dishes
	vos := make([]dishApp.DishVO, len(dishes))
	_ = copier.Copy(&vos, &dishes)
	return vos, nil
}
