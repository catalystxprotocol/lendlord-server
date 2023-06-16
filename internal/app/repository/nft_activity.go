package repository

import (
	"fmt"
	"math"

	"github.com/lendlord/lendlord-server/internal/app/models/dto"
	"github.com/lendlord/lendlord-server/internal/app/models/entity"
	"github.com/lendlord/lendlord-server/internal/app/models/form"
	"gorm.io/gorm"
)

var _ INftActivityRepo = (*NftActivityRepo)(nil)

type INftActivityRepo interface {
	GetNftActivity(where map[string]interface{}, column ...string) (entity.NftActivities, error)
	GetNftActivities(where map[string]interface{}, column ...string) ([]entity.NftActivities, error)
	GetNftActivityList(params *form.PageQueryNftActivityListForm) (total int64, totalPage int64, nftList []dto.NftActivityList, err error)
}

type NftActivityRepo struct {
	db *gorm.DB
}

func NewNftActivityRepo(db *gorm.DB) INftActivityRepo {
	return &NftActivityRepo{
		db: db,
	}
}

func (n NftActivityRepo) GetNftActivity(where map[string]interface{}, column ...string) (info entity.NftActivities, err error) {
	Get := n.db.Where(where)
	if len(column) > 0 {
		Get.Select(column)
	}
	err = Get.First(&info).Error

	return
}

func (n NftActivityRepo) GetNftActivities(where map[string]interface{}, column ...string) (list []entity.NftActivities, err error) {
	Get := n.db.Where(where)
	if len(column) > 0 {
		Get.Select(column)
	}
	err = Get.Find(&list).Error

	return
}

func (n NftActivityRepo) GetNftActivityList(params *form.PageQueryNftActivityListForm) (total int64, totalPage int64, list []dto.NftActivityList, err error) {
	query := n.db.Table(entity.TableNftActivity).Select([]string{
		entity.NftActivityModelFields.Status,
		entity.NftActivityModelFields.LenderAddr,
		entity.NftActivityModelFields.RenterAddr,
		entity.NftActivityModelFields.Price,
		entity.NftActivityModelFields.Duration,
		entity.NftActivityModelFields.EndTimestamp,
		entity.NftActivityModelFields.StartTimestamp,

		entity.ChainConfigModelFields.ChainId,
		entity.ChainConfigModelFields.TokenSymbol,
		entity.ChainConfigModelFields.TokenIcon,
		entity.ChainConfigModelFields.Decimals,
	}).
		Joins(fmt.Sprintf("left join %s on %s = %s", entity.TableChainConfig, entity.NftActivityModelFields.ChainConfigId, entity.ChainConfigModelFields.Id)).
		Where(map[string]interface{}{entity.NftActivityModelFields.NftId: params.NftId})

	if params.EventType != 0 {
		switch params.EventType {
		case entity.Listed:
			query = query.Where(map[string]interface{}{entity.NftActivityModelFields.Status: entity.Listed})
		case entity.Rented:
			query = query.Where(map[string]interface{}{entity.NftActivityModelFields.Status: entity.Rented})
		}
	} else {
		query = query.Where(fmt.Sprintf("%s IN ", entity.NftActivityModelFields.Status)+"?", []int{entity.Listed, entity.Rented})
	}

	query.Count(&total)
	totalPage = int64(math.Ceil(float64(total) / float64(params.PageSize)))
	query.Offset((params.PageNumber - 1) * params.PageSize).Limit(params.PageSize).Order(entity.NftActivityModelFields.StartTimestamp + " DESC")

	err = query.Scan(&list).Error
	return
}
