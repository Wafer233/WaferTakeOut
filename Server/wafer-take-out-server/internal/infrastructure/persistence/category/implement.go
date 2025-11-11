package categoryImpl

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repo *CategoryRepository) Insert(ctx context.Context, entity *category.Category) error {

	err := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Create(entity).Error

	return err
}

func (repo *CategoryRepository) GetsByPaged(ctx context.Context, name string,
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

func (repo *CategoryRepository) UpdateById(ctx context.Context, entity *category.Category) error {

	err := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Where("id = ?", entity.ID).
		Select("type", "name", "sort", "update_time", "update_user").
		Updates(entity).Error
	if err != nil {
		return err
	}
	return nil
}
