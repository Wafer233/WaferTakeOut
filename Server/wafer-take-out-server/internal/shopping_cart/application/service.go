package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/domain"
)

type ShoppingCartService struct {
	repo domain.ShoppingCartRepository
}

func NewShoppingCartService(repo domain.ShoppingCartRepository) *ShoppingCartService {
	return &ShoppingCartService{
		repo: repo,
	}
}

func (svc *ShoppingCartService) Create(ctx context.Context, dto *AddDTO) error {

	// 这里因为他不支持cookie我暂时不搞userid
	// TODO 有时间搞没时间算了
	userId := int64(1)

	cart := &domain.ShoppingCart{

		Name:       "", //冗余字段
		Image:      "", //冗余字段
		UserId:     userId,
		DishId:     dto.DishId,
		SetmealId:  dto.SetMealId,
		DishFlavor: dto.DishFlavor,
		Number:     1,          //就是1
		Amount:     float64(0), //冗余字段
		CreateTime: time.Now(),
	}

	err := svc.repo.Create(ctx, cart)

	return err
}
