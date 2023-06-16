package repository

import (
	"fmt"
	"math"
	"time"

	"github.com/lendlord/lendlord-server/internal/app/enum"
	"github.com/lendlord/lendlord-server/internal/app/models/dto"
	"github.com/lendlord/lendlord-server/internal/app/models/entity"
	"github.com/lendlord/lendlord-server/internal/app/models/form"
	"gorm.io/gorm"
)

type INftRepo interface {
	GetNft(where map[string]interface{}, column ...string) (entity.Nfts, error)
	GetNfts(where map[string]interface{}, column ...string) ([]entity.Nfts, error)
	GetNftListByUser(params *form.PageQueryNftListByUserForm) (total int64, totalPage int64, nftList []dto.NftListByUser, err error)
	QueryNftDetail(nftId int) (detail *dto.NftDetail, err error)
	GetNftLendInfo(nftId int) (lendInfo *dto.NftLendInfo, err error)
}

type NftRepo struct {
	db *gorm.DB
}

func NewNftRepo(db *gorm.DB) INftRepo {
	return &NftRepo{
		db: db,
	}
}

func (n NftRepo) GetNft(where map[string]interface{}, column ...string) (info entity.Nfts, err error) {
	Get := n.db.Where(where)
	if len(column) > 0 {
		Get.Select(column)
	}
	err = Get.First(&info).Error

	return
}

func (n NftRepo) GetNfts(where map[string]interface{}, column ...string) (list []entity.Nfts, err error) {
	Get := n.db.Where(where)
	if len(column) > 0 {
		Get.Select(column)
	}
	err = Get.Find(&list).Error

	return
}

func (n NftRepo) GetNftListByUser(params *form.PageQueryNftListByUserForm) (total int64, totalPage int64, nftList []dto.NftListByUser, err error) {
	query := n.db.Table(entity.TableNft).Select([]string{
		entity.NftFields.Id + " AS nft_id",
		entity.NftFields.TokenId,
		entity.NftFields.NftAddr,
		entity.NftFields.OwnerAddr,
		entity.NftFields.UseLink,
		entity.NftFields.Name + " AS nft_name",
		entity.NftFields.Image,

		entity.ChainConfigModelFields.ChainId,
		entity.ChainConfigModelFields.TokenSymbol,
		entity.ChainConfigModelFields.TokenIcon,
		entity.ChainConfigModelFields.Decimals,
		entity.NftCollectionFields.Name + " AS collection_name",

		entity.NftOrderFields.ShelfStatus,
		entity.NftOrderFields.LenderPrice,
		entity.NftOrderFields.LenderEndTimestamp,
		entity.NftOrderFields.OnftId,
		entity.NftOrderFields.Id + " AS order_id",
		entity.NftOrderFields.RenterAddr,
		entity.NftOrderFields.RenterPrice,
		entity.NftOrderFields.RenterDuration,
		entity.NftOrderFields.RenterStartTimestamp,
		entity.NftOrderFields.RenterEndTimestamp,
		entity.NftOrderFields.LenderAddr,
	}).
		Joins(fmt.Sprintf("inner join %s on %s = %s", entity.TableChainConfig, entity.NftFields.ChainConfigId, entity.ChainConfigModelFields.Id)).
		Joins(fmt.Sprintf("inner join %s on %s = %s", entity.TableNftCollection, entity.NftFields.CollectionId, entity.NftCollectionFields.Id)).
		Joins(fmt.Sprintf("left join %s on %s = %s", entity.TableNftOrder, entity.NftFields.Id, entity.NftOrderFields.NftId))

	// "t_nfts.owner_addr = ? AND (t_nft_orders.shelf_status IN ? OR t_nft_orders.id = null)"
	DeListRedeemStatement := fmt.Sprintf("%s = ? AND %s IN ?", entity.NftOrderFields.LenderAddr, entity.NftOrderFields.ShelfStatus)

	NoListStatement := fmt.Sprintf("%s = ? AND ISNULL(%v)", entity.NftFields.OwnerAddr, entity.NftOrderFields.Id)

	// "t_nft_orders.renter_addr = ? AND t_nft_orders.renter_start_timestamp <= ? AND t_nft_orders.renter_end_timestamp >= ?"
	rentedStatement := fmt.Sprintf("%s = ? AND %v <= ? AND %v >= ?", entity.NftOrderFields.RenterAddr, entity.NftOrderFields.RenterStartTimestamp, entity.NftOrderFields.RenterEndTimestamp)

	rentedStatementByCollection := fmt.Sprintf("%s = ? AND %v <= ? AND %v >= ? AND %s = ?", entity.NftOrderFields.RenterAddr,
		entity.NftOrderFields.RenterStartTimestamp, entity.NftOrderFields.RenterEndTimestamp, entity.NftFields.CollectionId)
	// 上架过期
	listExceed := fmt.Sprintf("%s = ? AND %s = ? AND NOT (%s <= ? AND %s >= ?)", entity.NftOrderFields.LenderAddr,
		entity.NftOrderFields.ShelfStatus, entity.NftOrderFields.LenderStartTimestamp, entity.NftOrderFields.LenderEndTimestamp)

	// 上架
	list := fmt.Sprintf("%s = ? AND %s = ? AND %s <= ? AND %s >= ?", entity.NftOrderFields.LenderAddr, entity.NftOrderFields.ShelfStatus,
		entity.NftOrderFields.LenderStartTimestamp, entity.NftOrderFields.LenderEndTimestamp)

	// "t_nfts.owner_addr = ? AND (t_nft_orders.shelf_status IN ? OR t_nft_orders.id = null)"
	DeListRedeemStatementByC := fmt.Sprintf("%s = ? AND %s IN ? AND %s = ?", entity.NftOrderFields.LenderAddr, entity.NftOrderFields.ShelfStatus, entity.NftFields.CollectionId)

	NoListStatementByC := fmt.Sprintf("%s = ? AND ISNULL(%v) AND %s = ?", entity.NftFields.OwnerAddr, entity.NftOrderFields.Id, entity.NftFields.CollectionId)

	// "t_nft_orders.renter_addr = ? AND t_nft_orders.renter_start_timestamp <= ? AND t_nft_orders.renter_end_timestamp >= ?"
	rentedStatementByC := fmt.Sprintf("%s = ? AND %v <= ? AND %v >= ? AND %s = ?", entity.NftOrderFields.RenterAddr,
		entity.NftOrderFields.RenterStartTimestamp, entity.NftOrderFields.RenterEndTimestamp, entity.NftFields.CollectionId)

	// 上架过期
	listExceedByC := fmt.Sprintf("%s = ? AND %s = ? AND NOT (%s <= ? AND %s >= ?) AND %s = ?", entity.NftOrderFields.LenderAddr,
		entity.NftOrderFields.ShelfStatus, entity.NftOrderFields.LenderStartTimestamp, entity.NftOrderFields.LenderEndTimestamp, entity.NftFields.CollectionId)

	// 上架
	listByC := fmt.Sprintf("%s = ? AND %s = ? AND %s <= ? AND %s >= ? AND %s = ?", entity.NftOrderFields.LenderAddr, entity.NftOrderFields.ShelfStatus,
		entity.NftOrderFields.LenderStartTimestamp, entity.NftOrderFields.LenderEndTimestamp, entity.NftFields.CollectionId)

	now := time.Now().Unix()

	if params.Status != 0 && params.CollectionId != 0 {
		switch params.Status {
		case enum.Idle: // 获取当前账户未上架或已租赁NFT数据
			query = query.Where(DeListRedeemStatementByC, params.UserAddr, []uint{entity.DeList, entity.Redeem}, params.CollectionId).
				Or(NoListStatementByC, params.UserAddr, params.CollectionId).
				Or(rentedStatementByC, params.UserAddr, now, now, params.CollectionId).
				Or(listExceedByC, params.UserAddr, entity.List, now, now, params.CollectionId)

		case enum.Listed: //当前账户已上架（判断是否过期）
			query = query.Where(listByC, params.UserAddr, entity.List, now, now, params.CollectionId)
		case enum.Rented: // 已租赁NFT数据
			query = query.Where(rentedStatementByC, params.UserAddr, now, now, params.CollectionId)
		}
		// 按状态查
	} else if params.Status != 0 {
		switch params.Status {
		case enum.Idle: // 获取当前账户未上架或已租赁NFT数据
			query = query.Where(DeListRedeemStatement, params.UserAddr, []uint{entity.DeList, entity.Redeem}).
				Or(NoListStatement, params.UserAddr).
				Or(rentedStatement, params.UserAddr, now, now).
				Or(listExceed, params.UserAddr, entity.List, now, now)

		case enum.Listed: //当前账户已上架（判断是否过期）
			query = query.Where(list, params.UserAddr, entity.List, now, now)
		case enum.Rented: // 已租赁NFT数据
			query = query.Where(rentedStatement, params.UserAddr, now, now)
		}

	} else if params.CollectionId != 0 {
		query = query.Where(map[string]interface{}{
			entity.NftFields.OwnerAddr:    params.UserAddr,
			entity.NftFields.CollectionId: params.CollectionId,
		}).
			Or(map[string]interface{}{
				entity.NftOrderFields.LenderAddr: params.UserAddr,
				entity.NftFields.CollectionId:    params.CollectionId,
			}).
			Or(rentedStatementByCollection, params.UserAddr, now, now, params.CollectionId)
	} else {
		// 查全部
		query = query.Where(map[string]interface{}{entity.NftFields.OwnerAddr: params.UserAddr}).
			Or(map[string]interface{}{entity.NftOrderFields.LenderAddr: params.UserAddr}).
			Or(rentedStatement, params.UserAddr, now, now)
	}

	query.Count(&total)
	totalPage = int64(math.Ceil(float64(total) / float64(params.PageSize)))
	query.Offset((params.PageNumber - 1) * params.PageSize).Limit(params.PageSize)

	err = query.Scan(&nftList).Error
	return
}

func (n NftRepo) QueryNftDetail(nftId int) (detail *dto.NftDetail, err error) {
	err = n.db.Table(entity.TableNft).Select([]string{
		entity.NftFields.Id + " AS nft_id",
		entity.NftFields.TokenId,
		entity.NftFields.NftAddr,
		entity.NftFields.OwnerAddr,
		entity.NftFields.Name + " AS nft_name",
		entity.NftFields.Image,

		entity.ChainConfigModelFields.ChainId,
		entity.ChainConfigModelFields.TokenSymbol,
		entity.ChainConfigModelFields.TokenIcon,
		entity.ChainConfigModelFields.Decimals,
		entity.NftCollectionFields.Name + " AS collection_name",

		entity.NftOrderFields.ShelfStatus,
		entity.NftOrderFields.RenterEndTimestamp,
		entity.NftOrderFields.RenterStartTimestamp,
		entity.NftOrderFields.RenterAddr,
		entity.NftOrderFields.LenderAddr,
		entity.NftOrderFields.LenderPrice,
		entity.NftOrderFields.LenderMinDuration,
		entity.NftOrderFields.LenderStartTimestamp,
		entity.NftOrderFields.LenderEndTimestamp,
		entity.NftOrderFields.OnftId,
		entity.NftOrderFields.OnftAddr,
		entity.NftOrderFields.DonftAddr,
		entity.NftOrderFields.Id + " AS order_id",
	}).
		Joins(fmt.Sprintf("inner join %s on %s = %s", entity.TableChainConfig, entity.NftFields.ChainConfigId, entity.ChainConfigModelFields.Id)).
		Joins(fmt.Sprintf("inner join %s on %s = %s", entity.TableNftCollection, entity.NftFields.CollectionId, entity.NftCollectionFields.Id)).
		Joins(fmt.Sprintf("left join %s on %s = %s", entity.TableNftOrder, entity.NftFields.Id, entity.NftOrderFields.NftId)).
		Where(map[string]interface{}{entity.NftFields.Id: nftId}).Take(&detail).Error

	return
}

func (n NftRepo) GetNftLendInfo(nftId int) (lendInfo *dto.NftLendInfo, err error) {
	err = n.db.Table(entity.TableNft).Select([]string{
		entity.NftFields.TokenId,
		entity.NftFields.OwnerAddr,
		entity.NftFields.Image,
		entity.NftCollectionFields.Name + " AS collection_name",
		entity.NftFields.NftAddr,

		entity.ChainConfigModelFields.TokenName,
		entity.ChainConfigModelFields.TokenSymbol,
		entity.ChainConfigModelFields.TokenIcon,
		entity.ChainConfigModelFields.Decimals,
		entity.ChainConfigModelFields.ChainId,
		entity.ChainConfigModelFields.DonftContractAddr,
		entity.NftOrderFields.ShelfStatus,
		entity.NftOrderFields.LenderAddr,
		entity.NftOrderFields.DonftId,
	}).
		Joins(fmt.Sprintf("inner join %s on %s = %s", entity.TableChainConfig, entity.NftFields.ChainConfigId, entity.ChainConfigModelFields.Id)).
		Joins(fmt.Sprintf("inner join %s on %s = %s", entity.TableNftCollection, entity.NftFields.CollectionId, entity.NftCollectionFields.Id)).
		Joins(fmt.Sprintf("left join %s on %s = %s", entity.TableNftOrder, entity.NftFields.Id, entity.NftOrderFields.NftId)).
		Where(map[string]interface{}{entity.NftFields.Id: nftId}).Take(&lendInfo).Error

	return
}
