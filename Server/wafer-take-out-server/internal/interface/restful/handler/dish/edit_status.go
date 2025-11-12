package dishHandler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (h *DishHandler) EditDishStatus(c *gin.Context) {

	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("id读取失败"))
		return
	}

	status := c.Param("status")
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("status读取失败"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权无法获取id"))
		return
	}
	err = h.svc.StatusFlip(ctx, int64(idInt), statusInt, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务失败"))
		return
	}

	c.JSON(200, result.Success())

}
