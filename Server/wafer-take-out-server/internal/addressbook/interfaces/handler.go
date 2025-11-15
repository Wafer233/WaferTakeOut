package interfaces

import (
	"context"
	"net/http"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	svc *application.AddressService
}

func NewAddressHandler(svc *application.AddressService) *AddressHandler {
	return &AddressHandler{svc: svc}
}

func (h *AddressHandler) Create(c *gin.Context) {

	dto := application.AddressDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定错误"))
		return
	}

	userId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = h.svc.Create(ctx, &dto, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *AddressHandler) List(c *gin.Context) {

	userId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindByUserId(ctx, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}
