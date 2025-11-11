package categoryApp

import "context"

func (svc *CategoryService) DeleteCategory(ctx context.Context, id int64) error {
	err := svc.repo.DeleteById(ctx, id)
	return err
}
