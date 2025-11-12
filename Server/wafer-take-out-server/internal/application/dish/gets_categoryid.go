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

	for index, record := range records {
		record.ID = dishes[index].Id
		record.Name = dishes[index].Name
		record.CategoryId = dishes[index].CategoryId
		record.Price = dishes[index].Price
		record.Image = dishes[index].Image
		record.Description = dishes[index].Description
		record.Status = dishes[index].Status
		record.UpdateTime = dishes[index].UpdateTime.Format("2006-01-02 15:04:05")
		record.CategoryName = category.Name
	}
	return records, nil
}
