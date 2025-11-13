package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewAdminRouter(r *gin.Engine, h *CommonHandler) *gin.Engine {

	common := r.Group("/admin/common")
	common.Use(middleware.JWTAuthMiddleware())
	common.POST("upload", h.Upload)
	return r
}
