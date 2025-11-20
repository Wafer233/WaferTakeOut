package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/internal/domain"
	rpcClient "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/internal/infrastructure/rpc"
)

type ShoppingCartService struct {
	repo       domain.ShoppingCartRepository
	setMealSvc *rpcClient.SetMealService
	dishSvc    *rpcClient.DishService
}

func NewShoppingCartService(
	repo domain.ShoppingCartRepository,
	setMealSvc *rpcClient.SetMealService,
	dishSvc *rpcClient.DishService,
) *ShoppingCartService {
	return &ShoppingCartService{
		repo:       repo,
		setMealSvc: setMealSvc,
		dishSvc:    dishSvc,
	}
}

func (svc *ShoppingCartService) Add(ctx context.Context, dto *CartDTO, curId int64) error {

	dishId := dto.DishId
	setId := dto.SetMealId

	curCart, err := svc.repo.Find(ctx, curId, dishId, setId)
	if err != nil {
		return err
	}

	if curCart != nil && len(curCart) > 0 {
		num := curCart[0].Number + 1
		err = svc.repo.UpdateNumber(ctx, curCart[0].Id, num)
		if err != nil {
			return err
		}
		return nil
	}

	// 2.2.如果没有就创建
	// 2.2.1 获取name，image和amount
	// dish

	if dishId != 0 {
		img, name, price, er := svc.dishSvc.FindDetailById(ctx, dishId)
		if er != nil {
			return er
		}
		item := &domain.ShoppingCart{
			Name:       name,
			Image:      img,
			UserId:     curId,
			DishId:     dishId,
			SetmealId:  setId,
			DishFlavor: dto.DishFlavor,
			Number:     1,
			Amount:     price,
			CreateTime: time.Now(),
		}
		err = svc.repo.Create(ctx, item)
		return err

	} else {

		img, name, price, er := svc.setMealSvc.FindDetailById(ctx, setId)
		if er != nil {
			return er
		}
		item := &domain.ShoppingCart{
			Name:       name,
			Image:      img,
			UserId:     curId,
			SetmealId:  setId,
			DishId:     dishId,
			DishFlavor: dto.DishFlavor,
			Number:     1,
			Amount:     price,
			CreateTime: time.Now(),
		}
		err = svc.repo.Create(ctx, item)
		return err
	}
}

func (svc *ShoppingCartService) FindByUserId(ctx context.Context, userId int64) ([]CartVO, error) {

	records, err := svc.repo.Find(ctx, userId, 0, 0)

	vos := make([]CartVO, len(records))
	for i, r := range records {
		vos[i] = CartVO{
			Amount:     r.Amount,
			CreateTime: r.CreateTime.Format("2006-01-02 15:04"),
			DishFlavor: r.DishFlavor,
			DishId:     r.DishId,
			Id:         r.Id,
			Image:      r.Image,
			Name:       r.Name,
			Number:     r.Number,
			SetMealId:  r.SetmealId,
			UserID:     r.UserId,
		}

	}

	return vos, err
}

func (svc *ShoppingCartService) Sub(ctx context.Context, dto *CartDTO, curId int64) error {

	setId := dto.SetMealId
	dishId := dto.DishId

	curCart, err := svc.repo.Find(ctx, curId, dishId, setId)
	if err != nil {
		return err
	}

	num := curCart[0].Number - 1

	err = svc.repo.UpdateNumber(ctx, curCart[0].Id, num)

	return err

}

func (svc *ShoppingCartService) Delete(ctx context.Context, curId int64) error {

	err := svc.repo.Delete(ctx, curId)
	return err

}
