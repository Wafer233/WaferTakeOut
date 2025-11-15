package application

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/domain"
)

type AddressService struct {
	repo domain.AddressRepository
}

func NewAddressService(repo domain.AddressRepository) *AddressService {
	return &AddressService{repo: repo}
}

func (svc *AddressService) Create(ctx context.Context, dto *AddressDTO, userId int64) error {

	book := &domain.AddressBook{
		UserId:       userId,
		Consignee:    dto.Consignee,
		Sex:          dto.Sex,
		Phone:        dto.Phone,
		ProvinceCode: dto.ProvinceCode,
		ProvinceName: dto.ProvinceName,
		CityCode:     dto.CityCode,
		CityName:     dto.CityName,
		DistrictCode: dto.DistrictCode,
		DistrictName: dto.DistrictName,
		Detail:       dto.Detail,
		Label:        dto.Label,
		IsDefault:    dto.IsDefault,
	}

	err := svc.repo.Create(ctx, book)

	return err
}
