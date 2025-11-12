package dishHandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *DishHandler) GetDishId(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data, err := h.svc.GetDishId(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(data))

}
