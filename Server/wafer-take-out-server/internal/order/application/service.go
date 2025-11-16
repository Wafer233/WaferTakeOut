package application

import (
	"context"
	"errors"
	"time"

	addr "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/domain"
	orde "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/order/domain"
	cart "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/domain"
	"github.com/jinzhu/copier"
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

func (svc *OrderService) Page(ctx context.Context, dto *PageDTO,
	userId int64) (HistoryVO, error) {

	// 查历史
	records, total, err := svc.orderRepo.FindPage(ctx, dto.Page, dto.PageSize, userId, dto.Status)
	if err != nil || records == nil {
		return HistoryVO{}, err
	}
	// 查order detail
	ids := make([]int64, len(records))
	for i, record := range records {
		ids[i] = record.Id
	}

	detailsMap, err := svc.orderRepo.FindDetailByOrderIds(ctx, ids)

	//组装vo
	vo := make([]UserOrderVO, len(records))

	// 组装userorder
	for index, order := range records {
		tmpVO := UserOrderVO{}
		_ = copier.Copy(&tmpVO, &order)
		tmpVO.OrderTime = order.OrderTime.Format("2006-01-02 15:04:05")
		tmpVO.CheckoutTime = order.CheckoutTime.Format("2006-01-02 15:04:05")
		tmpVO.CancelTime = order.CancelTime.Format("2006-01-02 15:04:05")
		tmpVO.EstimatedDeliveryTime = order.EstimatedDeliveryTime.Format("2006-01-02 15:04:05")
		tmpVO.DeliveryTime = order.DeliveryTime.Format("2006-01-02 15:04:05")

		//组装特定的vo
		var detailVO []OrderDetail
		tmpDetail := detailsMap[order.Id]
		_ = copier.Copy(&detailVO, &tmpDetail)
		tmpVO.OrderDetails = detailVO

		vo[index] = tmpVO
	}

	history := HistoryVO{
		Total:   total,
		Records: vo,
	}

	return history, nil

}

func (svc *OrderService) GetOrder(ctx context.Context, orderId int64) (UserOrderVO, error) {

	order, err := svc.orderRepo.FindById(ctx, orderId)
	if err != nil || order == nil {
		return UserOrderVO{}, err
	}

	details, err := svc.orderRepo.FindDetailByOrderId(ctx, orderId)

	var detailVO []OrderDetail
	_ = copier.Copy(&detailVO, &details)

	var vo UserOrderVO
	_ = copier.Copy(&vo, &order)
	vo.OrderTime = order.OrderTime.Format("2006-01-02 15:04:05")
	vo.CheckoutTime = order.CheckoutTime.Format("2006-01-02 15:04:05")
	vo.CancelTime = order.CancelTime.Format("2006-01-02 15:04:05")
	vo.EstimatedDeliveryTime = order.EstimatedDeliveryTime.Format("2006-01-02 15:04:05")
	vo.DeliveryTime = order.DeliveryTime.Format("2006-01-02 15:04:05")
	vo.OrderDetails = detailVO

	return vo, nil

}

func (svc *OrderService) CreateSame(ctx context.Context, orderId int64, curID int64) error {
	// 根据订单id查询当前订单详情 detail: []*domain.OrderDetail
	details, err := svc.orderRepo.FindDetailByOrderId(ctx, orderId)
	if err != nil || len(details) == 0 {
		return err
	}

	// 将购物车对象批量添加到数据库 cart: []domain.ShoppingCart
	shopCart := make([]*cart.ShoppingCart, len(details))

	_ = copier.Copy(&shopCart, &details)
	for index, _ := range shopCart {
		shopCart[index].CreateTime = time.Now()
		shopCart[index].Id = 0
		shopCart[index].UserId = curID
		err = svc.cartRepo.Create(ctx, shopCart[index])
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *OrderService) Cancel(ctx context.Context, orderId int64) error {

	// 根据id查询订单
	order, err := svc.orderRepo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}
	// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消
	if order.Status > 2 {
		return errors.New("无法取消订单")
	}

	// 跳过微信流程
	if order.Status == 2 {
		order.PayStatus = REFUND
		err = svc.orderRepo.UpdateStatus(ctx, order)
		if err != nil {
			return err
		}
	}

	order.Status = CANCELLED
	order.CancelReason = "用户取消"
	order.CancelTime = time.Now()
	err = svc.orderRepo.UpdateStatus(ctx, order)
	if err != nil {
		return err
	}

	return nil
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
