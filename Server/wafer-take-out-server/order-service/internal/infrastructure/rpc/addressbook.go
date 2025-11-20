package rpc

import (
	"context"

	addressbookpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/proto/addressbook"
	"github.com/jinzhu/copier"
)

type AddressBookService struct {
	client addressbookpb.AddressBookServiceClient
}

func NewAddressBookService(client addressbookpb.AddressBookServiceClient) *AddressBookService {
	return &AddressBookService{client: client}
}

func (svc *AddressBookService) FindById(ctx context.Context, id int64) (*AddressBookVO, error) {

	req := addressbookpb.IdMessage{Id: id}
	resp, err := svc.client.FindById(ctx, &req)
	if err != nil {
		return &AddressBookVO{}, err
	}
	vo := &AddressBookVO{}
	_ = copier.Copy(vo, resp)
	return vo, nil
}
