package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/wechat"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) WxLogin(ctx context.Context, code string) (LoginVO, error) {

	resp, err := wechat.GetWxResp(code)
	if err != nil {
		return LoginVO{}, err
	}

	openid := resp.OpenId

	user := &domain.User{
		OpenId:     openid,
		CreateTime: time.Now(),
	}

	err = svc.repo.Upsert(ctx, user)
	if err != nil {
		return LoginVO{}, err
	}

	token, err := middleware.GenerateToken(user.Id)
	if err != nil {
		return LoginVO{}, err
	}

	vo := LoginVO{
		Id:     user.Id,
		OpenId: openid,
		Token:  token,
	}
	return vo, nil

}
