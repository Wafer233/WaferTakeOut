package infrastructure

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/domain"
	"gorm.io/gorm"
)

type DefaultCategoryRepository struct {
	db *gorm.DB
}

func NewDefaultCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &DefaultCategoryRepository{
		db: db,
	}
}

func (repo *DefaultCategoryRepository) Create(ctx context.Context, entity *domain.Category) error {

	err := repo.db.WithContext(ctx).
		Model(&domain.Category{}).
		Create(entity).Error

	return err
}

func (repo *DefaultCategoryRepository) FindPage(ctx context.Context, name string,
	curType int, page, pageSize int) ([]*domain.Category, int64, error) {

	offset := (page - 1) * pageSize
	total := int64(0)
	records := make([]*domain.Category, 0)

	db := repo.db.WithContext(ctx).Model(&domain.Category{})

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

func (repo *DefaultCategoryRepository) Update(ctx context.Context, entity *domain.Category) error {

	db := repo.db.WithContext(ctx).
		Model(&domain.Category{}).
		Where("id = ?", entity.ID).
		Select("name", "sort", "update_time", "update_user").
		Updates(entity)
	err := db.Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultCategoryRepository) UpdateStatus(ctx context.Context, entity *domain.Category) error {

	db := repo.db.WithContext(ctx).
		Model(&domain.Category{}).
		Where("id = ?", entity.ID).
		Select("status", "update_time", "update_user").
		Updates(entity)

	err := db.Error

	return err
}

func (repo *DefaultCategoryRepository) Delete(ctx context.Context, id int64) error {
	db := repo.db.WithContext(ctx).
		Model(&domain.Category{}).
		Where("id = ?", id).
		Delete(&domain.Category{})

	return db.Error
}

func (repo *DefaultCategoryRepository) FindByType(ctx context.Context, curType int) ([]*domain.Category, error) {
	entity := make([]*domain.Category, 0)
	db := repo.db.WithContext(ctx).
		Model(&domain.Category{})
	if curType > 0 {
		db.Where("type = ?", curType)
	}
	if err := db.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (repo *DefaultCategoryRepository) FindById(ctx context.Context, id int64) (*domain.Category, error) {
	entity := domain.Category{}
	db := repo.db.WithContext(ctx).
		Model(&domain.Category{}).
		Where("id = ?", id).
		First(&entity)
	err := db.Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
