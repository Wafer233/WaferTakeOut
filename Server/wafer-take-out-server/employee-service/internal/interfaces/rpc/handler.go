package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/employee-service/internal/application"
	employeepb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/employee-service/proto"
	"github.com/jinzhu/copier"
)

type EmployeeHandler struct {
	employeepb.UnimplementedEmployeeServiceServer
	svc *application.EmployeeService
}

func (h *EmployeeHandler) Login(ctx context.Context, req *employeepb.LoginRequest) (*employeepb.LoginResponse, error) {

	dto := application.LoginDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.Login(ctx, &dto)
	if err != nil {
		return nil, err
	}

	resp := &employeepb.LoginResponse{}
	_ = copier.Copy(&resp, vo)
	return resp, nil
}

func (h *EmployeeHandler) ListPage(ctx context.Context, req *employeepb.PageRequest) (*employeepb.PageResponse, error) {

	dto := application.PageDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.FindPage(ctx, &dto)
	if err != nil {
		return nil, err
	}

	resp := &employeepb.PageResponse{}
	_ = copier.Copy(&resp, vo)
	return resp, nil
}

func (h *EmployeeHandler) FindById(ctx context.Context, req *employeepb.IdRequest) (*employeepb.Employee, error) {

	id := req.Id

	vo, err := h.svc.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := &employeepb.Employee{}

	_ = copier.Copy(resp, vo)
	return resp, nil
}

func (h *EmployeeHandler) UpdateStatus(ctx context.Context, req *employeepb.UpdateStatusRequest) (*employeepb.EmptyResponse, error) {

	err := h.svc.UpdateStatus(ctx, int(req.Status), req.Id, req.CurId)
	if err != nil {
		return nil, err
	}
	return &employeepb.EmptyResponse{}, nil
}

func (h *EmployeeHandler) Update(ctx context.Context, req *employeepb.AddEmployeeRequest) (*employeepb.EmptyResponse, error) {

	dto := application.AddEmployeeDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Update(ctx, &dto, req.CurId)

	if err != nil {
		return nil, err
	}
	return &employeepb.EmptyResponse{}, nil
}

func (h *EmployeeHandler) Create(ctx context.Context, req *employeepb.AddEmployeeRequest) (*employeepb.EmptyResponse, error) {

	dto := application.AddEmployeeDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Create(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}
	return &employeepb.EmptyResponse{}, nil
}
func (h *EmployeeHandler) UpdatePassword(ctx context.Context, req *employeepb.PasswordRequest) (*employeepb.EmptyResponse, error) {

	dto := application.PasswordDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.UpdatePassword(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}
	return &employeepb.EmptyResponse{}, nil
}

func NewEmployeeHandler(svc *application.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: svc}
}
