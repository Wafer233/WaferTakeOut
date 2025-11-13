package application

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/domian"
)

type ShopService struct {
	cache domian.ShopCache
}

func NewShopService(cache domian.ShopCache) *ShopService {
	return &ShopService{cache: cache}
}
