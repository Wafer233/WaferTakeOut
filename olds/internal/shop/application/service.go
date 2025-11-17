package application

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/domian"
)

type ShopService struct {
	repo domian.ShopRepository
}

func NewShopService(cache domian.ShopRepository) *ShopService {
	return &ShopService{repo: cache}
}

func (svc *ShopService) UpdateStatus(ctx context.Context, status int) error {

	err := svc.repo.Update(ctx, status)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ShopService) FindStatus(ctx context.Context) (int, error) {

	status, err := svc.repo.Find(ctx)
	if err != nil {
		return 0, err
	}
	return status, nil
}
