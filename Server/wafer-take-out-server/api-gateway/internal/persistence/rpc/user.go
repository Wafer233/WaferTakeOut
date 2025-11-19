package rpc

import (
	"context"

	userApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/user"
	userpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/user"
	"github.com/jinzhu/copier"
)

type UserService struct {
	client userpb.UserServiceClient
}

func NewUserService(client userpb.UserServiceClient) *UserService {
	return &UserService{client: client}
}

func (svc *UserService) WxLogin(ctx context.Context,
	code string) (userApp.LoginVO, error) {

	req := userpb.LoginRequest{
		Code: code,
	}
	resp, err := svc.client.Login(ctx, &req)
	if err != nil {
		return userApp.LoginVO{}, err
	}

	vo := userApp.LoginVO{}
	_ = copier.Copy(&vo, resp)
	return vo, nil
}
