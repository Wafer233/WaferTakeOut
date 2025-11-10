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
		c.JSON(http.StatusBadRequest, result.Error(err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusInternalServerError, result.Error("没有登录验证token，无法识别id"))
		return
	}
	err = h.svc.UpdateEmployee(ctx, dto, id.(int64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}
