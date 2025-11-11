package categoryApp

import "context"

func (svc *CategoryService) FlipStatus(ctx context.Context, id int64, status int) error {
	err := svc.repo.UpdateStatusById(ctx, id, status)
	return err
}
