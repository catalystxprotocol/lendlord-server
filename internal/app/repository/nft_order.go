package repository

import (
	"github.com/lendlord/lendlord-server/internal/app/models/entity"
	"gorm.io/gorm"
)

type INftOrderRepo interface {
	GetNftOrder(where map[string]interface{}, column ...string) (entity.NftOrders, error)
	GetNftOrders(where map[string]interface{}, column ...string) ([]entity.NftOrders, error)
}

type NftOrderRepo struct {
	db *gorm.DB
}

func NewNftOrdersRepo(db *gorm.DB) INftOrderRepo {
	return &NftOrderRepo{
		db: db,
	}
}

func (n NftOrderRepo) GetNftOrder(where map[string]interface{}, column ...string) (info entity.NftOrders, err error) {
	Get := n.db.Where(where)
	if len(column) > 0 {
		Get.Select(column)
	}
	err = Get.First(&info).Error

	return
}

func (n NftOrderRepo) GetNftOrders(where map[string]interface{}, column ...string) (list []entity.NftOrders, err error) {
	Get := n.db.Where(where)
	if len(column) > 0 {
		Get.Select(column)
	}
	err = Get.Find(&list).Error

	return
}
