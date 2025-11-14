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

func (svc *ShoppingCartService) Create(ctx context.Context, dto *AddDTO) error {

	// 这里因为他不支持cookie我暂时不搞userid
	userId := int64(1)
	dishId := dto.DishId
	setId := dto.SetMealId
	shoppingCart := &cart.ShoppingCart{}

	// 1.先看有没有
	shoppingCart, err := svc.cartRepo.Find(ctx, userId, dishId, setId)
	if err != nil {
		return err
	}

	// 2.1.如果有就增加
	if shoppingCart != nil {
		num := shoppingCart.Number + 1

		err = svc.cartRepo.UpdateNumber(ctx, shoppingCart.Id, num)

		return err
	}
	// 2.2.如果没有就创建

	// 2.2.1 获取name，image和amount
	// dish
	if dishId != 0 {
		dishEntity, _, er := svc.dishRepo.FindById(ctx, dishId)
		if er != nil {
			return er
		}
		shoppingCart = &cart.ShoppingCart{
			Name:       dishEntity.Name,
			Image:      dishEntity.Image,
			UserId:     userId,
			DishId:     dishId,
			SetmealId:  setId,
			DishFlavor: dto.DishFlavor,
			Number:     1,
			Amount:     dishEntity.Price,
			CreateTime: time.Now(),
		}
	} else if setId != 0 {
		setEntity, _, er := svc.setRepo.FindById(ctx, setId)
		if er != nil {
			return er
		}

		shoppingCart = &cart.ShoppingCart{
			Name:       setEntity.Name,
			Image:      setEntity.Image,
			UserId:     userId,
			SetmealId:  setId,
			DishId:     dishId,
			DishFlavor: dto.DishFlavor,
			Number:     1,
			Amount:     setEntity.Price,
			CreateTime: time.Now(),
		}
	}

	err = svc.cartRepo.Create(ctx, shoppingCart)
	return err

}
