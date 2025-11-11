package commonHandler

import commonApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/common"

type CommonHandler struct {
	svc *commonApp.CommonService
}

func NewCommonHandler(svc *commonApp.CommonService) *CommonHandler {
	return &CommonHandler{svc: svc}
}
