package setmealHandler

import (
	"context"
	"net/http"
	"time"

	setmealApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/setmeal"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *SetMealHandler) GetSetMealsPaged(c *gin.Context) {
	dto := setmealApp.PageDTO{}
	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	vo, err := h.svc.PageQuery(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))

}
