package infrastructure

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/domain"
	"gorm.io/gorm"
)

type DefaultSetMealRepository struct {
	db *gorm.DB
}

func NewSetMealRepository(db *gorm.DB) domain.SetMealRepository {
	return &DefaultSetMealRepository{db: db}
}

func (repo *DefaultSetMealRepository) FindPage(ctx context.Context, categoryID int64,
	name string, page, pageSize, status int) ([]*domain.SetMeal, int64, error) {

	records := make([]*domain.SetMeal, 0)
	total := int64(0)

	tx := repo.db.WithContext(ctx).
		Model(&domain.SetMeal{}).Begin()

	if name != "" {
		tx = tx.Where("name = ?", name)
	}

	if categoryID != 0 {
		tx = tx.Where("category_id = ?", categoryID)
	}

	if status == 0 || status == 1 {
		tx = tx.Where("status = ?", status)
	}

	if err := tx.Count(&total).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	// record一定要给指针啊
	offset := (page - 1) * pageSize
	if err := tx.Offset(offset).
		Limit(pageSize).
		Find(&records).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
	}
	return records, total, nil

}

func (repo *DefaultSetMealRepository) Create(ctx context.Context,
	set *domain.SetMeal, dishes []*domain.SetMealDish) error {

	tx := repo.db.WithContext(ctx).Begin()

	if err := tx.Model(&domain.SetMeal{}).
		Create(&set).Error; err != nil {
		tx.Rollback()
		return err
	}

	for index, _ := range dishes {
		dishes[index].SetMealId = set.Id
	}

	if err := tx.Model(&domain.SetMealDish{}).
		Create(&dishes).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
	}
	return nil
}

func (repo *DefaultSetMealRepository) UpdateStatus(ctx context.Context,
	set *domain.SetMeal) error {

	db := repo.db.WithContext(ctx).Model(&domain.SetMeal{}).
		Where("id = ?", set.Id).
		Select("status", "update_time", "update_user").
		Updates(set)

	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultSetMealRepository) Delete(ctx context.Context, ids []int64) error {

	tx := repo.db.WithContext(ctx).Begin()

	if err := tx.Model(&domain.SetMeal{}).
		Where("id in (?)", ids).
		Delete(&domain.SetMeal{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&domain.SetMealDish{}).
		Where("setmeal_id in (?)", ids).
		Delete(&domain.SetMealDish{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (repo *DefaultSetMealRepository) FindById(ctx context.Context,
	id int64) (*domain.SetMeal, []*domain.SetMealDish, error) {

	var set domain.SetMeal
	var dishes []*domain.SetMealDish

	tx := repo.db.WithContext(ctx).Begin()

	if err := tx.Model(&domain.SetMeal{}).
		Where("id = ?", id).
		First(&set).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	if err := repo.db.WithContext(ctx).
		Model(&domain.SetMealDish{}).
		Where("setmeal_id = ?", id).
		Find(&dishes).Error; err != nil {
	}

	if err := tx.Commit().Error; err != nil {
		return nil, nil, err
	}
	return &set, dishes, nil

}

func (repo *DefaultSetMealRepository) Update(ctx context.Context,
	set *domain.SetMeal, dishes []*domain.SetMealDish) error {

	tx := repo.db.WithContext(ctx).Begin()

	if err := tx.Model(&domain.SetMeal{}).
		Where("id = ?", set.Id).
		Omit("status", "create_time", "create_user").
		Updates(set).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("setmeal_id = ?", set.Id).
		Delete(&domain.SetMealDish{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&domain.SetMealDish{}).
		Create(&dishes).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}
