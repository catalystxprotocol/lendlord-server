package entity

import "gorm.io/gorm"

const TableNft = "t_nfts"

type Nfts struct {
	gorm.Model
	ChainConfigId uint64 `gorm:"column:chain_config_id;type:bigint(20) unsigned;default:0;comment:链配置表 id;NOT NULL" json:"chain_config_id"`
	CollectionId  uint64 `gorm:"column:collection_id;type:bigint(20) unsigned;default:0;comment:collection 表 id;NOT NULL" json:"collection_id"`
	TokenId       string `gorm:"column:token_id;type:varchar(255);comment:nft token id;NOT NULL" json:"token_id"`
	NftAddr       string `gorm:"column:nft_addr;type:varchar(255);comment:nft 合约地址;NOT NULL" json:"nft_addr"`
	Name          string `gorm:"column:name;type:varchar(100);comment:nft 名称;NOT NULL" json:"name"`
	Image         string `gorm:"column:image;type:varchar(255);comment:nft 图片地址;NOT NULL" json:"image"`
	Description   string `gorm:"column:describe;type:text;comment:nft 介绍" json:"describe"`
	OwnerAddr     string `gorm:"column:owner_addr;type:varchar(255);comment:nft 拥有者地址;NOT NULL" json:"owner_addr"`
	UseLink       string `gorm:"column:use_link;type:varchar(255);comment:use 链接;NOT NULL" json:"use_link"`
}

var NftFields nftFields

// nftFields
type nftFields struct {
	Id            string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	ChainConfigId string
	CollectionId  string
	TokenId       string
	NftAddr       string
	Name          string
	Image         string
	Description   string
	OwnerAddr     string
	UseLink       string
}

func init() {
	// init model filed value
	NftFields = InitModelFields(TableNft, NftFields).(nftFields)
}
