package categoryApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
)

type EditCategoryDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort,string"`
	Type int    `json:"type,string"`
}

func (svc *CategoryService) EditCategory(ctx context.Context, dto *EditCategoryDTO, curId int64) error {

	entity := category.Category{
		ID:         dto.ID,
		Type:       dto.Type,
		Name:       dto.Name,
		Sort:       dto.Sort,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}
	err := svc.repo.UpdateById(ctx, &entity)
	if err != nil {
		return err
	}
	return nil
}
