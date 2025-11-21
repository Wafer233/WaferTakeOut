package rpc

import (
	"context"

	setmealApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/setmeal"
	setmealpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/setmeal"
	"github.com/jinzhu/copier"
)

type SetMealService struct {
	client setmealpb.SetmealServiceClient
}

func NewSetMealService(client setmealpb.SetmealServiceClient) *SetMealService {
	return &SetMealService{client: client}
}

func (svc *SetMealService) Create(ctx context.Context, dto *setmealApp.SetMealDTO, curId int64) error {

	req := setmealpb.SetMealRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId

	_, err := svc.client.Create(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) Deletes(ctx context.Context, ids []int64) error {

	req := setmealpb.IdsMessage{
		Ids: ids,
	}
	_, err := svc.client.Deletes(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) Update(ctx context.Context, dto *setmealApp.SetMealDTO, curId int64) error {

	req := setmealpb.SetMealRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.Update(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) UpdateStatus(ctx context.Context, id int64,
	curId int64, status int) error {
	req := setmealpb.StatusRequest{
		Id:     id,
		CurId:  curId,
		Status: int32(status),
	}
	_, err := svc.client.UpdateStatus(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) FindById(ctx context.Context, id int64) (setmealApp.SetMealVO, error) {

	req := setmealpb.IdMessage{Id: id}
	resp, err := svc.client.FindById(ctx, &req)
	if err != nil {
		return setmealApp.SetMealVO{}, err
	}
	vo := setmealApp.SetMealVO{}
	_ = copier.Copy(&vo, resp)
	return vo, nil
}

func (svc *SetMealService) FindPage(ctx context.Context, dto *setmealApp.PageDTO) (setmealApp.PageVO, error) {

	req := setmealpb.PageRequest{}
	_ = copier.Copy(&req, dto)
	resp, err := svc.client.FindPage(ctx, &req)
	if err != nil {
		return setmealApp.PageVO{}, err
	}
	vo := setmealApp.PageVO{}
	_ = copier.Copy(&vo, resp)
	return vo, nil
}

func (svc *SetMealService) FindByCategoryId(ctx context.Context, cid int64) ([]setmealApp.FindByCategoryVO, error) {

	req := setmealpb.IdMessage{
		Id: cid,
	}
	resp, err := svc.client.FindByCategoryId(ctx, &req)
	if err != nil {
		return []setmealApp.FindByCategoryVO{}, err
	}
	message := resp.Setmeals
	vos := make([]setmealApp.FindByCategoryVO, len(message))
	_ = copier.Copy(&vos, &message)
	return vos, nil

}

func (svc *SetMealService) FindDishById(ctx context.Context, setId int64) ([]setmealApp.DishVO, error) {

	req := setmealpb.IdMessage{
		Id: setId,
	}
	resp, err := svc.client.FindDishById(ctx, &req)
	if err != nil {
		return []setmealApp.DishVO{}, err
	}
	message := resp.Dishes
	vos := make([]setmealApp.DishVO, len(message))
	_ = copier.Copy(&vos, &message)
	return vos, nil
}
