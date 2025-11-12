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
