package application

import (
	"context"
	"time"

	addr "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/domain"
	orde "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/order/domain"
	cart "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/domain"
)

type OrderService struct {
	orderRepo orde.OrderRepository
	cartRepo  cart.ShoppingCartRepository
	addrRepo  addr.AddressRepository
}

func (svc *OrderService) Submit(ctx context.Context, dto *SubmitDTO, curId int64) (SubmitVO, error) {
	// 处理各类异常 （地址空，购物车空）
	address, err := svc.addrRepo.FindById(ctx, dto.AddressBookId)
	if err != nil || address == nil {
		return SubmitVO{}, err
	}

	shopCart, err := svc.cartRepo.Find(ctx, curId, 0, 0)
	if err != nil || len(shopCart) == 0 {
		return SubmitVO{}, err
	}

	// 插入1条数据
	estimateTime, _ := time.Parse("2006-01-02 15:04:05", dto.EstimatedDeliveryTime)
	orderEntity := &orde.Order{
		Id:                    0,
		Number:                time.Now().Format(time.RFC3339Nano),
		Status:                1,
		UserId:                curId,
		AddressBookId:         dto.AddressBookId,
		OrderTime:             time.Now(),
		CheckoutTime:          MYSQL_MIN_TIME,
		PayMethod:             dto.PayMethod,
		PayStatus:             0,
		Amount:                dto.Amount,
		Remark:                dto.Remark,
		Phone:                 address.Phone,
		Address:               address.DistrictName,
		UserName:              address.Consignee,
		Consignee:             address.Consignee,
		CancelReason:          "",
		RejectionReason:       "",
		CancelTime:            MYSQL_MIN_TIME,
		EstimatedDeliveryTime: estimateTime,
		DeliveryStatus:        dto.DeliveryStatus,
		DeliveryTime:          MYSQL_MIN_TIME,
		PackAmount:            dto.PackAmount,
		TableWareNumber:       dto.TablewareNumber,
		TableWareStatus:       dto.TablewareStatus,
	}

	err = svc.orderRepo.Create(ctx, orderEntity)
	if err != nil {
		return SubmitVO{}, err
	}

	// 插入n条detail

	details := make([]*orde.OrderDetail, len(shopCart))
	for i, _ := range shopCart {
		details[i] = &orde.OrderDetail{
			//Id: v.Id,
			Name:       shopCart[i].Name,
			Image:      shopCart[i].Image,
			OrderId:    orderEntity.Id,
			DishId:     shopCart[i].DishId,
			SetMealId:  shopCart[i].SetmealId,
			DishFlavor: shopCart[i].DishFlavor,
			Number:     shopCart[i].Number,
			Amount:     shopCart[i].Amount,
		}
	}
	err = svc.orderRepo.CreateDetail(ctx, details)

	// 清空购物车
	err = svc.cartRepo.Delete(ctx, curId)
	// 封装vo返回

	vo := SubmitVO{
		Id:          orderEntity.Id,
		OrderTime:   orderEntity.OrderTime.Format("2006-01-02 15:04:05"),
		OrderNumber: orderEntity.Number,
		OrderAmount: orderEntity.Amount,
	}
	return vo, nil
}

func (svc *OrderService) Payment(ctx context.Context, dto *PaymentDTO) (PaymentVO, error) {

	// 订单号查询
	order, err := svc.orderRepo.FindByNumber(ctx, dto.OrderNumber)
	if err != nil || order == nil {
		return PaymentVO{}, err
	}

	// 根据订单id更新订单的状态、支付方式、支付状态、结账时间
	order.Status = TO_BE_CONFIRMED
	order.PayStatus = PAID
	order.CheckoutTime = time.Now()
	err = svc.orderRepo.UpdateStatus(ctx, order)

	if err != nil {
		return PaymentVO{}, err
	}

	return PaymentVO{
		EstimatedDeliveryTime: order.EstimatedDeliveryTime.Format("2006-01-02 15:04:05"),
	}, nil

}

func NewOrderService(
	orderRepo orde.OrderRepository,
	cartRepo cart.ShoppingCartRepository,
	addrRepo addr.AddressRepository,
) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
		addrRepo:  addrRepo,
	}
}
