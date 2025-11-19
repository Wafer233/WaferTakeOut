package rest

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interface/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewCommonRouter(r *gin.Engine, h *CommonHandler) *gin.Engine {

	common := r.Group("/admin/common")
	common.Use(middleware.EmployeeAuthMiddleware())

	common.POST("upload", h.Upload)
	return r
}
