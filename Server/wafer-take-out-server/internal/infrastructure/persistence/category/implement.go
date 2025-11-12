package categoryImpl

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
	"gorm.io/gorm"
)

type DefaultCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) category.CategoryRepository {
	return &DefaultCategoryRepository{
		db: db,
	}
}

func (repo *DefaultCategoryRepository) Insert(ctx context.Context, entity *category.Category) error {

	err := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Create(entity).Error

	return err
}

func (repo *DefaultCategoryRepository) GetsByPaged(ctx context.Context, name string,
	curType int, page, pageSize int) ([]*category.Category, int64, error) {

	offset := (page - 1) * pageSize
	total := int64(0)
	records := make([]*category.Category, 0)

	db := repo.db.WithContext(ctx).Model(&category.Category{})

	if name != "" {
		db = db.Where("name = ?", name)
	}

	if curType > 0 {
		db = db.Where("type = ?", curType)
	}

	err := db.Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	err = db.
		Offset(offset).
		Limit(pageSize).
		Find(&records).
		Error

	if err != nil {
		return nil, 0, err
	}
	return records, total, nil

}

func (repo *DefaultCategoryRepository) UpdateById(ctx context.Context, entity *category.Category) error {

	db := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Where("id = ?", entity.ID).
		Select("name", "sort", "update_time", "update_user").
		Updates(entity)
	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultCategoryRepository) UpdateStatusById(ctx context.Context, entity *category.Category) error {

	db := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Where("id = ?", entity.ID).
		Select("status", "update_time", "update_user").
		Updates(entity)

	err := db.Error

	return err
}

func (repo *DefaultCategoryRepository) DeleteById(ctx context.Context, id int64) error {
	db := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Where("id = ?", id).
		Delete(&category.Category{})

	return db.Error
}

func (repo *DefaultCategoryRepository) GetsByType(ctx context.Context, curType int) ([]*category.Category, error) {

	entity := make([]*category.Category, 0)
	db := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Where("type = ?", curType).
		Find(&entity)
	err := db.Error
	if err != nil {
		return nil, err
	}
	return entity, nil

}

func (repo *DefaultCategoryRepository) GetById(ctx context.Context, id int64) (*category.Category, error) {
	entity := category.Category{}
	db := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Where("id = ?", id).
		First(&entity)
	err := db.Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
