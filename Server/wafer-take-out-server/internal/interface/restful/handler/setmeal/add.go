package setmealHandler

import (
	"context"
	"net/http"
	"time"

	setmealApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/setmeal"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *SetMealHandler) AddSetMeal(c *gin.Context) {

	dto := setmealApp.AddSetMealDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}
	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("无权限"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = h.svc.AddSetMeal(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
