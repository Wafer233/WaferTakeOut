package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/domain"
	"github.com/redis/go-redis/v9"
)

//缓存一致性 = 让缓存永远跟数据库保持同步。
//解决方式 = “写请求一定要删缓存”，读请求走 Cache→DB 双查策略。

type CachedDishRepository struct {
	repo  domain.DishRepository
	cache *redis.Client
}

func NewCachedDishRepository(
	repo domain.DishRepository,
	cache *redis.Client,
) domain.DishRepository {
	return &CachedDishRepository{
		repo:  repo,
		cache: cache,
	}
}

func (c *CachedDishRepository) Delete(ctx context.Context, ids []int64) error {
	// 这里涉及到增删改，所以要删除缓存

	err := c.repo.Delete(ctx, ids)

	if err != nil {
		return err
	}

	err = c.cleanCache(ctx, "dish_*")

	return err
}

func (c *CachedDishRepository) Create(ctx context.Context, dish *domain.Dish, flavors []*domain.Flavor) error {
	// 这里涉及到增删改，所以要删除缓存
	err := c.repo.Create(ctx, dish, flavors)
	if err != nil {
		return err
	}

	err = c.cleanCache(ctx, "dish_*")
	return err
}

func (c *CachedDishRepository) FindById(ctx context.Context, id int64) (*domain.Dish, []*domain.Flavor, error) {
	// 这里不用缓存
	return c.repo.FindById(ctx, id)
}

func (c *CachedDishRepository) FindByCategoryId(ctx context.Context, categoryId int64) ([]*domain.Dish, error) {
	// 这里涉及到查，要优先使用缓存

	var dishes []*domain.Dish
	//  1.构建key
	key := fmt.Sprintf("dish_%d", categoryId)

	//	2.查缓存
	// 如果 val = ""   那么 err = redis.Nil
	val, err := c.cache.Get(ctx, key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(val), &dishes)
		if err != nil {
			return nil, errors.New("反序列化失败")
		}
		return dishes, nil
	}

	// 3.如果没有缓存，查mysql
	dishes, err = c.repo.FindByCategoryId(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	// 4.从新写缓存
	bytes, err := json.Marshal(dishes)
	if err != nil {
		return nil, errors.New("序列化失败")
	}
	c.cache.Set(ctx, key, bytes, time.Hour)
	return dishes, nil

}

func (c *CachedDishRepository) FindPage(ctx context.Context, s string, i int64, i2 int, i3 int, i4 int) ([]*domain.Dish, int64, error) {
	// 不建议用缓存
	return c.repo.FindPage(ctx, s, i, i2, i3, i4)

}

func (c *CachedDishRepository) UpdateStatus(ctx context.Context, dish *domain.Dish) error {
	// 这里涉及到增删改，要删除缓存
	err := c.repo.UpdateStatus(ctx, dish)
	if err != nil {
		return err
	}
	err = c.cleanCache(ctx, "dish_*")

	return err
}

func (c *CachedDishRepository) Update(ctx context.Context, dish *domain.Dish, flavors []*domain.Flavor) error {
	err := c.repo.Update(ctx, dish, flavors)
	if err != nil {
		return err
	}

	err = c.cleanCache(ctx, "dish_*")
	return err
}

func (c *CachedDishRepository) FindByIds(ctx context.Context, int64s []int64) (map[int64]string, map[int64]string, error) {
	// 这里不用缓存
	return c.repo.FindByIds(ctx, int64s)
}

func (c *CachedDishRepository) cleanCache(ctx context.Context, pattern string) error {
	iter := c.cache.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		c.cache.Del(ctx, iter.Val())
	}

	return iter.Err()
}
