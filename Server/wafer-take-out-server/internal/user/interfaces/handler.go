package interfaces

import (
	"context"
	"net/http"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *application.UserService
}

func NewUserHandler(svc *application.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) WxLogin(c *gin.Context) {
	dto := application.LoginDTO{}
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

	//这里小程序不支持cookie所以很烦....
	//只能在header里头验证了
	//c.SetCookie(
	//	"token",
	//	vo.Token,
	//	1*60*60,
	//	"/",
	//	"",
	//	false,
	//	false,
	//)

	c.JSON(http.StatusOK, result.SuccessData(vo))
}
