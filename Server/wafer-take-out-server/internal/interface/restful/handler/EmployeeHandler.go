package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/auth"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	svc *employeeApp.EmployeeService
}

func NewEmployeeHandler(svc *employeeApp.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		svc: svc,
	}
}

func (h *EmployeeHandler) Login(c *gin.Context) {
	var dto employeeApp.LoginDTO
	var vo *employeeApp.LoginVO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("invalid request"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	vo, err = h.svc.Login(ctx, &dto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, result.Error(err.Error()))
		return
	}

	token, err := auth.GenerateToken(vo.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("生成Token失败"))
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

func (h *EmployeeHandler) AddEmployee(c *gin.Context) {
	var dto *employeeApp.AddEmployeeDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("invalid request"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	id, _ := c.Get("ID")

	err = h.svc.AddEmployee(ctx, dto, id.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.Success())

}

func (h *EmployeeHandler) Page(c *gin.Context) {
	var dto *employeeApp.PageDTO

	err := c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("invalid request"))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	vo, err := h.svc.PageQuery(ctx, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.SuccessData(vo))
}

func (h *EmployeeHandler) StatusFlip(c *gin.Context) {
	var dto *employeeApp.StatusFlipsDTO
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil || (status != 1 && status != 0) {
		c.JSON(http.StatusBadRequest, result.Error("status参数错误"))
		return
	}

	err = c.ShouldBindQuery(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("id参数错误"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	err = h.svc.StatusFlips(ctx, status, dto.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(err.Error()))
	}
	c.JSON(http.StatusOK, result.Success)
}
