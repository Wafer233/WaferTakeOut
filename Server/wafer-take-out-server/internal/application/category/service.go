package categoryApp

import categoryImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/category"

type CategoryService struct {
	repo *categoryImpl.CategoryRepository
}

func NewCategoryService(repo *categoryImpl.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}
