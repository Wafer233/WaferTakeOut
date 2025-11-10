package employeeApp

import (
	"context"
	"errors"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/employee"
	employeeInfra "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/employee"
)

type EmployeeService struct {
	repo *employeeInfra.EmployeeRepository
}

func NewEmployeeService(repo *employeeInfra.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
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

func (svc *EmployeeService) AddEmployee(ctx context.Context, dto *AddEmployeeDTO, id int64) error {

	entity := &employee.Employee{
		IDNumber: dto.IDNumber,
		Name:     dto.Name,
		Phone:    dto.Phone,
		Sex:      dto.Sex,
		Username: dto.UserName,
	}
	entity.Password = "123456"
	entity.Status = 1
	entity.CreateTime = time.Now()
	entity.UpdateTime = time.Now()

	entity.CreateUser = id
	entity.UpdateUser = id

	err := svc.repo.Insert(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}
