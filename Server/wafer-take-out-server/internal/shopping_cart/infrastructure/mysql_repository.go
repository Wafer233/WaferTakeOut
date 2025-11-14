package infrastructure

import (
	"context"
	"errors"

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

func (repo *DefaultShoppingCartRepository) Find(ctx context.Context, uid int64, did int64, sid int64) (*domain.ShoppingCart, error) {

	cart := &domain.ShoppingCart{}

	tx := repo.db.WithContext(ctx).
		Model(&domain.ShoppingCart{}).
		Where("user_id =?", uid)

	if did != 0 {
		tx = tx.Where("dish_id =?", did)
	}

	if sid != 0 {
		tx = tx.Where("setmeal_id =?", sid)
	}

	err := tx.First(cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (repo *DefaultShoppingCartRepository) UpdateNumber(ctx context.Context, cartId int64, num int) error {

	err := repo.db.WithContext(ctx).
		Model(&domain.ShoppingCart{}).
		Where("id =?", cartId).
		Update("number", num).Error

	return err

}

func (repo *DefaultShoppingCartRepository) Create(ctx context.Context, cart *domain.ShoppingCart) error {

	err := repo.db.WithContext(ctx).
		Model(&domain.ShoppingCart{}).
		Create(&cart).Error

	return err
}
