package repository

import (
	"fmt"
	"github.com/lendlord/lendlord-server/internal/app/models/dto"
	"github.com/lendlord/lendlord-server/internal/app/models/entity"
	"gorm.io/gorm"
)

type INftCollectionRepo interface {
	GetNftCollectionsByUser(userAddr string) (list []dto.CollectionByUserList, err error)
	GetNftCollectionSum() (sum int64, err error)
}

type NftCollectionRepo struct {
	db *gorm.DB
}

func NewNftCollectionRepo(db *gorm.DB) INftCollectionRepo {
	return &NftCollectionRepo{
		db: db,
	}
}

func (n NftCollectionRepo) GetNftCollectionsByUser(userAddr string) (list []dto.CollectionByUserList, err error) {
	err = n.db.Table(entity.TableNftCollection).Select([]string{
		entity.NftCollectionFields.Id,
		entity.NftCollectionFields.Name,
	}).
		Joins(fmt.Sprintf("inner join %s on %s = %s", entity.TableNft, entity.NftCollectionFields.Id, entity.NftFields.CollectionId)).
		Joins(fmt.Sprintf("left join %s on %s = %s", entity.TableNftOrder, entity.NftOrderFields.CollectionId, entity.NftFields.CollectionId)).
		Where(map[string]interface{}{entity.NftFields.OwnerAddr: userAddr}).
		Or(map[string]interface{}{entity.NftOrderFields.LenderAddr: userAddr}).
		Or(map[string]interface{}{entity.NftOrderFields.RenterAddr: userAddr}).
		Group(entity.NftCollectionFields.Id).Scan(&list).Error

	return
}

func (n NftCollectionRepo) GetNftCollectionSum() (sum int64, err error) {
	err = n.db.Table(entity.TableNft).Distinct(entity.NftFields.CollectionId).Count(&sum).Error
	return
}
