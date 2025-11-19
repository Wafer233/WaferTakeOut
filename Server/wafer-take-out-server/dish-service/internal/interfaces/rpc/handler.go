package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/application"
	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/proto"
	"github.com/jinzhu/copier"
)

type DishHandler struct {
	dishpb.UnimplementedDishServiceServer
	svc *application.DishService
}

func NewDishHandler(svc *application.DishService) *DishHandler {
	return &DishHandler{svc: svc}
}

func (h *DishHandler) Create(ctx context.Context,
	req *dishpb.DishRequest) (*dishpb.EmptyResponse, error) {

	dto := application.DishDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Create(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}
	return &dishpb.EmptyResponse{}, nil
}

func (h *DishHandler) Update(ctx context.Context,
	req *dishpb.DishRequest) (*dishpb.EmptyResponse, error) {
	dto := application.DishDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Update(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}
	return &dishpb.EmptyResponse{}, nil
}

func (h *DishHandler) FindByPage(ctx context.Context,
	req *dishpb.PageRequest) (*dishpb.PageResponse, error) {

	dto := application.PageDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.FindPage(ctx, &dto)
	if err != nil {
		return nil, err
	}
	resp := dishpb.PageResponse{}
	_ = copier.Copy(&resp, &vo)
	return &resp, nil
}

func (h *DishHandler) FindByCategoryId(ctx context.Context,
	req *dishpb.IDMessage) (*dishpb.RecordsResponse, error) {

	categoryId := req.Id

	vo, err := h.svc.FindByCategoryId(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	records := make([]*dishpb.Record, len(vo))
	_ = copier.Copy(&records, &vo)

	return &dishpb.RecordsResponse{
		Records: records,
	}, nil
}

func (h *DishHandler) FindById(ctx context.Context,
	req *dishpb.IDMessage) (*dishpb.DishResponse, error) {

	vo, err := h.svc.FindById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := dishpb.DishResponse{}
	_ = copier.Copy(&resp, &vo)
	return &resp, nil
}

func (h *DishHandler) UpdateStatus(ctx context.Context,
	req *dishpb.UpdateRequest) (*dishpb.EmptyResponse, error) {

	err := h.svc.UpdateStatus(ctx, req.Id, int(req.Status), req.CurId)
	if err != nil {
		return nil, err
	}
	return &dishpb.EmptyResponse{}, nil
}

func (h *DishHandler) Delete(ctx context.Context,
	req *dishpb.IdsMessage) (*dishpb.EmptyResponse, error) {

	err := h.svc.Delete(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	return &dishpb.EmptyResponse{}, nil
}

func (h *DishHandler) FindByCategoryIdFlavor(ctx context.Context,
	req *dishpb.IDMessage) (*dishpb.DishesResponse, error) {

	vo, err := h.svc.FindByCategoryIdFlavor(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	dishes := make([]*dishpb.DishResponse, len(vo))
	_ = copier.Copy(&dishes, &vo)

	return &dishpb.DishesResponse{
		Dishes: dishes,
	}, nil
}

// 给setmeal用的
func (h *DishHandler) FindDescriptionById(ctx context.Context,
	req *dishpb.IDMessage) (*dishpb.DescriptionResponse, error) {

	des, img, err := h.svc.FindDescriptionById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp := dishpb.DescriptionResponse{
		Description: des,
		Image:       img,
	}
	return &resp, nil
}
