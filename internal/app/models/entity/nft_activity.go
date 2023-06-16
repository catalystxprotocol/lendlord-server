package entity

import "gorm.io/gorm"

const TableNftActivity = "t_nft_activities"

type NftActivities struct {
	gorm.Model
	NftOrderId     uint64 `gorm:"column:nft_order_id;type:bigint(20) unsigned;default:0;comment:nft 订单表 id;NOT NULL" json:"nft_order_id"`
	ChainConfigId  uint64 `gorm:"column:chain_config_id;type:bigint(20) unsigned;default:0;comment:链配置表 id;NOT NULL" json:"chain_config_id"`
	CollectionId   uint64 `gorm:"column:collection_id;type:bigint(20) unsigned;default:0;comment:collection 表 id;NOT NULL" json:"collection_id"`
	NftId          uint64 `gorm:"column:nft_id;type:bigint(20) unsigned;default:0;comment:nft 表id;NOT NULL" json:"nft_id"`
	Status         uint   `gorm:"column:status;type:tinyint(4) unsigned;default:0;comment:活动状态（1：List 2：Delist 3：Rent, 4：Redeem）;NOT NULL" json:"status"`
	LenderAddr     string `gorm:"column:lender_addr;type:varchar(255);comment:nft 出租人地址;NOT NULL" json:"lender_addr"`
	RenterAddr     string `gorm:"column:renter_addr;type:varchar(255);comment:nft 承租人地址;NOT NULL" json:"renter_addr"`
	OnftId         string `gorm:"column:onft_id;type:varchar(255);comment:onft token id;NOT NULL" json:"onft_id"`
	OnftAddr       string `gorm:"column:onft_addr;type:varchar(255);comment:onft 合约地址;NOT NULL" json:"onft_addr"`
	DonftId        string `gorm:"column:donft_id;type:varchar(255);comment:donft token id;NOT NULL" json:"donft_id"`
	DonftAddr      string `gorm:"column:donft_addr;type:varchar(255);comment:donft 合约地址;NOT NULL" json:"donft_addr"`
	Price          string `gorm:"column:price;type:varchar(255);comment:租赁每日价格/承租价格;NOT NULL" json:"price"`
	StartTimestamp uint   `gorm:"column:start_timestamp;type:int(11) unsigned;default:0;comment:租赁开始时间/承租开始时间;NOT NULL" json:"start_timestamp"`
	EndTimestamp   uint   `gorm:"column:end_timestamp;type:int(11) unsigned;default:0;comment:租赁结束时间/承租结束时间;NOT NULL" json:"end_timestamp"`
	Duration       uint   `gorm:"column:duration;type:int(11) unsigned;default:0;comment:租赁小持续时间/承租者租赁持续时间;NOT NULL" json:"duration"`
	TxHash         string `gorm:"column:tx_hash;type:varchar(100);comment:交易 hash" json:"tx_hash"`
}

var NftActivityModelFields nftActivityModelFields

// nftActivityModelFields
type nftActivityModelFields struct {
	Id             string
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
	NftOrderId     string
	ChainConfigId  string
	CollectionId   string
	NftId          string
	Status         string
	LenderAddr     string
	RenterAddr     string
	OnftId         string
	OnftAddr       string
	DonftId        string
	DonftAddr      string
	Price          string
	TxHash         string
	StartTimestamp string
	EndTimestamp   string
	Duration       string
}

func init() {
	// init model filed value
	NftActivityModelFields = InitModelFields(TableNftActivity, NftActivityModelFields).(nftActivityModelFields)
}

const (
	Listed int = iota + 1
	_
	Rented
)
