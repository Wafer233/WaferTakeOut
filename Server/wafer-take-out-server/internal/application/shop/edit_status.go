package shopApp

import "context"

func (svc *ShopService) EditStatus(ctx context.Context, status int) error {

	err := svc.cache.Set(ctx, status)
	if err != nil {
		return err
	}
	return nil
}
