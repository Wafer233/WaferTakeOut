package dishHandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *DishHandler) GetDishesCategory(c *gin.Context) {

	categoryId := c.Query("categoryId")

	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("获取categoryId失败"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Second)
	defer cancel()

	records, err := h.svc.GetDishCategory(ctx, int64(categoryIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(records))
}
