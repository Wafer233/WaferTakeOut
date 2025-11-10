package employeeHandler

import (
	"context"
	"net/http"
	"time"

	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/auth"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *EmployeeHandler) Login(c *gin.Context) {
	var dto employeeApp.LoginDTO
	var vo *employeeApp.LoginVO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("invalid request"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	vo, err = h.svc.Login(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, result.Error(err.Error()))
		return
	}

	token, err := auth.GenerateToken(vo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("生成Token失败"))
		return
	}

	c.SetCookie(
		"token",
		token,
		1*60*60,
		"/",
		"",
		false,
		false,
	)
	vo.Token = token
	c.JSON(http.StatusOK, result.SuccessData(vo))
}
