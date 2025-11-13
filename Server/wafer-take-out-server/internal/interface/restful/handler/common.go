package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	commonApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/common"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/result"
	"github.com/gin-gonic/gin"
)

type CommonHandler struct {
	svc *commonApp.CommonService
}

func NewCommonHandler(svc *commonApp.CommonService) *CommonHandler {
	return &CommonHandler{svc: svc}
}

func (h *CommonHandler) Upload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error("上传错误"))
		return
	}

	originalFilename := file.Filename
	suffix := filepath.Ext(originalFilename)
	if suffix == "" {
		suffix = ".jpg"
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), suffix)

	savePath := filepath.Join("C:\\Users\\wangz\\Documents\\wafer-take-out-pic", fileName)

	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Error("上传图片失败"))
		return
	}

	imgURL := fmt.Sprintf("http://localhost/media/%s", fileName)

	c.JSON(http.StatusOK, result.SuccessData(imgURL))
}
