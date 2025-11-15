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

func (svc *AddressService) FindDefault(ctx context.Context, userId int64) (AddressBookVO, error) {

	addr, err := svc.repo.FindByUserIdDefault(ctx, userId)
	if err != nil || addr == nil {
		return AddressBookVO{}, err
	}

	vo := AddressBookVO{
		Id:           addr.Id,
		UserId:       addr.UserId,
		Consignee:    addr.Consignee,
		Sex:          addr.Sex,
		Phone:        addr.Phone,
		ProvinceCode: addr.ProvinceCode,
		ProvinceName: addr.ProvinceName,
		CityCode:     addr.CityCode,
		CityName:     addr.CityName,
		DistrictCode: addr.DistrictCode,
		DistrictName: addr.DistrictName,
		Detail:       addr.Detail,
		Label:        addr.Label,
		IsDefault:    addr.IsDefault,
	}

	return vo, nil
}

func (svc *AddressService) UpdateDefault(ctx context.Context, userId int64, addrId int64) error {
	//先看有没有default的
	addr, err := svc.repo.FindByUserIdDefault(ctx, userId)
	if err != nil {
		return err
	}

	// 有就更新改地址id
	if addr != nil {
		tmpId := addr.Id
		err = svc.repo.UpdateDefault(ctx, userId, tmpId, 0)
		if err != nil {
			return err
		}
	}

	// 然后更新新的
	err = svc.repo.UpdateDefault(ctx, userId, addrId, 1)
	return err

}

func (svc *AddressService) FindById(ctx context.Context, id int64) (AddressBookVO, error) {

	addr, err := svc.repo.FindById(ctx, id)
	if err != nil {
		return AddressBookVO{}, err
	}

	vo := AddressBookVO{
		Id:           addr.Id,
		UserId:       addr.UserId,
		Consignee:    addr.Consignee,
		Sex:          addr.Sex,
		Phone:        addr.Phone,
		ProvinceCode: addr.ProvinceCode,
		ProvinceName: addr.ProvinceName,
		CityCode:     addr.CityCode,
		CityName:     addr.CityName,
		DistrictCode: addr.DistrictCode,
		DistrictName: addr.DistrictName,
		Detail:       addr.Detail,
		Label:        addr.Label,
		IsDefault:    addr.IsDefault,
	}

	return vo, nil
}

func (svc *AddressService) DeleteById(ctx context.Context, addrId int64) error {
	err := svc.repo.DeleteById(ctx, addrId)
	return err
}

func (svc *AddressService) Update(ctx context.Context,
	dto *AddressDTO, userId int64) error {

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

	err := svc.repo.Update(ctx, book)

	return err
}
