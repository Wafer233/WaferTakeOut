package categoryHandler

import (
	"context"
	"net/http"
	"time"

	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *CategoryHandler) EditCategory(c *gin.Context) {

	dto := categoryApp.EditCategoryDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定请求失败"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权，获取当前id失败"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()

	err = h.svc.EditCategory(ctx, &dto, curId.(int64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务失败"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
