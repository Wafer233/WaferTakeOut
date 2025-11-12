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

	err := r.db.
		WithContext(ctx).
		Where("username = ?", username).
		First(&model).Error
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
	err := r.db.
		WithContext(ctx).
		Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) GetByUsernamePaged(ctx context.Context,
	name string, page int, pageSize int) (int64, []employee.Employee, error) {

	var employees []employee.Employee
	var total int64

	db := r.db.
		WithContext(ctx).
		Model(&employee.Employee{})

	if name != "" {
		db = db.Where("name = ?", name)
	}

	err := db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	offset := (page - 1) * pageSize
	err = db.
		Offset(offset).
		Limit(pageSize).
		Find(&employees).Error

	if err != nil {
		return 0, nil, err
	}

	return total, employees, nil

}

func (r *EmployeeRepository) UpdateStatusByID(ctx context.Context,
	status int, id int64) error {

	db := r.db.WithContext(ctx).
		Model(&employee.Employee{}).
		Where("id = ?", id).
		Update("status", status)

	return db.Error

}

func (r *EmployeeRepository) GetById(ctx context.Context,
	id int64) (*employee.Employee, error) {

	var entity employee.Employee

	err := r.db.WithContext(ctx).
		Model(&employee.Employee{}).
		First(&entity, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *EmployeeRepository) UpdateById(ctx context.Context,
	employee *employee.Employee) error {

	err := r.db.WithContext(ctx).
		Model(&employee).
		Where("id = ?", employee.Id).
		Updates(employee).Error

	return err

}
