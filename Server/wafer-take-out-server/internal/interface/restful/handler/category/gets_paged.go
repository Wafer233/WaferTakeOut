package categoryHandler

import (
	"context"
	"net/http"
	"time"

	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *CategoryHandler) GetCategories(c *gin.Context) {

	dto := categoryApp.PageDTO{}

	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	vo, err := h.svc.PageQuery(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}
