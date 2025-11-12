package shopApp

import "context"

func (svc *ShopService) GetStatus(ctx context.Context) (int, error) {

	status, err := svc.cache.Get(ctx)
	if err != nil {
		return 0, err
	}
	return status, nil
}
