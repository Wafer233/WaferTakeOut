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

func (repo *DefaultShoppingCartRepository) Find(ctx context.Context, uid int64,
	did int64, sid int64) ([]*domain.ShoppingCart, error) {

	cart := make([]*domain.ShoppingCart, 0)

	tx := repo.db.WithContext(ctx).
		Model(&domain.ShoppingCart{}).
		Where("user_id =?", uid)

	if did != 0 {
		tx = tx.Where("dish_id =?", did)
	}

	if sid != 0 {
		tx = tx.Where("setmeal_id =?", sid)
	}

	if err := tx.Find(&cart).Error; err != nil {
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

func (repo *DefaultShoppingCartRepository) FindByUserId(ctx context.Context, userId int64) ([]*domain.ShoppingCart, error) {

	records := make([]*domain.ShoppingCart, 0)
	err := repo.db.WithContext(ctx).
		Model(&domain.ShoppingCart{}).
		Where("user_id =?", userId).
		Find(&records).Error

	if err != nil {
		return nil, err
	}
	return records, err
}

func (repo *DefaultShoppingCartRepository) Delete(ctx context.Context, userId int64) error {

	err := repo.db.WithContext(ctx).
		Model(&domain.ShoppingCart{}).
		Where("user_id =?", userId).
		Delete(&domain.ShoppingCart{}).Error
	return err
}
