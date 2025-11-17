package application

import (
	"context"
	"errors"
	"time"

	addr "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/internal/domain"
	orde "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/order/domain"
	cart "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/random"
	"github.com/jinzhu/copier"
)

type OrderService struct {
	orderRepo orde.OrderRepository
	cartRepo  cart.ShoppingCartRepository
	addrRepo  domain.AddressRepository
	domainSvc *orde.OrderDomainService
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
	now := time.Now()
	orderEntity := &orde.Order{
		Id:                    0,
		Number:                random.GenerateOrderID(),
		Status:                orde.PENDING_PAYMENT,
		UserId:                curId,
		AddressBookId:         dto.AddressBookId,
		OrderTime:             &now,
		CheckoutTime:          nil,
		PayMethod:             dto.PayMethod,
		PayStatus:             orde.UN_PAID,
		Amount:                dto.Amount,
		Remark:                dto.Remark,
		Phone:                 address.Phone,
		Address:               address.DistrictName,
		UserName:              address.Consignee,
		Consignee:             address.Consignee,
		CancelReason:          "",
		RejectionReason:       "",
		CancelTime:            nil,
		EstimatedDeliveryTime: &estimateTime,
		DeliveryStatus:        dto.DeliveryStatus,
		DeliveryTime:          nil,
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
	now := time.Now()
	order.Status = orde.TO_BE_CONFIRMED
	order.PayStatus = orde.PAID
	order.CheckoutTime = &now
	err = svc.orderRepo.UpdateStatus(ctx, order)

	if err != nil || order.EstimatedDeliveryTime == nil {
		return PaymentVO{}, err
	}

	estimateStr := order.EstimatedDeliveryTime.Format("2006-01-02 15:04:05")

	return PaymentVO{
		EstimatedDeliveryTime: estimateStr,
	}, nil

}

func (svc *OrderService) Page(ctx context.Context, dto *UserPageDTO,
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
		tmpVO.OrderTime = svc.domainSvc.ParseTime(order.OrderTime)
		tmpVO.CheckoutTime = svc.domainSvc.ParseTime(order.CheckoutTime)
		tmpVO.CancelTime = svc.domainSvc.ParseTime(order.CancelTime)
		tmpVO.EstimatedDeliveryTime = svc.domainSvc.ParseTime(order.EstimatedDeliveryTime)
		tmpVO.DeliveryTime = svc.domainSvc.ParseTime(order.DeliveryTime)

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

func (svc *OrderService) FindPageAdmin(ctx context.Context, dto *AdminPageDTO) (ListAdminOrderVO, error) {

	// 查历史
	layout := "2006-01-02 15:04:05"
	beginTime, _ := time.Parse(layout, dto.BeginTime)
	endTime, _ := time.Parse(layout, dto.EndTime)

	records, total, err := svc.orderRepo.FindPageAdmin(ctx, beginTime, endTime, dto.Number,
		dto.Page, dto.PageSize, dto.Phone, dto.Status)

	if err != nil || records == nil {
		return ListAdminOrderVO{}, err
	}
	// 查order detail
	//ids := make([]int64, len(records))
	//for i, record := range records {
	//	ids[i] = record.Id
	//}

	//detailsMap, err := svc.orderRepo.FindDetailByOrderIds(ctx, ids)

	//组装vo
	vo := make([]AdminOrderVO, len(records))

	// 组装userorder
	for index, order := range records {
		tmpVO := AdminOrderVO{}
		_ = copier.Copy(&tmpVO, &order)

		tmpVO.OrderTime = svc.domainSvc.ParseTime(order.OrderTime)
		tmpVO.CheckoutTime = svc.domainSvc.ParseTime(order.CheckoutTime)
		tmpVO.CancelTime = svc.domainSvc.ParseTime(order.CancelTime)
		tmpVO.EstimatedDeliveryTime = svc.domainSvc.ParseTime(order.EstimatedDeliveryTime)
		tmpVO.DeliveryTime = svc.domainSvc.ParseTime(order.DeliveryTime)
		vo[index] = tmpVO

		////组装特定的vo
		//var detailVO []OrderDetail
		//tmpDetail := detailsMap[order.Id]
		//_ = copier.Copy(&detailVO, &tmpDetail)
		//tmpVO.OrderDetails = detailVO
		//
		//vo[index].
	}

	list := ListAdminOrderVO{
		Total:  total,
		Orders: vo,
	}

	return list, nil

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

	vo.OrderTime = svc.domainSvc.ParseTime(order.OrderTime)
	vo.CheckoutTime = svc.domainSvc.ParseTime(order.CheckoutTime)
	vo.CancelTime = svc.domainSvc.ParseTime(order.CancelTime)
	vo.EstimatedDeliveryTime = svc.domainSvc.ParseTime(order.EstimatedDeliveryTime)
	vo.DeliveryTime = svc.domainSvc.ParseTime(order.DeliveryTime)
	vo.OrderDetails = detailVO

	return vo, nil

}

func (svc *OrderService) UserCancel(ctx context.Context, orderId int64) error {
	// 根据id查询订单
	order, err := svc.orderRepo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}
	//// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消
	if order.Status > 3 {
		return errors.New("无法取消")
	}

	if order.Status != orde.PENDING_PAYMENT {
		order.PayStatus = orde.REFUND
	}

	now := time.Now()
	order.CancelTime = &now
	order.CancelReason = "用户取消"
	order.Status = orde.CANCELLED
	err = svc.orderRepo.UpdateStatus(ctx, order)

	return err
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

func (svc *OrderService) Confirm(ctx context.Context, orderId int64) error {
	// 根据id查询订单
	order, err := svc.orderRepo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}

	order.Status = orde.CONFIRMED
	err = svc.orderRepo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) Rejection(ctx context.Context, dto *RejectionDTO) error {
	// 根据id查询订单
	order, err := svc.orderRepo.FindById(ctx, dto.Id)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}

	// 跳过微信流程
	order.PayStatus = orde.REFUND
	order.Status = orde.CANCELLED
	order.RejectionReason = dto.RejectionReason
	now := time.Now()
	order.CancelTime = &now
	err = svc.orderRepo.UpdateStatus(ctx, order)

	return err

}

func (svc *OrderService) Cancel(ctx context.Context, dto *CancelDTO) error {

	// 根据id查询订单
	order, err := svc.orderRepo.FindById(ctx, dto.Id)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}
	// 跳过微信流程

	order.PayStatus = orde.REFUND
	order.Status = orde.CANCELLED
	order.CancelReason = dto.CancelReason
	now := time.Now()
	order.CancelTime = &now
	err = svc.orderRepo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) Delivery(ctx context.Context, orderId int64) error {
	// 根据id查询订单
	order, err := svc.orderRepo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}

	order.Status = orde.DELIVERY_IN_PROGRESS
	if order.DeliveryStatus == 1 {
		now := time.Now()
		order.DeliveryTime = &now
	} else {
		order.DeliveryTime = order.EstimatedDeliveryTime
	}
	err = svc.orderRepo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) Complete(ctx context.Context, orderId int64) error {
	// 根据id查询订单
	order, err := svc.orderRepo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}
	// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消

	order.Status = orde.COMPLETED
	now := time.Now()
	order.DeliveryTime = &now
	err = svc.orderRepo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) GetStatistics(ctx context.Context) (StatisticsVO, error) {

	confirmed, deliveryInProgress, toBeConfirmed, err := svc.orderRepo.GetTotalByStatus(ctx)

	if err != nil {
		return StatisticsVO{}, err
	}
	return StatisticsVO{
		Confirmed:          confirmed,
		DeliveryInProgress: deliveryInProgress,
		ToBeConfirmed:      toBeConfirmed,
	}, nil
}

func NewOrderService(
	orderRepo orde.OrderRepository,
	cartRepo cart.ShoppingCartRepository,
	addrRepo domain.AddressRepository,
) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
		addrRepo:  addrRepo,
	}
}
