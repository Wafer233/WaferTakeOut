package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	addressbookApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/addressbook"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/persistence/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	svc *rpc.AddressService
}

func NewAddressHandler(svc *rpc.AddressService) *AddressHandler {
	return &AddressHandler{svc: svc}
}

func (h *AddressHandler) Create(c *gin.Context) {

	dto := addressbookApp.AddressDTO{}
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

func (h *AddressHandler) GetDefault(c *gin.Context) {
	userId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindDefault(ctx, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *AddressHandler) UpdateDefault(c *gin.Context) {
	userId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}
	var addrId addressbookApp.DefaultIdDTO
	err := c.ShouldBindJSON(&addrId)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定错误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = h.svc.UpdateDefault(ctx, userId.(int64), addrId.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *AddressHandler) GetById(c *gin.Context) {

	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindById(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *AddressHandler) Delete(c *gin.Context) {

	idStr := c.Query("id")
	idInt, _ := strconv.Atoi(idStr)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := h.svc.DeleteById(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *AddressHandler) Update(c *gin.Context) {

	dto := addressbookApp.AddressDTO{}
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

	err = h.svc.Update(ctx, &dto, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}
