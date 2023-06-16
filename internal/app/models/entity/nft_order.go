package entity

import "gorm.io/gorm"

const TableNftOrder = "t_nft_orders"

type NftOrders struct {
	gorm.Model
	ChainConfigId        uint64 `gorm:"column:chain_config_id;type:bigint(20) unsigned;default:0;comment:链配置表 id;NOT NULL" json:"chain_config_id"`
	CollectionId         uint64 `gorm:"column:collection_id;type:bigint(20) unsigned;default:0;comment:collection 表 id;NOT NULL" json:"collection_id"`
	NftId                uint64 `gorm:"column:nft_id;type:bigint(20) unsigned;default:0;comment:nft 表 id;NOT NULL" json:"nft_id"`
	ShelfStatus          uint   `gorm:"column:shelf_status;type:tinyint(4) unsigned;default:0;comment:上架状态（1：List 2：Delist 3：Redeem）;NOT NULL" json:"shelf_status"`
	OnftId               string `gorm:"column:onft_id;type:varchar(255);comment:onft token id;NOT NULL" json:"onft_id"`
	OnftAddr             string `gorm:"column:onft_addr;type:varchar(255);comment:onft 合约地址;NOT NULL" json:"onft_addr"`
	DonftId              string `gorm:"column:donft_id;type:varchar(255);comment:donft token id;NOT NULL" json:"donft_id"`
	DonftAddr            string `gorm:"column:donft_addr;type:varchar(255);comment:donft 合约地址;NOT NULL" json:"donft_addr"`
	LenderAddr           string `gorm:"column:lender_addr;type:varchar(255);comment:nft 出租人地址;NOT NULL" json:"lender_addr"`
	LenderPrice          string `gorm:"column:lender_price;type:varchar(255);comment:租赁 nft 价格;NOT NULL" json:"lender_price"`
	LenderMinDuration    uint   `gorm:"column:lender_min_duration;type:int(11) unsigned;default:0;comment:租赁最小持续时间;NOT NULL" json:"lender_min_duration"`
	LenderStartTimestamp uint   `gorm:"column:lender_start_timestamp;type:int(11) unsigned;default:0;comment:租赁开始时间戳;NOT NULL" json:"lender_start_timestamp"`
	LenderEndTimestamp   uint   `gorm:"column:lender_end_timestamp;type:int(11) unsigned;default:0;comment:租赁结束时间戳;NOT NULL" json:"lender_end_timestamp"`
	RenterAddr           string `gorm:"column:renter_addr;type:varchar(255);comment:承租人地址;NOT NULL" json:"renter_addr"`
	RenterPrice          string `gorm:"column:renter_price;type:varchar(255);comment:承租价格;NOT NULL" json:"renter_price"`
	RenterDuration       uint   `gorm:"column:renter_duration;type:int(11) unsigned;default:0;comment:承租者租赁持续时间;NOT NULL" json:"renter_duration"`
	RenterStartTimestamp uint   `gorm:"column:renter_start_timestamp;type:int(11) unsigned;default:0;comment:承租者开始时间戳;NOT NULL" json:"renter_start_timestamp"`
	RenterEndTimestamp   uint   `gorm:"column:renter_end_timestamp;type:int(11) unsigned;default:0;comment:承租者结束时间戳;NOT NULL" json:"renter_end_timestamp"`
}

var NftOrderFields nftOrderFields

// nftCollectionFields
type nftOrderFields struct {
	Id                   string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	ChainConfigId        string
	CollectionId         string
	NftId                string
	ShelfStatus          string
	OnftId               string
	OnftAddr             string
	DonftId              string
	DonftAddr            string
	LenderAddr           string
	LenderPrice          string
	LenderMinDuration    string
	LenderStartTimestamp string
	LenderEndTimestamp   string
	RenterAddr           string
	RenterPrice          string
	RenterDuration       string
	RenterStartTimestamp string
	RenterEndTimestamp   string
}

func init() {
	// init model filed value
	NftOrderFields = InitModelFields(TableNftOrder, NftOrderFields).(nftOrderFields)
}

const (
	List uint = iota + 1
	DeList
	Redeem
)
