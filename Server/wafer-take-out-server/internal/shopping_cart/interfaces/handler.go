package interfaces

import (
	"context"
	"net/http"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type ShoppingCartHandler struct {
	svc *application.ShoppingCartService
}

func NewShoppingCartHandler(svc *application.ShoppingCartService) *ShoppingCartHandler {
	return &ShoppingCartHandler{
		svc: svc,
	}
}

// Add
// 如果从购物车界面添加的话，前端不回有preload
func (h *ShoppingCartHandler) Add(c *gin.Context) {
	dto := application.CartDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求信息有误"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = h.svc.Add(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func (h *ShoppingCartHandler) List(c *gin.Context) {

	userId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindByUserId(ctx, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *ShoppingCartHandler) Sub(c *gin.Context) {
	//
	//dto := application.CartDTO{}
	//err := c.ShouldBindJSON(&dto)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, result.Error("错误请求"))
	//	return
	//}
	//
	//curId, exist := c.Get("CurID")
	//if !exist {
	//	c.JSON(http.StatusUnauthorized, result.Error("未授权"))
	//	return
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//
	//err = h.svc.Sub(ctx, &dto, curId.(int64))
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
	//	return
	//}

	c.JSON(http.StatusOK, result.Success())
}

func (h *ShoppingCartHandler) Delete(c *gin.Context) {

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.Delete(ctx, curId.(int64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}
