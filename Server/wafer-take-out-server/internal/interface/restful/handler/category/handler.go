package categoryHandler

import (
	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/category"
)

type CategoryHandler struct {
	svc *categoryApp.CategoryService
}

func NewCategoryHandler(svc *categoryApp.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc: svc,
	}
}
