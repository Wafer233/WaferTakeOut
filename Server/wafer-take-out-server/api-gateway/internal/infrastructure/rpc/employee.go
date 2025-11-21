package rpc

import (
	"context"

	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/employee"
	employeepb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/employee"
	"github.com/jinzhu/copier"
)

type EmployeeService struct {
	client employeepb.EmployeeServiceClient
}

func NewEmployeeService(client employeepb.EmployeeServiceClient) *EmployeeService {
	return &EmployeeService{client: client}
}

func (svc *EmployeeService) Login(ctx context.Context, dto *employeeApp.LoginDTO) (employeeApp.LoginVO, error) {

	req := employeepb.LoginRequest{}
	vo := employeeApp.LoginVO{}

	_ = copier.Copy(&req, dto)
	resp, err := svc.client.Login(ctx, &req)
	if err != nil {
		return employeeApp.LoginVO{}, err
	}

	_ = copier.Copy(&vo, resp)
	return vo, nil
}

func (svc *EmployeeService) FindPage(ctx context.Context, dto *employeeApp.PageDTO) (employeeApp.PageVO, error) {

	req := employeepb.PageRequest{}
	vo := employeeApp.PageVO{}
	_ = copier.Copy(&req, dto)
	resp, err := svc.client.ListPage(ctx, &req)
	if err != nil {
		return employeeApp.PageVO{}, err
	}

	_ = copier.Copy(&vo, resp)
	return vo, nil

}

func (svc *EmployeeService) FindById(ctx context.Context, id int64) (employeeApp.Employee, error) {

	req := employeepb.IdRequest{
		Id: id,
	}
	vo := employeeApp.Employee{}

	resp, err := svc.client.FindById(ctx, &req)
	if err != nil {
		return employeeApp.Employee{}, err
	}
	_ = copier.Copy(&vo, resp)
	return vo, nil
}

func (svc *EmployeeService) UpdateStatus(ctx context.Context, status int, id int64, curId int64) error {

	req := employeepb.UpdateStatusRequest{
		Id:     id,
		Status: int32(status),
		CurId:  curId,
	}

	_, err := svc.client.UpdateStatus(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *EmployeeService) Update(ctx context.Context, dto *employeeApp.AddEmployeeDTO, curId int64) error {
	req := employeepb.AddEmployeeRequest{}

	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.Update(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *EmployeeService) Create(ctx context.Context, dto *employeeApp.AddEmployeeDTO, curId int64) error {
	req := employeepb.AddEmployeeRequest{}

	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.Create(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *EmployeeService) UpdatePassword(ctx context.Context, dto *employeeApp.PasswordDTO, curId int64) error {
	req := employeepb.PasswordRequest{}
	_ = copier.Copy(&req, dto)
	req.CurId = curId
	_, err := svc.client.UpdatePassword(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}
