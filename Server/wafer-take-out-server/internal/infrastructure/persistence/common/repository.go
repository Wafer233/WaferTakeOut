package commonImpl

import "gorm.io/gorm"

type CommonRepository struct {
	db *gorm.DB
}

func NewCommonRepository(db *gorm.DB) *CommonRepository {
	return &CommonRepository{db: db}
}
