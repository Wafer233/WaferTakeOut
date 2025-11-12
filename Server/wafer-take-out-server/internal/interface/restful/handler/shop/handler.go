package shopHandler

import shopApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/shop"

type ShopHandler struct {
	svc *shopApp.ShopService
}

func NewShopHandler(svc *shopApp.ShopService) *ShopHandler {
	return &ShopHandler{svc: svc}
}
