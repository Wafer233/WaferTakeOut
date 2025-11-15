package application

import (
	"context"
	"time"

	dish "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/domain"
	setm "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/domain"
	cart "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/domain"
)

type ShoppingCartService struct {
	cartRepo cart.ShoppingCartRepository
	dishRepo dish.DishRepository
	setRepo  setm.SetMealRepository
}

func NewShoppingCartService(cartRepo cart.ShoppingCartRepository,
	dishRepo dish.DishRepository,
	setRepo setm.SetMealRepository,
) *ShoppingCartService {
	return &ShoppingCartService{
		cartRepo: cartRepo,
		dishRepo: dishRepo,
		setRepo:  setRepo,
	}
}

func (svc *ShoppingCartService) Add(ctx context.Context, dto *CartDTO, curId int64) error {

	dishId := dto.DishId
	setId := dto.SetMealId

	curCart, err := svc.cartRepo.Find(ctx, curId, dishId, setId)
	if err != nil {
		return err
	}

	if curCart != nil && len(curCart) > 0 {
		num := curCart[0].Number + 1
		err = svc.cartRepo.UpdateNumber(ctx, curCart[0].Id, num)
		if err != nil {
			return err
		}
		return nil
	}

	// 2.2.如果没有就创建
	// 2.2.1 获取name，image和amount
	// dish

	if dishId != 0 {
		dishEntity, _, er := svc.dishRepo.FindById(ctx, dishId)
		if er != nil {
			return er
		}
		item := &cart.ShoppingCart{
			Name:       dishEntity.Name,
			Image:      dishEntity.Image,
			UserId:     curId,
			DishId:     dishId,
			SetmealId:  setId,
			DishFlavor: dto.DishFlavor,
			Number:     1,
			Amount:     dishEntity.Price,
			CreateTime: time.Now(),
		}
		err = svc.cartRepo.Create(ctx, item)
		return err

	} else {
		setEntity, _, er := svc.setRepo.FindById(ctx, setId)
		if er != nil {
			return er
		}
		item := &cart.ShoppingCart{
			Name:       setEntity.Name,
			Image:      setEntity.Image,
			UserId:     curId,
			SetmealId:  setId,
			DishId:     dishId,
			DishFlavor: dto.DishFlavor,
			Number:     1,
			Amount:     setEntity.Price,
			CreateTime: time.Now(),
		}
		err = svc.cartRepo.Create(ctx, item)
		return err
	}
}

func (svc *ShoppingCartService) FindByUserId(ctx context.Context, userId int64) ([]RecordVO, error) {

	records, err := svc.cartRepo.Find(ctx, userId, 0, 0)

	vos := make([]RecordVO, len(records))
	for i, r := range records {
		vos[i] = RecordVO{
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

	curCart, err := svc.cartRepo.Find(ctx, curId, dishId, setId)
	if err != nil {
		return err
	}

	num := curCart[0].Number - 1

	err = svc.cartRepo.UpdateNumber(ctx, curCart[0].Id, num)

	return err

}

func (svc *ShoppingCartService) Delete(ctx context.Context, curId int64) error {

	err := svc.cartRepo.Delete(ctx, curId)
	return err

}
