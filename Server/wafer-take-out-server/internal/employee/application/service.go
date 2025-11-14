package application

import (
	"context"
	"errors"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/domain"
	employeeInfra "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/infrastructure"
	"github.com/jinzhu/copier"
)

type EmployeeService struct {
	repo *employeeInfra.DefaultEmployeeRepository
}

func NewEmployeeService(repo *employeeInfra.DefaultEmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}

func (svc *EmployeeService) Login(ctx context.Context, dto *LoginDTO) (*LoginVO, error) {

	entity, err := svc.repo.FindByUsername(ctx, dto.Username)
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

func (svc *EmployeeService) FindPage(ctx context.Context, dto *PageDTO) (*PageVO, error) {

	total, employees, err := svc.repo.FindPage(ctx, dto.Name, dto.Page, dto.PageSize)

	if err != nil {
		return nil, err
	}

	empVO := make([]Employee, len(employees))
	for i, e := range employees {
		empVO[i] = Employee{
			Id:         e.Id,
			Name:       e.Name,
			Username:   e.Username,
			Password:   e.Password,
			Phone:      e.Phone,
			Sex:        e.Sex,
			IDNumber:   e.IDNumber,
			Status:     e.Status,
			CreateTime: e.CreateTime.Format("2006-01-02 15:04"),
			UpdateTime: e.UpdateTime.Format("2006-01-02 15:04"),
			CreateUser: e.CreateUser,
			UpdateUser: e.UpdateUser,
		}
	}
	return &PageVO{
		Total:     total,
		Employees: empVO,
	}, nil
}

func (svc *EmployeeService) FindById(ctx context.Context, id int64) (*Employee, error) {

	employee, err := svc.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	employeeVO := Employee{
		Id:         employee.Id,
		Name:       employee.Name,
		Username:   employee.Username,
		Password:   employee.Password,
		Phone:      employee.Phone,
		Sex:        employee.Sex,
		IDNumber:   employee.IDNumber,
		Status:     employee.Status,
		CreateTime: employee.UpdateTime.Format("2006-01-02 15:04"),
		UpdateTime: employee.UpdateTime.Format("2006-01-02 15:04"),
		CreateUser: employee.CreateUser,
		UpdateUser: employee.UpdateUser,
	}

	return &employeeVO, nil
}

func (svc *EmployeeService) Update(ctx context.Context, dto *AddEmployeeDTO, updateId int64) error {

	entity := &domain.Employee{}

	err := copier.Copy(&entity, &dto)
	if err != nil {
		return err
	}
	entity.UpdateUser = updateId
	entity.UpdateTime = time.Now()
	err = svc.repo.Update(ctx, entity)
	return err
}

func (svc *EmployeeService) UpdateStatus(ctx context.Context, status int, id int64, curId int64) error {

	emp := &domain.Employee{
		Id:         id,
		Status:     status,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}

	err := svc.repo.UpdateStatus(ctx, emp)
	return err
}

func (svc *EmployeeService) Create(ctx context.Context, dto *AddEmployeeDTO, id int64) error {

	entity := &domain.Employee{
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

	err := svc.repo.Create(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (svc *EmployeeService) UpdatePassword(ctx context.Context, dto *PasswordDTO, curid int64) error {

	id := curid
	old := dto.OldPassword
	neo := dto.NewPassword
	err := svc.repo.UpdatePassword(ctx, id, old, neo)

	return err

}
