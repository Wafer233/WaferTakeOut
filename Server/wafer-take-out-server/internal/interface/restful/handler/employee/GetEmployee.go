package employeeHandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	vo, err := h.svc.GetEmployee(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))

}
