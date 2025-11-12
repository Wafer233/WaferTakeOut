package setmealHandler

import setmealApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/setmeal"

type SetMealHandler struct {
	svc *setmealApp.SetMealService
}

func NewSetMealHandler(svc *setmealApp.SetMealService) *SetMealHandler {
	return &SetMealHandler{
		svc: svc,
	}
}
