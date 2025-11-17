package infrastructure

import (
	"context"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DefaultUserRepository struct {
	db *gorm.DB
}

func NewDefaultUserRepository(db *gorm.DB) domain.UserRepository {
	return &DefaultUserRepository{db: db}
}

func (repo *DefaultUserRepository) Upsert(ctx context.Context,
	user *domain.User) error {

	err := repo.db.WithContext(ctx).
		Model(&domain.User{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "openid"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "phone", "sex", "id_number", "avatar"}),
		}).Create(user).Error

	err = repo.db.WithContext(ctx).
		Where("openid = ?", user.OpenId).
		First(user).Error

	return err

}
