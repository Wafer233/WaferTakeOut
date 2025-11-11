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

func (repo *CategoryRepository) GetsByNamePaged(ctx context.Context, name string,
	page, pageSize int) ([]*category.Category, int64, error) {

	offset := (page - 1) * pageSize
	total := int64(0)
	records := make([]*category.Category, 0)

	err := repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Where("name = ?", name).
		Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	err = repo.db.WithContext(ctx).
		Model(&category.Category{}).
		Where("name = ?", name).
		Offset(offset).
		Limit(pageSize).
		Find(&records).
		Error

	if err != nil {
		return nil, 0, err
	}
	return records, total, nil

}
