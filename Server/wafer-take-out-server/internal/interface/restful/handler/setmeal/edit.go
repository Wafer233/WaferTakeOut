package setmealHandler

import (
	"context"
	"net/http"
	"time"

	setmealApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/setmeal"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *SetMealHandler) EditSetMeal(c *gin.Context) {

	dto := setmealApp.AddSetMealDTO{}

	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("读取token失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	err := h.svc.Edit(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
