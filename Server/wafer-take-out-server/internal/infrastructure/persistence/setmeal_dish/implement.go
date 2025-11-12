package setmealDishImpl

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal_dish"
	"gorm.io/gorm"
)

type DefaultSetMealDishRepository struct {
	db *gorm.DB
}

func NewSetMealDishRepository(db *gorm.DB) *DefaultSetMealDishRepository {
	return &DefaultSetMealDishRepository{db: db}
}

func (repo *DefaultSetMealDishRepository) Inserts(ctx context.Context,
	sets []*setmeal_dish.SetMealDish) error {

	db := repo.db.WithContext(ctx).
		Model(&setmeal_dish.SetMealDish{}).
		Create(sets)

	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultSetMealDishRepository) GetsBySetMealId(ctx context.Context,
	id int64) ([]*setmeal_dish.SetMealDish, error) {

	dishes := make([]*setmeal_dish.SetMealDish, 0)

	err := repo.db.WithContext(ctx).
		Model(&setmeal_dish.SetMealDish{}).
		Where("setmeal_id = ?", id).
		Find(&dishes).Error

	if err != nil || len(dishes) == 0 {
		return nil, err
	}
	return dishes, nil
}

func (repo *DefaultSetMealDishRepository) UpdatesBySetMealId(ctx context.Context,
	dishes []*setmeal_dish.SetMealDish) error {

	tx := repo.db.WithContext(ctx).
		Model(&setmeal_dish.SetMealDish{}).
		Begin()

	if err := tx.Where("setmeal_id = ?", dishes[0].SetMealId).
		Delete(&setmeal_dish.SetMealDish{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&dishes).Error; err != nil {
		tx.Rollback()
		return err
	}

	err := tx.Commit().Error
	if err != nil {
		return err
	}
	return nil

}
