package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/category-service/internal/application"
	categorypb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/category-service/proto"
	"github.com/jinzhu/copier"
)

type CategoryHandler struct {
	categorypb.UnimplementedCategoryServiceServer
	svc *application.CategoryService
}

func NewCategoryHandler(svc *application.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}

func (h *CategoryHandler) Create(ctx context.Context,
	req *categorypb.CategoryRequest) (*categorypb.EmptyResponse, error) {

	dto := application.CategoryDTO{}
	_ = copier.Copy(&dto, req)
	curId := req.CurId
	err := h.svc.Create(ctx, &dto, curId)
	if err != nil {
		return nil, err
	}
	return &categorypb.EmptyResponse{}, nil
}

func (h *CategoryHandler) Delete(ctx context.Context,
	req *categorypb.IdRequest) (*categorypb.EmptyResponse, error) {

	id := req.Id
	err := h.svc.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &categorypb.EmptyResponse{}, nil
}

func (h *CategoryHandler) Update(ctx context.Context,
	req *categorypb.CategoryRequest) (*categorypb.EmptyResponse, error) {

	dto := application.CategoryDTO{}
	_ = copier.Copy(&dto, req)
	curId := req.CurId
	err := h.svc.Update(ctx, &dto, curId)
	if err != nil {
		return nil, err
	}
	return &categorypb.EmptyResponse{}, nil
}

func (h *CategoryHandler) UpdateStatus(ctx context.Context,
	req *categorypb.UpdateStatusRequest) (*categorypb.EmptyResponse, error) {

	err := h.svc.UpdateStatus(ctx, req.Id, int(req.Status), req.CurId)
	if err != nil {
		return nil, err
	}
	return &categorypb.EmptyResponse{}, nil
}

func (h *CategoryHandler) FindPage(ctx context.Context, req *categorypb.PageRequest) (*categorypb.PageResponse, error) {

	dto := application.PageDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.FindPage(ctx, &dto)
	if err != nil {
		return nil, err
	}

	resp := &categorypb.PageResponse{}
	_ = copier.Copy(resp, &vo)
	return resp, nil
}
func (h *CategoryHandler) FindByType(ctx context.Context,
	req *categorypb.FindTypeRequest) (*categorypb.FindTypeResponse, error) {

	vo, err := h.svc.FindByType(ctx, int(req.CurType))
	if err != nil {
		return nil, err
	}
	resp := &categorypb.FindTypeResponse{}
	_ = copier.Copy(resp, &vo)
	return resp, nil
}

// 这个是给dish用的
func (h *CategoryHandler) FindNameById(ctx context.Context,
	req *categorypb.IdRequest) (*categorypb.NameResponse, error) {

	name, err := h.svc.FindNameById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := &categorypb.NameResponse{
		Name: name,
	}
	return resp, nil
}
