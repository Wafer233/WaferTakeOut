package categoryApp

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
)

type CategoryService struct {
	repo category.CategoryRepository
}

func NewCategoryService(repo category.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}
