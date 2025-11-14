package infrastructure

import (
	"context"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/domian"
	"github.com/redis/go-redis/v9"
)

type DefaultShopRepository struct {
	rbd *redis.Client
}

func NewDefaultShopRepository(rbd *redis.Client) domian.ShopRepository {
	return &DefaultShopRepository{rbd: rbd}
}

func (c *DefaultShopRepository) Update(ctx context.Context, status int) error {
	key := "ShopStatus"
	rdb := c.rbd.Set(ctx, key, status, time.Hour)

	err := rdb.Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *DefaultShopRepository) Find(ctx context.Context) (int, error) {
	key := "ShopStatus"
	statusStr, err := c.rbd.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	status, _ := strconv.Atoi(statusStr)

	return status, nil
}
