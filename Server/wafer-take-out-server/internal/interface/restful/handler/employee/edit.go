package employeeHandler

import (
	"context"
	"net/http"
	"time"

	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *EmployeeHandler) EditEmployee(c *gin.Context) {

	dto := &employeeApp.AddEmployeeDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	err = h.svc.UpdateEmployee(ctx, dto, id.(int64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}
