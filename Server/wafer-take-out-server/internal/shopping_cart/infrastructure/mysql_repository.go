package infrastructure

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/domain"
	"gorm.io/gorm"
)

type DefaultShoppingCartRepository struct {
	db *gorm.DB
}

func NewDefaultShoppingCartRepository(db *gorm.DB) domain.ShoppingCartRepository {
	return &DefaultShoppingCartRepository{
		db: db,
	}
}

func (repo *DefaultShoppingCartRepository) Create(ctx context.Context, cart *domain.ShoppingCart) error {

	err := repo.db.WithContext(ctx).
		Model(&domain.ShoppingCart{}).
		Create(&cart).Error

	return err
}
