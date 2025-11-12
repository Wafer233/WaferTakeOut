package setmealImpl

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal"
	"gorm.io/gorm"
)

type DefaultSetMealRepository struct {
	db *gorm.DB
}

func NewSetMealRepository(db *gorm.DB) *DefaultSetMealRepository {
	return &DefaultSetMealRepository{db: db}
}

func (repo *DefaultSetMealRepository) GetsPaged(ctx context.Context, categoryID int64,
	name string, page, pageSize, status int) ([]*setmeal.SetMeal, int64, error) {

	total := int64(0)
	records := make([]*setmeal.SetMeal, 0)

	db := repo.db.WithContext(ctx).
		Model(&setmeal.SetMeal{})

	if categoryID != 0 {
		db = db.Where("category_id = ?", categoryID)
	}

	if name != "" {
		db = db.Where("name = ?", name)
	}

	if status == 0 || status == 1 {
		db = db.Where("status = ?", status)
	}

	err := db.Count(&total).Error
	if err != nil || total == 0 {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = db.Offset(offset).
		Limit(pageSize).
		Find(records).Error

	if err != nil {
		return nil, 0, err
	}

	return records, total, nil

}

func (repo *DefaultSetMealRepository) Insert(ctx context.Context, set *setmeal.SetMeal) error {

	db := repo.db.
		WithContext(ctx).
		Model(&setmeal.SetMeal{}).
		Create(set)

	if db.Error != nil {
		return db.Error
	}
	return nil
}
