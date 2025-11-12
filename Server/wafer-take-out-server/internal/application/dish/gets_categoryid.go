package dishApp

import (
	"context"
)

func (svc *DishService) GetDishCategory(ctx context.Context, id int64) ([]*Record, error) {

	dishes, err := svc.dishRepo.GetsByCategoryId(ctx, id)
	if err != nil || len(dishes) == 0 {
		return nil, err
	}

	records := make([]*Record, len(dishes))

	category, err := svc.cateRepo.GetById(ctx, dishes[0].CategoryId)
	if err != nil {
		return nil, err
	}

	for index, d := range dishes {
		records[index] = &Record{
			ID:           d.Id,
			Name:         d.Name,
			CategoryId:   d.CategoryId,
			Price:        d.Price,
			Image:        d.Image,
			Description:  d.Description,
			Status:       d.Status,
			UpdateTime:   d.UpdateTime.Format("2006-01-02 15:04"),
			CategoryName: category.Name,
		}
	}
	return records, nil
}
