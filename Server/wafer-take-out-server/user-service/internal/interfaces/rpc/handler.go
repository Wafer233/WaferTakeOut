package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/user-service/internal/application"
	userpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/user-service/proto"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	svc *application.UserService
	userpb.UnimplementedUserServiceServer
}

func NewUserHandler(svc *application.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Login(ctx context.Context,
	req *userpb.LoginRequest) (*userpb.LoginResponse, error) {

	vo, err := h.svc.WxLogin(ctx, req.Code)
	if err != nil {
		return nil, err
	}
	resp := userpb.LoginResponse{}
	_ = copier.Copy(&resp, &vo)
	return &resp, nil

}
