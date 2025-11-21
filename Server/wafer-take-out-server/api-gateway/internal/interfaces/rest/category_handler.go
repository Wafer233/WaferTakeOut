package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	svc *rpc.CategoryService
}

func NewCategoryHandler(svc *rpc.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc: svc,
	}
}
func (h *CategoryHandler) Create(c *gin.Context) {

	var dto categoryApp.CategoryDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定JSON错误"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("获取curId失败"))
		return
	}
	err = h.svc.Create(ctx, &dto, id.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("获取id失败"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = h.svc.Delete(ctx, int64(idInt))

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务失败"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *CategoryHandler) Update(c *gin.Context) {

	dto := categoryApp.CategoryDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("绑定请求失败"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权，获取当前id失败"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = h.svc.Update(ctx, &dto, curId.(int64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务失败"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func (h *CategoryHandler) UpdateStatus(c *gin.Context) {

	status := c.Param("status")
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		c.JSON(http.StatusOK, result.Error("获取status失败"))
		return
	}
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, result.Error("获取id失败"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}
	err = h.svc.UpdateStatus(ctx, int64(idInt), statusInt, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *CategoryHandler) ListPage(c *gin.Context) {

	dto := categoryApp.PageDTO{}

	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	vo, err := h.svc.FindPage(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *CategoryHandler) ListByType(c *gin.Context) {
	curType := c.Query("type")
	curTypeInt, _ := strconv.Atoi(curType)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	vo, err := h.svc.FindByType(ctx, curTypeInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(vo))
}
