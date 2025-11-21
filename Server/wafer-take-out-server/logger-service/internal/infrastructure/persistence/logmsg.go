package persistence

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/logger-service/internal/domain"
	"gorm.io/gorm"
)

type DefaultLogMsgRepository struct {
	db *gorm.DB
}

func NewDefaultLogMsgRepository(db *gorm.DB) *DefaultLogMsgRepository {
	return &DefaultLogMsgRepository{db: db}
}

func (repo *DefaultLogMsgRepository) Save(logmsg *domain.LogMsg) error {

	err := repo.db.
		Model(&domain.LogMsg{}).
		Create(logmsg).Error
	if err != nil {
		return err
	}
	return nil
}
