package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/domain"
	"github.com/redis/go-redis/v9"
)

//缓存一致性 = 让缓存永远跟数据库保持同步。
//解决方式 = “写请求一定要删缓存”，读请求走 Cache→DB 双查策略。

type CachedDishRepository struct {
	repo  *DefaultDishRepository
	cache *redis.Client
}

func (c *CachedDishRepository) FindDescriptionById(ctx context.Context,
	id int64) (string, string, error) {

	des, img, err := c.repo.FindDescriptionById(ctx, id)
	if err != nil {
		return "", "", err
	}
	return des, img, nil

}

func (c *CachedDishRepository) FindByCategoryIdFlavor(ctx context.Context, categoryId int64) ([]*domain.Dish,
	map[int64][]*domain.Flavor, error) {
	type Cache struct {
		Dishes  []*domain.Dish             `json:"dishes"`
		Flavors map[int64][]*domain.Flavor `json:"flavors"`
	}

	// 生成key
	key := fmt.Sprintf("dish_%d", categoryId)

	// 尝试从Redis里头读取
	val, err := c.cache.Get(ctx, key).Result()
	if err == nil && val != "" {
		var cache Cache
		if err = json.Unmarshal([]byte(val), &cache); err == nil {
			return cache.Dishes, cache.Flavors, nil
		}
		_ = c.cache.Del(ctx, key).Err()
	}

	// 查数据库
	dishes, flavors, err := c.repo.FindByCategoryIdFlavor(ctx, categoryId)
	if err != nil {
		return nil, nil, err
	}

	// 写入缓存
	newCache := Cache{
		Dishes:  dishes,
		Flavors: flavors,
	}

	bytes, _ := json.Marshal(&newCache)
	err = c.cache.Set(ctx, key, bytes, 1*time.Hour).Err()
	if err != nil {
		return nil, nil, err
	}

	return dishes, flavors, nil
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
	return c.repo.FindByCategoryId(ctx, categoryId)
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

func NewCachedDishRepository(
	repo *DefaultDishRepository,
	cache *redis.Client,
) domain.DishRepository {
	return &CachedDishRepository{
		repo:  repo,
		cache: cache,
	}
}
