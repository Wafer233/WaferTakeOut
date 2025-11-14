package infrastructure

import (
	"context"
	"errors"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/domain"
	"gorm.io/gorm"
)

type DefaultDishRepository struct {
	db *gorm.DB
}

func NewDefaultDishRepository(db *gorm.DB) domain.DishRepository {
	return &DefaultDishRepository{
		db: db,
	}

}

func (repo *DefaultDishRepository) FindPage(ctx context.Context, name string, categoryId int64,
	status int, page int, pageSize int) ([]*domain.Dish, int64, error) {

	dishes := make([]*domain.Dish, 0)
	total := int64(0)

	db := repo.db.WithContext(ctx).
		Model(&domain.Dish{})

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

func (repo *DefaultDishRepository) Delete(ctx context.Context, ids []int64) error {
	db := repo.db.WithContext(ctx).
		Model(&domain.Dish{}).
		Where("id in (?)", ids).
		Delete(&domain.Dish{})

	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultDishRepository) Create(ctx context.Context, dish *domain.Dish, flavors []*domain.Flavor) error {

	tx := repo.db.WithContext(ctx).Begin()

	// 创建Dish
	if err := tx.Model(&domain.Dish{}).
		Create(dish).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 获取DishId
	if dish.Id == 0 {
		tx.Rollback()
		return errors.New("获取菜品id失败")
	}

	// 更新flavor中的DishId
	for _, f := range flavors {
		f.DishId = dish.Id
	}

	//创建flavor
	if err := tx.Model(&domain.Flavor{}).
		Create(&flavors).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repo *DefaultDishRepository) UpdateStatus(ctx context.Context, entity *domain.Dish) error {

	db := repo.db.WithContext(ctx).
		Model(&domain.Dish{}).
		Where("id = ?", entity.Id).
		Select("status", "update_time", "update_user").
		Updates(entity)

	return db.Error

}

func (repo *DefaultDishRepository) FindByCategoryId(ctx context.Context, id int64) ([]*domain.Dish, error) {

	dishes := make([]*domain.Dish, 0)

	db := repo.db.WithContext(ctx).
		Model(&domain.Dish{}).
		Where("category_id = ?", id).
		Find(&dishes)

	err := db.Error
	if err != nil || len(dishes) == 0 {
		return nil, err
	}

	return dishes, nil
}

func (repo *DefaultDishRepository) FindById(ctx context.Context, id int64) (*domain.Dish, []*domain.Flavor, error) {

	dish := &domain.Dish{}
	flavors := make([]*domain.Flavor, 0)
	tx := repo.db.WithContext(ctx).Begin()
	if err := tx.Model(&domain.Dish{}).
		Where("id = ?", id).
		First(dish).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	if err := tx.Model(&domain.Flavor{}).
		Where("dish_id = ?", id).
		Find(&flavors).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, nil, err
	}
	return dish, flavors, nil

}

func (repo *DefaultDishRepository) Update(ctx context.Context, dish *domain.Dish, flavors []*domain.Flavor) error {

	tx := repo.db.WithContext(ctx).Begin()

	if err := tx.Model(&domain.Dish{}).
		Where("id = ?", dish.Id).
		Omit("id", "status", "create_time", "create_user").
		Updates(dish).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("dish_id = ?", dish.Id).
		Delete(&domain.Flavor{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(flavors) == 0 {
		return tx.Commit().Error
	}
	if err := tx.Model(&domain.Flavor{}).
		Create(&flavors).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repo *DefaultDishRepository) FindByIds(ctx context.Context,
	ids []int64) (map[int64]string, map[int64]string, error) {

	// 这里是自定义模型 只能scan
	type DishDetail struct {
		Id          int64  `gorm:"column:id;primaryKey;type:bigint,autoIncrement"`
		Description string `gorm:"column:description;type:varchar(255)"`
		Image       string `gorm:"column:image;type:varchar(255)"`
	}

	var detail []*DishDetail
	descriptions := make(map[int64]string)
	images := make(map[int64]string)

	err := repo.db.WithContext(ctx).
		Model(&domain.Dish{}).
		Where("id in (?)", ids).
		Select("id,description,image").
		Scan(&detail).Error

	if err != nil {
		return descriptions, images, err
	}

	for _, value := range detail {
		descriptions[value.Id] = value.Description
		images[value.Id] = value.Image
	}

	return descriptions, images, nil

}
