package dishImpl

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/dish"
	"gorm.io/gorm"
)

type DefaultDishRepository struct {
	db *gorm.DB
}

func NewDishRepository(db *gorm.DB) dish.DishRepository {
	return &DefaultDishRepository{
		db: db,
	}

}

func (repo *DefaultDishRepository) GetsPaged(ctx context.Context, name string, categoryId int64,
	status int, page int, pageSize int) ([]*dish.Dish, int64, error) {

	dishes := make([]*dish.Dish, 0)
	total := int64(0)
	db := repo.db.WithContext(ctx).
		Model(&dish.Dish{})

	if name != "" {
		db = db.Where("name = ?", name)
	}

	if categoryId != 0 {
		db = db.Where("category_id = ?", categoryId)
	}

	if status == 0 || status == 1 {
		db = db.Where("status = ?", status)
	}

	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = db.Offset(offset).
		Limit(pageSize).
		Find(&dishes).Error

	return dishes, total, err

}

func (repo *DefaultDishRepository) DeletesById(ctx context.Context, ids []int64) error {
	db := repo.db.WithContext(ctx).
		Model(&dish.Dish{}).
		Where("id in (?)", ids).
		Delete(&dish.Dish{})

	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultDishRepository) Insert(ctx context.Context, entity *dish.Dish) error {

	db := repo.db.WithContext(ctx).
		Model(&dish.Dish{}).
		Create(&entity)

	err := db.Error
	if err != nil {
		return err
	}

	return err
}

func (repo *DefaultDishRepository) UpdateStatusById(ctx context.Context,
	entity *dish.Dish) error {

	db := repo.db.WithContext(ctx).
		Model(&dish.Dish{}).
		Where("id = ?", entity.Id).
		Select("status", "update_time", "update_user").
		Updates(entity)

	return db.Error

}

func (repo *DefaultDishRepository) GetsByCategoryId(ctx context.Context, id int64) ([]*dish.Dish, error) {

	dishes := make([]*dish.Dish, 0)
	db := repo.db.WithContext(ctx).
		Model(&dish.Dish{}).
		Where("category_id = ?", id).
		Find(&dishes)

	err := db.Error
	if err != nil || len(dishes) == 0 {
		return nil, err
	}

	return dishes, nil
}

func (repo *DefaultDishRepository) GetById(ctx context.Context, id int64) (*dish.Dish, error) {

	entity := &dish.Dish{}
	db := repo.db.WithContext(ctx).
		Model(&dish.Dish{}).
		Where("id = ?", id).
		First(entity)

	err := db.Error
	if err != nil {
		return nil, err
	}
	return entity, nil

}
