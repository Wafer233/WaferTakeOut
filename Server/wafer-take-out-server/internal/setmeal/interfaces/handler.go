package interfaces

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type SetMealHandler struct {
	svc *application.SetMealService
}

func NewSetMealHandler(svc *application.SetMealService) *SetMealHandler {
	return &SetMealHandler{
		svc: svc,
	}
}

func (h *SetMealHandler) Page(c *gin.Context) {
	dto := application.PageDTO{}
	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入有误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.PageQuery(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))

}

func (h *SetMealHandler) GetById(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("id错误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	vo, err := h.svc.IdQuery(ctx, int64(id))
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

	err = h.svc.StatusFlip(ctx, int64(id), curId.(int64), status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}

func (h *SetMealHandler) Update(c *gin.Context) {

	dto := application.AddSetMealDTO{}

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

	err := h.svc.Edit(ctx, &dto, curId.(int64))
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

	dto := application.AddSetMealDTO{}
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

	err = h.svc.AddSetMeal(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
