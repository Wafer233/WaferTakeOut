package rest

import (
	"context"
	"net/http"
	"time"

	userApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/user"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *rpc.UserService
}

func NewUserHandler(svc *rpc.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) WxLogin(c *gin.Context) {
	dto := userApp.LoginDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	vo, err := h.svc.WxLogin(ctx, dto.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用内部服务失败"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}
