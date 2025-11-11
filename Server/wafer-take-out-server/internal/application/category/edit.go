package categoryApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
)

// 这里前端有个bug，就是如果我不点击对应的栏，从新输入的话，他前端返回的sort就是int类型
// 如果点击了，就是string类型，导致我反序列化失败。这里我就用string。
type EditCategoryDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort,string"`
	//Type int    `json:"type,string"`
}

func (svc *CategoryService) EditCategory(ctx context.Context, dto *EditCategoryDTO, curId int64) error {

	entity := category.Category{
		ID:         dto.ID,
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
