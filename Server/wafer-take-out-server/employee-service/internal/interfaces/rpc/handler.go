package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/employee-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/employee-service/proto"
	"github.com/jinzhu/copier"
)

type EmployeeHandler struct {
	proto.UnimplementedEmployeeServiceServer
	svc *application.EmployeeService
}

func (h *EmployeeHandler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	dto := application.LoginDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.Login(ctx, &dto)
	if err != nil {
		return nil, err
	}

	resp := &proto.LoginResponse{}
	_ = copier.Copy(&resp, vo)
	return resp, nil
}

func (h *EmployeeHandler) ListPage(ctx context.Context, req *proto.PageRequest) (*proto.PageResponse, error) {

	dto := application.PageDTO{}
	_ = copier.Copy(&dto, req)
	vo, err := h.svc.FindPage(ctx, &dto)
	if err != nil {
		return nil, err
	}

	resp := &proto.PageResponse{}
	_ = copier.Copy(&resp, vo)
	return resp, nil
}

func (h *EmployeeHandler) FindById(ctx context.Context, req *proto.IdRequest) (*proto.Employee, error) {

	id := req.Id

	vo, err := h.svc.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := &proto.Employee{}

	_ = copier.Copy(resp, vo)
	return resp, nil
}

func (h *EmployeeHandler) UpdateStatus(ctx context.Context, req *proto.UpdateStatusRequest) (*proto.EmptyResponse, error) {

	err := h.svc.UpdateStatus(ctx, int(req.Status), req.Id, req.CurId)
	if err != nil {
		return nil, err
	}
	return &proto.EmptyResponse{}, nil
}

func (h *EmployeeHandler) Update(ctx context.Context, req *proto.AddEmployeeRequest) (*proto.EmptyResponse, error) {

	dto := application.AddEmployeeDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Update(ctx, &dto, req.CurId)

	if err != nil {
		return nil, err
	}
	return &proto.EmptyResponse{}, nil
}

func (h *EmployeeHandler) Create(ctx context.Context, req *proto.AddEmployeeRequest) (*proto.EmptyResponse, error) {

	dto := application.AddEmployeeDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Create(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}
	return &proto.EmptyResponse{}, nil
}
func (h *EmployeeHandler) UpdatePassword(ctx context.Context, req *proto.PasswordRequest) (*proto.EmptyResponse, error) {

	dto := application.PasswordDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.UpdatePassword(ctx, &dto, req.CurId)
	if err != nil {
		return nil, err
	}
	return &proto.EmptyResponse{}, nil
}

func NewEmployeeHandler(svc *application.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: svc}
}
