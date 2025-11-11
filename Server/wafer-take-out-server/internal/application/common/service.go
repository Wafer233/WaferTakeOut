package commonApp

import commonInfra "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/common"

type CommonService struct {
	repo *commonInfra.CommonRepository
}

func NewCommonService(repo *commonInfra.CommonRepository) *CommonService {
	return &CommonService{repo: repo}
}
