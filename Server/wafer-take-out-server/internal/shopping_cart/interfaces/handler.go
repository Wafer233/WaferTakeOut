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

func (h *ShoppingCartHandler) Create(c *gin.Context) {
	dto := application.AddDTO{}
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

	err = h.svc.Create(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

//func (h *ShoppingCartHandler) ListByUserId(c *gin.Context) {
//
//	userId := int64(1)
//	vo := h.svc.FindByUserId(ctx, userId)
//}
