package dishHandler

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *DishHandler) DeleteDishes(c *gin.Context) {
	ids := c.Query("ids")

	idStrings := strings.Split(ids, ",")
	idArr := make([]int64, len(idStrings))

	for _, value := range idStrings {
		id, err := strconv.Atoi(value)
		if err != nil {
			c.JSON(http.StatusOK, result.Error("输入错误"))
			return
		}
		idArr = append(idArr, int64(id))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := h.svc.DeleteDishes(ctx, idArr)
	if err != nil {
		c.JSON(http.StatusOK, result.Error("调用服务失败"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}
