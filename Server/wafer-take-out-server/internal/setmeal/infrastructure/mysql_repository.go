package infrastructure

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/domain"
	"gorm.io/gorm"
)

type DefaultSetMealRepository struct {
	db *gorm.DB
}

func NewSetMealRepository(db *gorm.DB) *DefaultSetMealRepository {
	return &DefaultSetMealRepository{db: db}
}

func (repo *DefaultSetMealRepository) GetsPaged(ctx context.Context, categoryID int64,
	name string, page, pageSize, status int) ([]*domain.SetMeal, int64, error) {

	records := make([]*domain.SetMeal, 0)
	total := int64(0)

	db := repo.db.WithContext(ctx).
		Model(&domain.SetMeal{})

	if name != "" {
		db = db.Where("name = ?", name)
	}

	if categoryID != 0 {
		db = db.Where("category_id = ?", categoryID)
	}

	if status == 0 || status == 1 {
		db = db.Where("status = ?", status)
	}

	err := db.Count(&total).Error
	if err != nil || total == 0 {
		return nil, 0, err
	}

	// record一定要给指针啊
	offset := (page - 1) * pageSize
	db = db.Offset(offset).
		Limit(pageSize).
		Find(&records)

	err = db.Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil

}

func (repo *DefaultSetMealRepository) Insert(ctx context.Context, set *domain.SetMeal) error {

	db := repo.db.
		WithContext(ctx).
		Model(&domain.SetMeal{}).
		Create(set)

	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (repo *DefaultSetMealRepository) UpdateStatusById(ctx context.Context,
	set *domain.SetMeal) error {

	db := repo.db.WithContext(ctx).Model(&domain.SetMeal{}).
		Where("id = ?", set.Id).
		Select("status", "update_time", "update_user").
		Updates(set)

	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultSetMealRepository) DeletesByIds(ctx context.Context, ids []int64) error {

	db := repo.db.WithContext(ctx).
		Model(&domain.SetMeal{}).
		Where("id in (?)", ids).
		Delete(&domain.SetMeal{})

	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultSetMealRepository) GetById(ctx context.Context,
	id int64) (*domain.SetMeal, error) {

	var set domain.SetMeal
	err := repo.db.WithContext(ctx).
		Model(&domain.SetMeal{}).
		Where("id = ?", id).
		First(&set).Error

	if err != nil {
		return nil, err
	}
	return &set, nil

}

func (repo *DefaultSetMealRepository) UpdateById(ctx context.Context,
	set *domain.SetMeal) error {

	err := repo.db.WithContext(ctx).
		Model(&domain.SetMeal{}).
		Where("id = ?", set.Id).
		Omit("id", "status", "create_time", "create_user").
		Updates(set).Error

	if err != nil {
		return err
	}
	return nil
}
