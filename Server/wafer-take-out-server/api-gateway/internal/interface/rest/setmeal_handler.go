package rest

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	setmealApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/setmeal"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/persistence/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

type SetMealHandler struct {
	svc *rpc.SetMealService
}

func NewSetMealHandler(svc *rpc.SetMealService) *SetMealHandler {
	return &SetMealHandler{
		svc: svc,
	}
}

func (h *SetMealHandler) ListPage(c *gin.Context) {
	dto := setmealApp.PageDTO{}
	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	vo, err := h.svc.FindPage(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))

}

func (h *SetMealHandler) ListById(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("id错误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindById(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))

}

func (h *SetMealHandler) UpdateStatus(c *gin.Context) {

	// 这个地方前端有问题，无论怎么样他是0， 例如
	// "0?id=35"
	// 后端测试没问题
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("status获取失败"))
		return
	}

	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("id获取失败"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("无权限"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = h.svc.UpdateStatus(ctx, int64(id), curId.(int64), status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}

func (h *SetMealHandler) Update(c *gin.Context) {

	dto := setmealApp.SetMealDTO{}

	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("读取token失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	err := h.svc.Update(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func (h *SetMealHandler) Delete(c *gin.Context) {

	idsStr := c.Query("ids")
	idsArr := strings.Split(idsStr, ",")
	ids := make([]int64, 0)

	for _, value := range idsArr {
		valueInt, err := strconv.Atoi(value)
		ids = append(ids, int64(valueInt))
		if err != nil {
			c.JSON(http.StatusBadRequest, result.Error("输入错误"))
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := h.svc.Deletes(ctx, ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *SetMealHandler) Create(c *gin.Context) {

	dto := setmealApp.SetMealDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}
	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("无权限"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = h.svc.Create(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func (h *SetMealHandler) ListByCategoryId(c *gin.Context) {
	categoryIdStr := c.Query("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindByCategoryId(ctx, int64(categoryId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *SetMealHandler) ListDishById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.FindDishById(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))

}
