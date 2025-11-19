package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	emplApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/persistence/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	client *rpc.EmployeeService
}

func NewEmployeeHandler(client *rpc.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		client: client,
	}
}

func (h *EmployeeHandler) Login(c *gin.Context) {
	var dto emplApp.LoginDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	vo, err := h.client.Login(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.SetCookie(
		"token",
		vo.Token,
		1*60*60,
		"/",
		"",
		false,
		false,
	)

	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *EmployeeHandler) Logout(c *gin.Context) {

	c.SetCookie(
		"token",
		"",
		-1,
		"/",
		"",
		false,
		false,
	)
	c.JSON(http.StatusOK, result.Success())
}

func (h *EmployeeHandler) ListPage(c *gin.Context) {
	var dto emplApp.PageDTO

	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	vo, err := h.client.FindPage(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *EmployeeHandler) List(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	vo, err := h.client.FindById(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))

}

func (h *EmployeeHandler) UpdateStatus(c *gin.Context) {
	var dto *emplApp.StatusFlipsDTO
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)

	if err != nil || (status != 1 && status != 0) {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}

	err = c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("没有权限"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	err = h.client.UpdateStatus(ctx, status, dto.ID, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
	}
	c.JSON(http.StatusOK, result.Success)
}

func (h *EmployeeHandler) Update(c *gin.Context) {

	dto := &emplApp.AddEmployeeDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	err = h.client.Update(ctx, dto, id.(int64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var dto *emplApp.AddEmployeeDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("未授权导致无法读取CurId"))
		return
	}

	err = h.client.Create(ctx, dto, id.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func (h *EmployeeHandler) UpdatePassword(c *gin.Context) {
	//想骂人了这里根本没有传入empID,我在token自己获取了
	var dto emplApp.PasswordDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("绑定失败"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	curId, exist := c.Get("CurID")
	if !exist {
		c.JSON(http.StatusUnauthorized, result.Error("无权限"))
		return
	}

	err = h.client.UpdatePassword(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("密码错误或内部错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
