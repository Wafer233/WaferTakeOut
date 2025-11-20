package rpc

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/addressbook-service/internal/application"
	addressbookpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/addressbook-service/proto"
	"github.com/jinzhu/copier"
)

type AddressBookHandler struct {
	svc *application.AddressService
	addressbookpb.UnimplementedAddressBookServiceServer
}

func NewAddressBookHandler(svc *application.AddressService) *AddressBookHandler {
	return &AddressBookHandler{svc: svc}
}

func (h *AddressBookHandler) Create(ctx context.Context,
	req *addressbookpb.AddressBookRequest) (*addressbookpb.EmptyMessage, error) {

	dto := application.AddressDTO{}
	_ = copier.Copy(&dto, req)
	userId := req.UserId
	err := h.svc.Create(ctx, &dto, userId)
	if err != nil {
		return nil, err
	}
	return &addressbookpb.EmptyMessage{}, nil
}

func (h *AddressBookHandler) FindByUserId(ctx context.Context,
	req *addressbookpb.IdMessage) (*addressbookpb.AddressBookListResponse, error) {

	vos, err := h.svc.FindByUserId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	messages := make([]*addressbookpb.AddressBookResponse, len(vos))
	_ = copier.Copy(&messages, &vos)
	return &addressbookpb.AddressBookListResponse{
		AddressBooks: messages,
	}, nil
}

func (h *AddressBookHandler) FindDefault(ctx context.Context,
	req *addressbookpb.IdMessage) (*addressbookpb.AddressBookResponse, error) {

	vo, err := h.svc.FindDefault(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := &addressbookpb.AddressBookResponse{}
	_ = copier.Copy(resp, &vo)
	return resp, nil
}

func (h *AddressBookHandler) UpdateDefault(ctx context.Context,
	req *addressbookpb.UpdateDefaultRequest) (*addressbookpb.EmptyMessage, error) {

	err := h.svc.UpdateDefault(ctx, req.UserId, req.Id)
	if err != nil {
		return nil, err
	}
	return &addressbookpb.EmptyMessage{}, nil
}

func (h *AddressBookHandler) FindById(ctx context.Context,
	req *addressbookpb.IdMessage) (*addressbookpb.AddressBookResponse, error) {

	vo, err := h.svc.FindById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := &addressbookpb.AddressBookResponse{}
	_ = copier.Copy(resp, &vo)
	return resp, nil
}

func (h *AddressBookHandler) DeleteById(ctx context.Context,
	req *addressbookpb.IdMessage) (*addressbookpb.EmptyMessage, error) {

	err := h.svc.DeleteById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &addressbookpb.EmptyMessage{}, nil
}

func (h *AddressBookHandler) Update(ctx context.Context,
	req *addressbookpb.AddressBookRequest) (*addressbookpb.EmptyMessage, error) {

	dto := application.AddressDTO{}
	_ = copier.Copy(&dto, req)
	err := h.svc.Update(ctx, &dto, req.UserId)
	if err != nil {
		return nil, err
	}
	return &addressbookpb.EmptyMessage{}, nil
}
