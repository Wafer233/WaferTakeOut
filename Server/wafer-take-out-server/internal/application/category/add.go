package categoryApp

import (
	"context"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
)

// 这里我都要骂人了，json绑定一直失败，前端文档说传进来的是int实际上是string！！
type AddCategoryDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Sort string `json:"sort"`
	Type string `json:"type"`
}

func (svc *CategoryService) AddCategory(ctx context.Context, dto *AddCategoryDTO, curId int64) error {

	intSort, _ := strconv.Atoi(dto.Sort)
	intType, _ := strconv.Atoi(dto.Type)

	entity := category.Category{
		ID:         dto.ID,
		Name:       dto.Name,
		Sort:       intSort,
		Type:       intType,
		Status:     1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		CreateUser: curId,
		UpdateUser: curId,
	}

	err := svc.repo.Insert(ctx, &entity)
	return err
}
