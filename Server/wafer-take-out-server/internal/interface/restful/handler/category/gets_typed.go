package categoryHandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *CategoryHandler) GetCategoriesTyped(c *gin.Context) {
	curType := c.Query("type")
	if curType == "" {
		c.JSON(http.StatusBadRequest, result.Error("获取总类失败"))
		return
	}
	curTypeInt, _ := strconv.Atoi(curType)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	vo, err := h.svc.TypeQuery(ctx, curTypeInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}
