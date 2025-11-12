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

func (repo *DefaultFlavorRepository) Insert(ctx context.Context, flavors []*flavor.Flavor) error {

	if len(flavors) == 0 {
		return nil
	}

	db := repo.db.WithContext(ctx).
		Model(&flavor.Flavor{}).
		Create(&flavors)

	err := db.Error
	return err
}
