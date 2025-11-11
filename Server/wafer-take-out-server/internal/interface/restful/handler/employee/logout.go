package employeeHandler

import (
	"net/http"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

func (h *EmployeeHandler) Logout(c *gin.Context) {

	c.SetCookie(
		"token",
		"",
		-1,
		"/",
		"",
		false,
		false,
	)
	c.JSON(http.StatusOK, result.Success())
}
