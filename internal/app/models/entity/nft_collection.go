package entity

import "gorm.io/gorm"

const TableNftCollection = "t_nft_collections"

type NftCollections struct {
	gorm.Model
	ChainConfigId uint64 `gorm:"column:chain_config_id;type:bigint(20) unsigned;default:0;comment:链配置表 id;NOT NULL" json:"chain_config_id"`
	Name          string `gorm:"column:name;type:varchar(100);comment:collection 名称;NOT NULL" json:"name"`
	Symbol        string `gorm:"column:symbol;type:varchar(100);comment:collection 符号;NOT NULL" json:"symbol"`
	ImageLogo     string `gorm:"column:image_logo;type:varchar(255);comment:collection 图片 logo;NOT NULL" json:"image_logo"`
	ContractAddr  string `gorm:"column:contract_addr;type:varchar(255);comment:合约地址;NOT NULL" json:"contract_addr"`
	//OwnerAddr     string `gorm:"column:owner_addr;type:varchar(255);comment:拥有者地址;NOT NULL" json:"owner_addr"`
	Verified uint `gorm:"column:verified;type:tinyint(4) unsigned;default:0;comment:是否验证（0：未验证 1：已验证）;NOT NULL" json:"verified"`
}

var NftCollectionFields nftCollectionFields

// nftCollectionFields
type nftCollectionFields struct {
	Id            string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	ChainConfigId string
	Name          string
	Symbol        string
	ImageLogo     string
	ContractAddr  string
	//OwnerAddr     string
	RenterAddr string
	Verified   string
}

func init() {
	// init model filed value
	NftCollectionFields = InitModelFields(TableNftCollection, NftCollectionFields).(nftCollectionFields)
}
