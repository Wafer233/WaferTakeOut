package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *AddressHandler) *gin.Engine {

	user := r.Group("/user")
	user.Use(middleware.UserAuthMiddleware())
	user.POST("/addressBook", h.Create)
	return r
}
