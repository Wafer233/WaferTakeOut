package rest

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interfaces/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewAddressBookRouter(r *gin.Engine, h *AddressHandler) *gin.Engine {

	user := r.Group("/user/addressBook")
	user.Use(middleware.UserAuthMiddleware())

	user.POST("", h.Create)
	user.GET("/list", h.List)
	user.GET("/default", h.GetDefault)
	user.PUT("/default", h.UpdateDefault)
	user.GET("/:id", h.GetById)
	user.DELETE("", h.Delete)
	user.PUT("", h.Update)
	return r
}
