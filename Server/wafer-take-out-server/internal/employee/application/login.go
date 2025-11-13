package application

import (
	"context"
	"errors"
)

type LoginDTO struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginVO struct {
	ID       int64  `json:"id"`       // 主键值
	UserName string `json:"userName"` // 用户名
	Name     string `json:"name"`     // 姓名
	Token    string `json:"token"`    // JWT令牌
}

func (svc *EmployeeService) Login(ctx context.Context, dto *LoginDTO) (*LoginVO, error) {

	entity, err := svc.repo.GetByUsername(ctx, dto.Username)
	if err != nil {
		return nil, err
	}
	if entity.Password != dto.Password {
		return nil, errors.New("invalid password")
	}

	token := ""

	return &LoginVO{
		ID:       entity.Id,
		UserName: entity.Username,
		Name:     entity.Name,
		Token:    token,
	}, nil
}
