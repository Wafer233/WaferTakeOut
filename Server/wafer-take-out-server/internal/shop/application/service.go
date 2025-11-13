package application

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/domian"
)

type ShopService struct {
	cache domian.ShopRepository
}

func NewShopService(cache domian.ShopRepository) *ShopService {
	return &ShopService{cache: cache}
}

func (svc *ShopService) EditStatus(ctx context.Context, status int) error {

	err := svc.cache.Set(ctx, status)
	if err != nil {
		return err
	}
	return nil
}

func (svc *ShopService) GetStatus(ctx context.Context) (int, error) {

	status, err := svc.cache.Get(ctx)
	if err != nil {
		return 0, err
	}
	return status, nil
}
