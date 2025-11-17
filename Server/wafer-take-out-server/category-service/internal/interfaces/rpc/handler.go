package rpc

import "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/category-service/internal/application"

type CategoryHandler struct {
	svc *application.CategoryAppService
}

func NewCategoryHandler(svc *application.CategoryAppService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}
