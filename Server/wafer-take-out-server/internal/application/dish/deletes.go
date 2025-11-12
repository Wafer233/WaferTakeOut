package dishApp

import "context"

func (svc *DishService) DeleteDishes(ctx context.Context, idArr []int64) error {

	err := svc.dishRepo.DeletesById(ctx, idArr)
	if err != nil {
		return err
	}
	return nil
}
