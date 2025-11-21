package rest

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

type DishHandler struct {
	svc *rpc.DishService
}

func NewDishHandler(svc *rpc.DishService) *DishHandler {
	return &DishHandler{svc: svc}
}

func (h *DishHandler) Page(c *gin.Context) {

	dto := dishApp.PageDTO{}
	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	vo, err := h.svc.FindPage(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *DishHandler) ListByCategory(c *gin.Context) {

	categoryId := c.Query("categoryId")

	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("获取categoryId失败"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Second)
	defer cancel()

	records, err := h.svc.FindByCategoryId(ctx, int64(categoryIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(records))
}

func (h *DishHandler) GetById(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data, err := h.svc.FindById(ctx, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.SuccessData(data))

}

func (h *DishHandler) UpdateStatus(c *gin.Context) {

	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("id读取失败"))
		return
	}

	status := c.Param("status")
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("status读取失败"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权无法获取id"))
		return
	}
	err = h.svc.UpdateStatus(ctx, int64(idInt), statusInt, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务失败"))
		return
	}

	c.JSON(200, result.Success())

}

func (h *DishHandler) Update(c *gin.Context) {

	dto := dishApp.DishDTO{}

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = h.svc.Update(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func (h *DishHandler) Delete(c *gin.Context) {
	ids := c.Query("ids")

	idStrings := strings.Split(ids, ",")
	idArr := make([]int64, len(idStrings))

	for _, value := range idStrings {
		id, err := strconv.Atoi(value)
		if err != nil {
			c.JSON(http.StatusOK, result.Error("输入错误"))
			return
		}
		idArr = append(idArr, int64(id))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := h.svc.Delete(ctx, idArr)
	if err != nil {
		c.JSON(http.StatusOK, result.Error("调用服务失败"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}

func (h *DishHandler) Create(c *gin.Context) {

	dto := dishApp.DishDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("输入错误"))
		return
	}
	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = h.svc.Create(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("调用服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success())
}

func (h *DishHandler) ListByCategoryIdFlavor(c *gin.Context) {
	categoryId := c.Query("categoryId")
	categoryIdInt, _ := strconv.Atoi(categoryId)

	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Second)
	defer cancel()

	vo, err := h.svc.FindByCategoryIdFlavor(ctx, int64(categoryIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))

}
