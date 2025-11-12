package shopImpl

import (
	"context"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/shop"
	"github.com/redis/go-redis/v9"
)

type DefaultShopCache struct {
	rbd *redis.Client
}

func NewDefaultShopCache(rbd *redis.Client) shop.ShopCache {
	return &DefaultShopCache{rbd: rbd}
}

func (c *DefaultShopCache) Set(ctx context.Context, status int) error {
	key := "ShopStatus"
	rdb := c.rbd.Set(ctx, key, status, time.Hour)

	err := rdb.Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *DefaultShopCache) Get(ctx context.Context) (int, error) {
	key := "ShopStatus"
	statusStr, err := c.rbd.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	status, _ := strconv.Atoi(statusStr)

	return status, nil
}
