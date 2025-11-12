package setmealApp

import "context"

func (svc *SetMealService) Deletes(ctx context.Context, ids []int64) error {

	err := svc.setRepo.DeletesByIds(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}
