package categoryApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
)

func (svc *CategoryService) FlipStatus(ctx context.Context, id int64, status int, curId int64) error {

	entity := &category.Category{
		ID:         id,
		Status:     status,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}
	err := svc.repo.UpdateStatusById(ctx, entity)
	return err
}
