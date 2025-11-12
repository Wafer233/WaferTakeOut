package employeeHandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *EmployeeHandler) StatusFlip(c *gin.Context) {
	var dto *employeeApp.StatusFlipsDTO
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)

	if err != nil || (status != 1 && status != 0) {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}

	err = c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("没有权限"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	err = h.svc.StatusFlips(ctx, status, dto.ID, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
	}
	c.JSON(http.StatusOK, result.Success)
}
