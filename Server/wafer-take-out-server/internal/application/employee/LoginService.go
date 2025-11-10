package employeeApp

import (
	"context"
	"errors"
)

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
