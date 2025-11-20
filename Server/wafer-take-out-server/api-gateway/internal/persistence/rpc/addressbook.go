package rpc

import (
	"context"

	addressbookApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/addressbook"
	addressbookpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/addressbook"
	"github.com/jinzhu/copier"
)

type AddressService struct {
	client addressbookpb.AddressBookServiceClient
}

func NewAddressBookService(client addressbookpb.AddressBookServiceClient) *AddressService {
	return &AddressService{client: client}
}

func (svc *AddressService) Create(ctx context.Context,
	dto *addressbookApp.AddressDTO, userId int64) error {

	req := addressbookpb.AddressBookRequest{}
	_ = copier.Copy(&req, dto)
	req.UserId = userId
	_, err := svc.client.Create(ctx, &req)
	if err != nil {
		return err
	}
	return nil

}

func (svc *AddressService) FindByUserId(ctx context.Context, userId int64) ([]addressbookApp.AddressBookVO, error) {

	req := addressbookpb.IdMessage{Id: userId}
	resp, err := svc.client.FindByUserId(ctx, &req)
	if err != nil {
		return nil, err
	}

	message := resp.AddressBooks
	vos := make([]addressbookApp.AddressBookVO, len(message))
	_ = copier.Copy(&vos, &message)
	return vos, nil
}

func (svc *AddressService) FindDefault(ctx context.Context,
	userId int64) (addressbookApp.AddressBookVO, error) {

	req := addressbookpb.IdMessage{Id: userId}
	resp, err := svc.client.FindDefault(ctx, &req)
	if err != nil {
		return addressbookApp.AddressBookVO{}, err
	}
	vo := addressbookApp.AddressBookVO{}
	_ = copier.Copy(&vo, resp)
	return vo, nil
}

func (svc *AddressService) UpdateDefault(ctx context.Context, userId int64, addrId int64) error {
	req := addressbookpb.UpdateDefaultRequest{
		UserId: userId,
		Id:     addrId,
	}
	_, err := svc.client.UpdateDefault(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *AddressService) FindById(ctx context.Context,
	id int64) (addressbookApp.AddressBookVO, error) {

	req := addressbookpb.IdMessage{Id: id}
	resp, err := svc.client.FindById(ctx, &req)
	if err != nil {
		return addressbookApp.AddressBookVO{}, err
	}
	vo := addressbookApp.AddressBookVO{}
	_ = copier.Copy(&vo, resp)
	return vo, nil

}

func (svc *AddressService) DeleteById(ctx context.Context, addrId int64) error {

	req := addressbookpb.IdMessage{Id: addrId}
	_, err := svc.client.DeleteById(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *AddressService) Update(ctx context.Context,
	dto *addressbookApp.AddressDTO, userId int64) error {

	req := addressbookpb.AddressBookRequest{}
	_ = copier.Copy(&req, dto)
	req.UserId = userId
	_, err := svc.client.Update(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}
