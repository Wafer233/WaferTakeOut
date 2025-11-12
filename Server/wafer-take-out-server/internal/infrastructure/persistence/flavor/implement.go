package flavorImpl

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/flavor"
	"gorm.io/gorm"
)

type DefaultFlavorRepository struct {
	db *gorm.DB
}

func NewFlavorRepository(db *gorm.DB) flavor.FlavorRepository {
	return &DefaultFlavorRepository{db: db}
}

func (repo *DefaultFlavorRepository) Inserts(ctx context.Context, flavors []*flavor.Flavor) error {

	if len(flavors) == 0 {
		return nil
	}

	err := repo.db.WithContext(ctx).
		Model(&flavor.Flavor{}).
		Create(&flavors).Error

	return err
}

func (repo *DefaultFlavorRepository) GetsByDishId(ctx context.Context, id int64) ([]*flavor.Flavor, error) {

	fal := make([]*flavor.Flavor, 0)
	err := repo.db.WithContext(ctx).
		Model(&flavor.Flavor{}).
		Where("dish_id = ?", id).
		Find(&fal).Error

	if err != nil {
		return nil, err
	}

	return fal, nil
}

func (repo *DefaultFlavorRepository) UpdatesByDishId(ctx context.Context, flavors []*flavor.Flavor) error {

	dishID := flavors[0].DishId

	tx := repo.db.WithContext(ctx).
		Model(&flavor.Flavor{}).
		Begin()

	err := tx.Where("dish_id = ?", dishID).Delete(&flavor.Flavor{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(flavors).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
