package application

import (
	"context"
	"strconv"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/domain"
)

type AddressService struct {
	repo domain.AddressRepository
}

func NewAddressService(repo domain.AddressRepository) *AddressService {
	return &AddressService{repo: repo}
}

func (svc *AddressService) Create(ctx context.Context, dto *AddressDTO, userId int64) error {

	label := strconv.Itoa(dto.Label)

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
		Label:        label,
		IsDefault:    dto.IsDefault,
	}

	err := svc.repo.Create(ctx, book)

	return err
}

func (svc *AddressService) FindByUserId(ctx context.Context, userId int64) ([]AddressBookVO, error) {

	books, err := svc.repo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	vos := make([]AddressBookVO, len(books))
	for i, v := range books {
		vos[i] = AddressBookVO{
			Id:           v.Id,
			UserId:       v.UserId,
			Consignee:    v.Consignee,
			Sex:          v.Sex,
			Phone:        v.Phone,
			ProvinceCode: v.ProvinceCode,
			ProvinceName: v.ProvinceName,
			CityCode:     v.CityCode,
			CityName:     v.CityName,
			DistrictCode: v.DistrictCode,
			DistrictName: v.DistrictName,
			Detail:       v.Detail,
			Label:        v.Label,
			IsDefault:    v.IsDefault,
		}
	}

	return vos, nil
}
