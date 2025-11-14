package interfaces

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/jwt"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	svc *application.EmployeeService
}

func NewEmployeeHandler(svc *application.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		svc: svc,
	}
}

func (h *EmployeeHandler) Login(c *gin.Context) {
	var dto application.LoginDTO
	var vo *application.LoginVO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("请求错误"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	vo, err = h.svc.Login(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	token, err := jwt.GenerateToken(vo.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, result.Error("未授权"))
		return
	}

	c.SetCookie(
		"token",
		token,
		1*60*60,
		"/",
		"",
		false,
		false,
	)
	vo.Token = token
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
	var dto application.PageDTO

	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("错误请求"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	vo, err := h.svc.FindPage(ctx, &dto)
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

	vo, err := h.svc.FindById(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))

}

func (h *EmployeeHandler) UpdateStatus(c *gin.Context) {
	var dto *application.StatusFlipsDTO
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

	err = h.svc.UpdateStatus(ctx, status, dto.ID, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
	}
	c.JSON(http.StatusOK, result.Success)
}

func (h *EmployeeHandler) Update(c *gin.Context) {

	dto := &application.AddEmployeeDTO{}
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
	err = h.svc.Update(ctx, dto, id.(int64))

	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var dto *application.AddEmployeeDTO

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

	err = h.svc.Create(ctx, dto, id.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("内部服务错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}

func (h *EmployeeHandler) UpdatePassword(c *gin.Context) {
	//想骂人了这里根本没有传入empID,我在token自己获取了
	var dto application.PasswordDTO
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

	err = h.svc.UpdatePassword(ctx, &dto, curId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("密码错误或内部错误"))
		return
	}

	c.JSON(http.StatusOK, result.Success())
}
