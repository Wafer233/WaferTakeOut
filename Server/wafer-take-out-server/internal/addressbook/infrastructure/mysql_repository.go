package infrastructure

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/domain"
	"gorm.io/gorm"
)

type DefaultAddressRepository struct {
	db *gorm.DB
}

func (repo *DefaultAddressRepository) FindByUserId(ctx context.Context, userId int64) ([]*domain.AddressBook, error) {
	books := make([]*domain.AddressBook, 0)

	err := repo.db.WithContext(ctx).Model(&domain.AddressBook{}).Where("user_id = ?", userId).Find(&books).Error
	if err != nil || len(books) == 0 {
		return nil, err
	}
	return books, nil
}

func (repo *DefaultAddressRepository) Create(ctx context.Context, book *domain.AddressBook) error {

	err := repo.db.WithContext(ctx).Model(&domain.AddressBook{}).Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func NewDefaultAddressRepository(db *gorm.DB) domain.AddressRepository {
	return &DefaultAddressRepository{db: db}
}
