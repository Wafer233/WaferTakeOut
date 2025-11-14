package interfaces

import (
	"context"
	"net/http"
	"strconv"
	"time"

	shopApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	svc *shopApp.ShopService
}

func NewShopHandler(svc *shopApp.ShopService) *ShopHandler {
	return &ShopHandler{svc: svc}
}

func (h *ShopHandler) Get(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	status, err := h.svc.FindStatus(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(status))
}

func (h *ShopHandler) Update(c *gin.Context) {

	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err = h.svc.UpdateStatus(ctx, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}
