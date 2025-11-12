package setmealHandler

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *SetMealHandler) DeleteSetMeal(c *gin.Context) {

	idsStr := c.Query("ids")
	idsArr := strings.Split(idsStr, ",")
	ids := make([]int64, 0)

	for _, value := range idsArr {
		valueInt, err := strconv.Atoi(value)
		ids = append(ids, int64(valueInt))
		if err != nil {
			c.JSON(http.StatusBadRequest, result.Error("输入错误"))
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := h.svc.Deletes(ctx, ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}
