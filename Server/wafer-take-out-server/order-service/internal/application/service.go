package application

import (
	"context"
	"errors"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/domain"
	rpc2 "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/pkg/random"
	"github.com/jinzhu/copier"
)

type OrderService struct {
	repo      domain.OrderRepository
	domainSvc *domain.OrderDomainService
	addrSvc   *rpc2.AddressBookService
	cartSvc   *rpc2.ShoppingCartService
}

func (svc *OrderService) Submit(ctx context.Context, dto *SubmitDTO, curId int64) (SubmitVO, error) {
	// 处理各类异常 （地址空，购物车空）
	address, err := svc.addrSvc.FindById(ctx, dto.AddressBookId)
	if err != nil || address == nil {
		return SubmitVO{}, err
	}

	shopCarts, err := svc.cartSvc.FindByUserId(ctx, curId)
	if err != nil || len(shopCarts) == 0 {
		return SubmitVO{}, err
	}

	// 插入1条数据
	estimateTime, _ := time.Parse("2006-01-02 15:04:05", dto.EstimatedDeliveryTime)
	now := time.Now()
	orderEntity := &domain.Order{
		Id:                    0,
		Number:                random.GenerateOrderID(),
		Status:                domain.PENDING_PAYMENT,
		UserId:                curId,
		AddressBookId:         dto.AddressBookId,
		OrderTime:             &now,
		CheckoutTime:          nil,
		PayMethod:             dto.PayMethod,
		PayStatus:             domain.UN_PAID,
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

	err = svc.repo.Create(ctx, orderEntity)
	if err != nil {
		return SubmitVO{}, err
	}

	// 插入n条detail

	details := make([]*domain.OrderDetail, len(shopCarts))
	for i, _ := range shopCarts {
		details[i] = &domain.OrderDetail{
			//Id: v.Id,
			Name:       shopCarts[i].Name,
			Image:      shopCarts[i].Image,
			OrderId:    orderEntity.Id,
			DishId:     shopCarts[i].DishId,
			SetMealId:  shopCarts[i].SetMealId,
			DishFlavor: shopCarts[i].DishFlavor,
			Number:     shopCarts[i].Number,
			Amount:     shopCarts[i].Amount,
		}
	}
	err = svc.repo.CreateDetail(ctx, details)

	// 清空购物车
	err = svc.cartSvc.Delete(ctx, curId)
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
	order, err := svc.repo.FindByNumber(ctx, dto.OrderNumber)
	if err != nil || order == nil {
		return PaymentVO{}, err
	}

	// 根据订单id更新订单的状态、支付方式、支付状态、结账时间
	now := time.Now()
	order.Status = domain.TO_BE_CONFIRMED
	order.PayStatus = domain.PAID
	order.CheckoutTime = &now
	err = svc.repo.UpdateStatus(ctx, order)

	if err != nil || order.EstimatedDeliveryTime == nil {
		return PaymentVO{}, err
	}

	estimateStr := order.EstimatedDeliveryTime.Format("2006-01-02 15:04:05")

	return PaymentVO{
		EstimatedDeliveryTime: estimateStr,
	}, nil

}

func (svc *OrderService) Page(ctx context.Context, dto *UserPageDTO,
	userId int64) (UserPageVO, error) {

	// 查历史
	records, total, err := svc.repo.FindPage(ctx, dto.Page, dto.PageSize, userId, dto.Status)
	if err != nil || records == nil {
		return UserPageVO{}, err
	}
	// 查order detail
	ids := make([]int64, len(records))
	for i, record := range records {
		ids[i] = record.Id
	}

	detailsMap, err := svc.repo.FindDetailByOrderIds(ctx, ids)

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

	history := UserPageVO{
		Total:   total,
		Records: vo,
	}

	return history, nil

}

func (svc *OrderService) FindPageAdmin(ctx context.Context, dto *AdminPageDTO) (AdminPageVO, error) {

	// 查历史
	layout := "2006-01-02 15:04:05"
	beginTime, _ := time.Parse(layout, dto.BeginTime)
	endTime, _ := time.Parse(layout, dto.EndTime)

	records, total, err := svc.repo.FindPageAdmin(ctx, beginTime, endTime, dto.Number,
		dto.Page, dto.PageSize, dto.Phone, dto.Status)

	if err != nil || records == nil {
		return AdminPageVO{}, err
	}
	// 查order detail
	//ids := make([]int64, len(records))
	//for i, record := range records {
	//	ids[i] = record.Id
	//}

	//detailsMap, err := svc.repo.FindDetailByOrderIds(ctx, ids)

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

	list := AdminPageVO{
		Total:   total,
		Records: vo,
	}

	return list, nil

}

func (svc *OrderService) GetOrder(ctx context.Context, orderId int64) (UserOrderVO, error) {

	order, err := svc.repo.FindById(ctx, orderId)
	if err != nil || order == nil {
		return UserOrderVO{}, err
	}

	details, err := svc.repo.FindDetailByOrderId(ctx, orderId)

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
	order, err := svc.repo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}
	//// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消
	if order.Status > 3 {
		return errors.New("无法取消")
	}

	if order.Status != domain.PENDING_PAYMENT {
		order.PayStatus = domain.REFUND
	}

	now := time.Now()
	order.CancelTime = &now
	order.CancelReason = "用户取消"
	order.Status = domain.CANCELLED
	err = svc.repo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) CreateSame(ctx context.Context, orderId int64, curID int64) error {
	// 根据订单id查询当前订单详情 detail: []*domain.OrderDetail
	details, err := svc.repo.FindDetailByOrderId(ctx, orderId)
	if err != nil || len(details) == 0 {
		return err
	}

	// 将购物车对象批量添加到数据库 cart: []domain.ShoppingCart
	shopCarts := make([]*rpc2.CartDTO, len(details))

	_ = copier.Copy(&shopCarts, &details)
	for index, _ := range shopCarts {
		shopCarts[index].DishFlavor = details[index].DishFlavor
		shopCarts[index].DishId = details[index].DishId
		shopCarts[index].SetMealId = details[index].SetMealId
		err = svc.cartSvc.Add(ctx, shopCarts[index], curID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *OrderService) Confirm(ctx context.Context, orderId int64) error {
	// 根据id查询订单
	order, err := svc.repo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}

	order.Status = domain.CONFIRMED
	err = svc.repo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) Rejection(ctx context.Context, dto *RejectionDTO) error {
	// 根据id查询订单
	order, err := svc.repo.FindById(ctx, dto.Id)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}

	// 跳过微信流程
	order.PayStatus = domain.REFUND
	order.Status = domain.CANCELLED
	order.RejectionReason = dto.RejectionReason
	now := time.Now()
	order.CancelTime = &now
	err = svc.repo.UpdateStatus(ctx, order)

	return err

}

func (svc *OrderService) Cancel(ctx context.Context, dto *CancelDTO) error {

	// 根据id查询订单
	order, err := svc.repo.FindById(ctx, dto.Id)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}
	// 跳过微信流程

	order.PayStatus = domain.REFUND
	order.Status = domain.CANCELLED
	order.CancelReason = dto.CancelReason
	now := time.Now()
	order.CancelTime = &now
	err = svc.repo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) Delivery(ctx context.Context, orderId int64) error {
	// 根据id查询订单
	order, err := svc.repo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}

	order.Status = domain.DELIVERY_IN_PROGRESS
	if order.DeliveryStatus == 1 {
		now := time.Now()
		order.DeliveryTime = &now
	} else {
		order.DeliveryTime = order.EstimatedDeliveryTime
	}
	err = svc.repo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) Complete(ctx context.Context, orderId int64) error {
	// 根据id查询订单
	order, err := svc.repo.FindById(ctx, orderId)

	// 校验订单是否存在
	if err != nil || order == nil {
		return err
	}
	// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消

	order.Status = domain.COMPLETED
	now := time.Now()
	order.DeliveryTime = &now
	err = svc.repo.UpdateStatus(ctx, order)

	return err
}

func (svc *OrderService) GetStatistics(ctx context.Context) (StatisticsVO, error) {

	confirmed, deliveryInProgress, toBeConfirmed, err := svc.repo.GetTotalByStatus(ctx)

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
	repo domain.OrderRepository,
	domainSvc *domain.OrderDomainService,
	addrSvc *rpc2.AddressBookService,
	cartSvc *rpc2.ShoppingCartService,
) *OrderService {
	return &OrderService{
		repo:      repo,
		domainSvc: domainSvc,
		addrSvc:   addrSvc,
		cartSvc:   cartSvc,
	}
}
