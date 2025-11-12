package dishHandler

import (
	"context"
	"net/http"
	"time"

	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *DishHandler) GetDishedPaged(c *gin.Context) {

	dto := dishApp.PageDTO{}
	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer cancel()

	vo, err := h.svc.PageQuery(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}
