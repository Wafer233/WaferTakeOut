package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/domain"
	"gorm.io/gorm"
)

type DefaultOrderRepository struct {
	db *gorm.DB
}

func (repo *DefaultOrderRepository) GetTotalByStatus(ctx context.Context) (int64, int64, int64, error) {
	var confirmed, deliveryInProgress, toBeConfirmed int64

	tx := repo.db.WithContext(ctx).Model(&domain.Order{})

	if err := tx.Where("status = ?", domain.CONFIRMED).
		Count(&confirmed).Error; err != nil {
		tx.Rollback()
		return 0, 0, 0, nil
	}

	if err := tx.Where("status = ?", domain.DELIVERY_IN_PROGRESS).
		Count(&deliveryInProgress).Error; err != nil {
		tx.Rollback()
		return 0, 0, 0, nil
	}

	if err := tx.Where("status = ?", domain.TO_BE_CONFIRMED).
		Count(&toBeConfirmed).Error; err != nil {
		tx.Rollback()
		return 0, 0, 0, nil
	}

	return confirmed, deliveryInProgress, toBeConfirmed, nil
}

func (repo *DefaultOrderRepository) FindPageAdmin(ctx context.Context, begin time.Time, end time.Time,
	number string, page int, pageSize int, phone string, status int) ([]*domain.Order, int64, error) {
	total := int64(0)
	tx := repo.db.WithContext(ctx).Model(&domain.Order{}).Begin()

	if !begin.IsZero() {
		tx = tx.Where("order_time BETWEEN ? AND ?", begin, end)
	}

	if number != "" {
		tx = tx.Where("number = ?", number)
	}

	if phone != "" {
		tx = tx.Where("phone = ?", phone)
	}

	if status != 0 {
		tx = tx.Where("status = ?", status)
	}

	if err := tx.Count(&total).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	var orders []*domain.Order
	offset := (page - 1) * pageSize
	if err := tx.Limit(offset).Limit(pageSize).Find(&orders).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	return orders, total, nil

}

func (repo *DefaultOrderRepository) FindDetailByOrderId(ctx context.Context,
	id int64) ([]*domain.OrderDetail, error) {

	var details []*domain.OrderDetail
	err := repo.db.WithContext(ctx).Model(&domain.OrderDetail{}).
		Where("order_id = ?", id).Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (repo *DefaultOrderRepository) FindDetailByOrderIds(ctx context.Context,
	ids []int64) (map[int64][]*domain.OrderDetail, error) {

	tx := repo.db.WithContext(ctx).Begin()

	detailMap := make(map[int64][]*domain.OrderDetail)

	for _, id := range ids {

		// 每次循环必须重新声明一个新切片！！！
		var curDetails []*domain.OrderDetail

		if err := tx.Where("order_id = ?", id).Find(&curDetails).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		detailMap[id] = curDetails
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return detailMap, nil
}

func (repo *DefaultOrderRepository) FindPage(ctx context.Context, page int,
	pageSize int, userId int64, status int) ([]*domain.Order, int64, error) {

	total := int64(0)
	offset := (page - 1) * pageSize

	tx := repo.db.WithContext(ctx).
		Model(&domain.Order{}).Begin()

	if status != 0 {
		tx = tx.Where("status = ?", status)
	}

	if err := tx.Where("user_id = ? ", userId).Count(&total).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	var records []*domain.Order
	if err := tx.Where("user_id = ? ", userId).
		Limit(pageSize).Offset(offset).
		Find(&records).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	return records, total, nil

}

func (repo *DefaultOrderRepository) UpdateStatus(ctx context.Context, order *domain.Order) error {

	tx := repo.db.WithContext(ctx).
		Model(order).
		Where("id = ?", order.Id).
		Begin()

	if order.OrderTime != nil {
		err := tx.Update("order_time", order.OrderTime).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if order.CheckoutTime != nil {
		err := tx.Update("checkout_time", order.CheckoutTime).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if order.CancelTime != nil {
		err := tx.Update("cancel_time = ?", order.CancelTime).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if order.EstimatedDeliveryTime != nil {
		err := tx.Update("estimated_delivery_time = ?", order.EstimatedDeliveryTime).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if order.DeliveryTime != nil {
		err := tx.Update("delivery_time = ?", order.DeliveryTime).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Select("status", "pay_status", "cancel_reason", "rejection_reason").
		Updates(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

	err := repo.db.WithContext(ctx).
		Model(&domain.Order{}).Create(order).Error

	return err
}

func (repo *DefaultOrderRepository) CreateDetail(ctx context.Context, details []*domain.OrderDetail) error {

	if err := repo.db.WithContext(ctx).
		Model(&domain.OrderDetail{}).
		Create(&details).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DefaultOrderRepository) FindById(ctx context.Context, id int64) (*domain.Order, error) {

	var order *domain.Order

	err := repo.db.WithContext(ctx).Model(&domain.Order{}).
		Where("id = ?", id).Find(&order).Error

	if err != nil {
		return nil, err
	}
	return order, nil

}

func NewDefaultOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &DefaultOrderRepository{
		db: db,
	}
}
