package employeeImpl

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/employee"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) GetByUsername(ctx context.Context, username string) (*employee.Employee, error) {
	var model employee.Employee

	err := r.db.WithContext(ctx).Where("username = ?", username).First(&model).Error
	if err != nil {
		return nil, err
	}

	return &employee.Employee{
		Id:       model.Id,
		Username: model.Username,
		Password: model.Password,
		Name:     model.Name,
	}, nil
}

func (r *EmployeeRepository) Insert(ctx context.Context, model *employee.Employee) error {
	err := r.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}
