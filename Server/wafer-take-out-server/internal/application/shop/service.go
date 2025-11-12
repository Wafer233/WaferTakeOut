package shopApp

import "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/shop"

type ShopService struct {
	cache shop.ShopCache
}

func NewShopService(cache shop.ShopCache) *ShopService {
	return &ShopService{cache: cache}
}
