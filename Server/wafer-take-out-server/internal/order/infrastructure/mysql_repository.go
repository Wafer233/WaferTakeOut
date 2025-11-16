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

func (repo *DefaultOrderRepository) FindDetailByOrderId(ctx context.Context,
	id int64) ([]*domain.OrderDetail, error) {

	var details []*domain.OrderDetail
	err := repo.db.WithContext(ctx).Model(&domain.OrderDetail{}).
		Where("id = ?", id).Find(&details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (repo *DefaultOrderRepository) FindDetailByOrderIds(ctx context.Context,
	ids []int64) (map[int64][]*domain.OrderDetail, error) {

	tx := repo.db.WithContext(ctx).Model(domain.OrderDetail{}).Begin()
	var curDetails []*domain.OrderDetail
	detailMap := map[int64][]*domain.OrderDetail{}
	var err error

	for _, v := range ids {
		if err = tx.Where("order_id = ?", v).Find(&curDetails).Error; err != nil {
			tx.Rollback()
			return map[int64][]*domain.OrderDetail{}, err
		}
		detailMap[v] = curDetails
	}

	if err = tx.Commit().Error; err != nil {
		return map[int64][]*domain.OrderDetail{}, err
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

	err := repo.db.WithContext(ctx).
		Model(order).
		Where("id = ?", order.Id).
		Select("status", "pay_status", "checkout_time").
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
