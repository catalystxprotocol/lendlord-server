package repository

import (
	"github.com/lendlord/lendlord-server/internal/app/models/entity"
	"gorm.io/gorm"
)

type IChainConfigRepo interface {
	GetChainConfigs() ([]entity.ChainConfigs, error)
}

type ChainConfigRepo struct {
	db *gorm.DB
}

func NewChainConfigRepo(db *gorm.DB) IChainConfigRepo {
	return &ChainConfigRepo{
		db: db,
	}
}

func (r ChainConfigRepo) GetChainConfigs() (list []entity.ChainConfigs, err error) {
	err = r.db.Find(&list).Error
	return
}
