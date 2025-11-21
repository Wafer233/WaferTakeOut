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
	"go.uber.org/zap"
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
		zap.L().Error("错误请求")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	vo, err := h.client.Login(ctx, &dto)
	if err != nil {
		zap.L().Error("调用Login微服务错误")
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

	zap.L().Info("登录成功")
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
	zap.L().Info("登出成功")
	c.JSON(http.StatusOK, result.Success())
}

func (h *EmployeeHandler) ListPage(c *gin.Context) {
	var dto emplApp.PageDTO

	err := c.ShouldBindQuery(&dto)
	if err != nil {
		zap.L().Error("错误请求")
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	vo, err := h.client.FindPage(ctx, &dto)
	if err != nil {
		zap.L().Error("调用微服务错误")
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	zap.L().Info("获取员工页面成功")
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *EmployeeHandler) List(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		zap.L().Error("错误请求")
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	vo, err := h.client.FindById(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		zap.L().Error("调用rpc服务错误")
		return
	}

	zap.L().Info("获取Id列表信息成功")
	c.JSON(http.StatusOK, result.SuccessData(vo))

}

func (h *EmployeeHandler) UpdateStatus(c *gin.Context) {
	var dto *emplApp.StatusFlipsDTO
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)

	if err != nil || (status != 1 && status != 0) {
		zap.L().Error("错误请求")
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}

	err = c.ShouldBindQuery(&dto)
	if err != nil {
		zap.L().Error("错误请求")
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}

	curId, exist := c.Get("CurID")
	if !exist {
		zap.L().Error("没有授权")
		c.JSON(http.StatusUnauthorized, result.Error("没有权限"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	err = h.client.UpdateStatus(ctx, status, dto.ID, curId.(int64))
	if err != nil {
		zap.L().Error("调用微服务错误")
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
	}

	zap.L().Info("更新状态成功")
	c.JSON(http.StatusOK, result.Success)
}

func (h *EmployeeHandler) Update(c *gin.Context) {

	dto := &emplApp.AddEmployeeDTO{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		zap.L().Error("错误请求")
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, exist := c.Get("CurID")
	if !exist {
		zap.L().Error("获取权限失败，无法在上下文中获取当前ID")
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	err = h.client.Update(ctx, dto, id.(int64))

	if err != nil {
		zap.L().Error("调用rpc服务错误")
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	zap.L().Info("更新成功")
	c.JSON(http.StatusOK, result.Success())

}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var dto *emplApp.AddEmployeeDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		zap.L().Error("错误请求")
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, exist := c.Get("CurID")
	if !exist {
		zap.L().Error("获取当前ID失败")
		c.JSON(http.StatusUnauthorized, result.Error("未授权导致无法读取CurId"))
		return
	}

	err = h.client.Create(ctx, dto, id.(int64))
	if err != nil {
		zap.L().Error("内部服务错误")
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	zap.L().Info("创建员工成功")
	c.JSON(http.StatusOK, result.Success())
}

func (h *EmployeeHandler) UpdatePassword(c *gin.Context) {
	//想骂人了这里根本没有传入empID,我在token自己获取了
	var dto emplApp.PasswordDTO
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		zap.L().Error("错误请求")
		c.JSON(http.StatusInternalServerError, result.Error("绑定失败"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	curId, exist := c.Get("CurID")
	if !exist {
		zap.L().Error("获取当前Id失败")
		c.JSON(http.StatusUnauthorized, result.Error("无权限"))
		return
	}

	err = h.client.UpdatePassword(ctx, &dto, curId.(int64))
	if err != nil {
		zap.L().Error("内部服务错误")
		c.JSON(http.StatusInternalServerError, result.Error("密码错误或内部错误"))
		return
	}

	zap.L().Info("修改密码成功")
	c.JSON(http.StatusOK, result.Success())
}
