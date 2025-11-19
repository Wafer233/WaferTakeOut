package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/internal/application"
	setmealpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/proto"
	"github.com/jinzhu/copier"
)

type SetMealHandler struct {
	svc *application.SetMealService
	setmealpb.UnimplementedSetmealServiceServer
}

func NewSetMealHandler(svc *application.SetMealService) *SetMealHandler {
	return &SetMealHandler{svc: svc}
}

func (h *SetMealHandler) Create(ctx context.Context,
	req *setmealpb.SetMealRequest) (*setmealpb.EmptyRequest, error) {

	dto := application.SetMealDTO{}
	_ = copier.Copy(&dto, req)
	curId := req.CurId
	err := h.svc.Create(ctx, &dto, curId)
	if err != nil {
		return nil, err
	}
	return &setmealpb.EmptyRequest{}, nil
}

func (h *SetMealHandler) Deletes(ctx context.Context,
	req *setmealpb.IdsMessage) (*setmealpb.EmptyRequest, error) {

	ids := req.Ids
	err := h.svc.Deletes(ctx, ids)
	if err != nil {
		return nil, err
	}
	return &setmealpb.EmptyRequest{}, nil
}

func (h *SetMealHandler) Update(ctx context.Context,
	req *setmealpb.SetMealRequest) (*setmealpb.EmptyRequest, error) {
	dto := application.SetMealDTO{}
	_ = copier.Copy(&dto, req)
	curId := req.CurId

	err := h.svc.Update(ctx, &dto, curId)
	if err != nil {
		return nil, err
	}
	return &setmealpb.EmptyRequest{}, nil
}

func (h *SetMealHandler) UpdateStatus(ctx context.Context,
	req *setmealpb.StatusRequest) (*setmealpb.EmptyRequest, error) {

	err := h.svc.UpdateStatus(ctx, req.Id, req.CurId, int(req.Status))
	if err != nil {
		return nil, err
	}
	return &setmealpb.EmptyRequest{}, nil
}

func (h *SetMealHandler) FindById(ctx context.Context,
	req *setmealpb.IdMessage) (*setmealpb.SetMealResponse, error) {

	vo, err := h.svc.FindById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := setmealpb.SetMealResponse{}
	_ = copier.Copy(&resp, &vo)
	return &resp, nil

}

func (h *SetMealHandler) FindPage(ctx context.Context,
	req *setmealpb.PageRequest) (*setmealpb.PageResponse, error) {

	dto := application.PageDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.FindPage(ctx, &dto)
	if err != nil {
		return nil, err
	}
	resp := setmealpb.PageResponse{}
	_ = copier.Copy(&resp, &vo)
	return &resp, nil
}

func (h *SetMealHandler) FindByCategoryId(ctx context.Context,
	req *setmealpb.IdMessage) (*setmealpb.FindByCategoryResponse, error) {

	vo, err := h.svc.FindByCategoryId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	setmeals := make([]*setmealpb.SetMealMessage, len(vo))
	_ = copier.Copy(&setmeals, &vo)
	resp := setmealpb.FindByCategoryResponse{
		Setmeals: setmeals,
	}
	return &resp, nil
}

func (h *SetMealHandler) FindDishById(ctx context.Context,
	req *setmealpb.IdMessage) (*setmealpb.FindDishByIdResponse, error) {

	vo, err := h.svc.FindDishById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	dishes := make([]*setmealpb.DishMessage, len(vo))
	_ = copier.Copy(&dishes, &vo)
	resp := setmealpb.FindDishByIdResponse{
		Dishes: dishes,
	}
	return &resp, nil
}

func (h *SetMealHandler) FindDetailById(ctx context.Context,
	req *setmealpb.IdMessage) (*setmealpb.ShoppingCartResponse, error) {

	name, img, price, err := h.svc.FindDetailById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := setmealpb.ShoppingCartResponse{
		Name:  name,
		Image: img,
		Price: float32(price),
	}
	return &resp, nil
}
