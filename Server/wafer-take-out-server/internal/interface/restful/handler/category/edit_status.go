package categoryHandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *CategoryHandler) EditCategoryStatus(c *gin.Context) {

	status := c.Param("status")
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		c.JSON(http.StatusOK, result.Error("获取status失败"))
		return
	}
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, result.Error("获取id失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}
	err = h.svc.FlipStatus(ctx, int64(idInt), statusInt, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
	}
	c.JSON(http.StatusOK, result.Success())
}
