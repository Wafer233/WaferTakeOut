package employeeHandler

import (
	"context"
	"net/http"
	"time"

	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *EmployeeHandler) AddEmployee(c *gin.Context) {
	var dto *employeeApp.AddEmployeeDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("invalid request"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, _ := c.Get("CurId")

	err = h.svc.AddEmployee(ctx, dto, id.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
