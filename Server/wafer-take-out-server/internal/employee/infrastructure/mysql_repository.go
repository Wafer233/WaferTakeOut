package infrastructure

import (
	"context"
	"errors"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/domain"
	"gorm.io/gorm"
)

type DefaultEmployeeRepository struct {
	db *gorm.DB
}

func NewDefaultEmployeeRepository(db *gorm.DB) *DefaultEmployeeRepository {
	return &DefaultEmployeeRepository{db: db}
}

func (r *DefaultEmployeeRepository) FindByUsername(ctx context.Context, username string) (*domain.Employee, error) {
	var model domain.Employee

	err := r.db.
		WithContext(ctx).
		Where("username = ?", username).
		First(&model).Error
	if err != nil {
		return nil, err
	}

	return &domain.Employee{
		Id:       model.Id,
		Username: model.Username,
		Password: model.Password,
		Name:     model.Name,
	}, nil
}

func (r *DefaultEmployeeRepository) Create(ctx context.Context, model *domain.Employee) error {
	err := r.db.
		WithContext(ctx).
		Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DefaultEmployeeRepository) FindPage(ctx context.Context,
	name string, page int, pageSize int) (int64, []domain.Employee, error) {

	var employees []domain.Employee
	var total int64

	db := r.db.
		WithContext(ctx).
		Model(&domain.Employee{})

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

func (r *DefaultEmployeeRepository) UpdateStatus(ctx context.Context,
	emp *domain.Employee) error {

	id := emp.Id
	db := r.db.WithContext(ctx).
		Model(&domain.Employee{}).
		Where("id = ?", id).
		Select("update_time", "update_user", "status").
		Updates(emp)

	return db.Error

}

func (r *DefaultEmployeeRepository) FindById(ctx context.Context,
	id int64) (*domain.Employee, error) {

	var entity domain.Employee

	err := r.db.WithContext(ctx).
		Model(&domain.Employee{}).
		First(&entity, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *DefaultEmployeeRepository) Update(ctx context.Context,
	employee *domain.Employee) error {

	err := r.db.WithContext(ctx).
		Model(&employee).
		Where("id = ?", employee.Id).
		Updates(employee).Error

	return err

}

func (r *DefaultEmployeeRepository) UpdatePassword(ctx context.Context, id int64,
	old string, neo string) error {
	realPassword := ""

	tx := r.db.WithContext(ctx).
		Model(&domain.Employee{}).
		Begin()

	if err := tx.Where("id = ?", id).
		Select("password").
		Scan(&realPassword).Error; err != nil {
		tx.Rollback()
		return err
	}

	if realPassword != old {
		tx.Rollback()
		return errors.New("密码错误")
	}

	if err := tx.Where("id = ?", id).
		Update("password", neo).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}
