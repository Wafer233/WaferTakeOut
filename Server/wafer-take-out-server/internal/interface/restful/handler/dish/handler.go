package dishHandler

import (
	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/dish"
)

type DishHandler struct {
	svc *dishApp.DishService
}

func NewDishHandler(svc *dishApp.DishService) *DishHandler {
	return &DishHandler{svc: svc}
}
