package infrastructure

import (
	"context"
	"errors"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/order/domain"
	"gorm.io/gorm"
)

type DefaultOrderRepository struct {
	db *gorm.DB
}

func (repo *DefaultOrderRepository) UpdateStatus(ctx context.Context, order *domain.Order) error {

	err := repo.db.WithContext(ctx).
		Model(order).
		Where("id = ?", order.Id).
		Select("status, pay_status, checkout_time").
		Updates(order).Error

	return err
}

func (repo *DefaultOrderRepository) FindByNumber(ctx context.Context, number string) (*domain.Order, error) {

	var order *domain.Order

	err := repo.db.
		WithContext(ctx).
		Model(&domain.Order{}).
		Where("number = ?", number).
		First(&order).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (repo *DefaultOrderRepository) Create(ctx context.Context, order *domain.Order) error {

	if err := repo.db.WithContext(ctx).
		Model(&domain.Order{}).
		Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DefaultOrderRepository) CreateDetail(ctx context.Context, details []*domain.OrderDetail) error {

	if err := repo.db.WithContext(ctx).
		Model(&domain.OrderDetail{}).
		Create(&details).Error; err != nil {
		return err
	}
	return nil
}

func NewDefaultOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &DefaultOrderRepository{
		db: db,
	}
}
