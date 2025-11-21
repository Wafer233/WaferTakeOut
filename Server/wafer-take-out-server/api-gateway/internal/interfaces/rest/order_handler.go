package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	orderApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/order"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	svc *rpc.OrderService
}

func (h *OrderHandler) Submit(c *gin.Context) {
	var dto orderApp.SubmitDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}
	vo, err := h.svc.Submit(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))

}

func (h *OrderHandler) Payment(c *gin.Context) {

	var dto orderApp.PaymentDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定失败"))
		return
	}

	vo, err := h.svc.Payment(c, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *OrderHandler) ListPage(c *gin.Context) {

	var dto orderApp.UserPageDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	userID, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("无权限"))
		return
	}
	vo, err := h.svc.Page(ctx, &dto, userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *OrderHandler) ListAdminPage(c *gin.Context) {

	var dto orderApp.AdminPageDTO

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindPageAdmin(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *OrderHandler) CreateSame(c *gin.Context) {

	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	userID, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("无权限"))
		return
	}
	err := h.svc.CreateSame(ctx, int64(idInt), userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *OrderHandler) GetOrder(c *gin.Context) {

	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.GetOrder(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *OrderHandler) GetStatistics(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.GetStatistics(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *OrderHandler) UserCancel(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.UserCancel(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *OrderHandler) Rejection(c *gin.Context) {

	dto := orderApp.RejectionDTO{}

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.Rejection(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *OrderHandler) Confirm(c *gin.Context) {
	dto := orderApp.ConfirmDTO{}

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.Confirm(ctx, dto.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *OrderHandler) Cancel(c *gin.Context) {

	dto := orderApp.CancelDTO{}

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.Cancel(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *OrderHandler) Delivery(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.Delivery(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *OrderHandler) Complete(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.Complete(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func NewOrderHandler(svc *rpc.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}
