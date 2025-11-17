package infrastructure

import (
	"context"
	"errors"

	domain2 "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/internal/domain"
	"gorm.io/gorm"
)

type DefaultAddressRepository struct {
	db *gorm.DB
}

func (repo *DefaultAddressRepository) DeleteById(ctx context.Context, addrId int64) error {

	err := repo.db.WithContext(ctx).
		Model(&domain2.AddressBook{}).
		Where("id = ?", addrId).
		Delete(&domain2.AddressBook{}).Error

	return err
}

func (repo *DefaultAddressRepository) Update(ctx context.Context, book *domain2.AddressBook) error {

	err := repo.db.
		WithContext(ctx).
		Model(&domain2.AddressBook{}).
		Where("id = ?", book.Id).
		Updates(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultAddressRepository) UpdateDefault(ctx context.Context, userId int64,
	addrId int64, isDefault int) error {

	err := repo.db.WithContext(ctx).
		Model(&domain2.AddressBook{}).
		Where("user_id = ?", userId).
		Where("id = ?", addrId).
		Update("is_default", isDefault).Error

	return err

}

func (repo *DefaultAddressRepository) FindByUserIdDefault(ctx context.Context,
	userId int64) (*domain2.AddressBook, error) {

	var addr *domain2.AddressBook

	err := repo.db.WithContext(ctx).
		Model(&domain2.AddressBook{}).
		Where("user_id = ?", userId).
		Where("is_default = ?", 1).
		First(&addr).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func (repo *DefaultAddressRepository) FindByUserId(ctx context.Context, userId int64) ([]*domain2.AddressBook, error) {
	books := make([]*domain2.AddressBook, 0)

	err := repo.db.WithContext(ctx).Model(&domain2.AddressBook{}).Where("user_id = ?", userId).Find(&books).Error
	if err != nil || len(books) == 0 {
		return nil, err
	}
	return books, nil
}

func (repo *DefaultAddressRepository) Create(ctx context.Context, book *domain2.AddressBook) error {

	err := repo.db.WithContext(ctx).Model(&domain2.AddressBook{}).Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *DefaultAddressRepository) FindById(ctx context.Context,
	id int64) (*domain2.AddressBook, error) {

	var addr *domain2.AddressBook

	err := repo.db.WithContext(ctx).
		Model(&domain2.AddressBook{}).
		Where("id = ?", id).
		First(&addr).Error

	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return addr, nil
}

func NewDefaultAddressRepository(db *gorm.DB) domain2.AddressRepository {
	return &DefaultAddressRepository{db: db}
}
