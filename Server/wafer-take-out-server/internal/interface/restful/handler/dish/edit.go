package dishHandler

import (
	"context"
	"net/http"
	"time"

	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *DishHandler) EditDish(c *gin.Context) {

	dto := dishApp.DishDTO{}

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = h.svc.UpdateDish(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
