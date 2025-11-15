package infrastructure

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/domain"
	"gorm.io/gorm"
)

type DefaultAddressRepository struct {
	db *gorm.DB
}

func NewDefaultAddressRepository(db *gorm.DB) domain.AddressRepository {
	return &DefaultAddressRepository{db: db}
}

func (repo *DefaultAddressRepository) Create(ctx context.Context, book *domain.AddressBook) error {

	err := repo.db.WithContext(ctx).Model(&domain.AddressBook{}).Create(book).Error
	return err
}
