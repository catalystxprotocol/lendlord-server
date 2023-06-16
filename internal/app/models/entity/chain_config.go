package entity

import "gorm.io/gorm"

const TableChainConfig = "t_chain_configs"

type ChainConfigs struct {
	gorm.Model
	ChainId            string `gorm:"column:chain_id;type:varchar(100);comment:链 chain id;NOT NULL" json:"chain_id"`
	ChainName          string `gorm:"column:chain_name;type:varchar(100);comment:链名称;NOT NULL" json:"chain_name"`
	TokenName          string `gorm:"column:token_name;type:varchar(100);comment:币名称;NOT NULL" json:"token_name"`
	TokenSymbol        string `gorm:"column:token_symbol;type:varchar(100);comment:币符号;NOT NULL" json:"token_symbol"`
	Decimals           uint   `gorm:"column:decimals;type:tinyint(4) unsigned;default:0;comment:币位数;NOT NULL" json:"decimals"`
	TokenIcon          string `gorm:"column:token_icon;type:varchar(100);comment:币图标;NOT NULL" json:"token_icon"`
	RpcUrl             string `gorm:"column:rpc_url;type:varchar(255);comment:RPC地址;NOT NULL" json:"rpc_url"`
	BlockExplorerUrl   string `gorm:"column:block_explorer_url;type:varchar(255);comment:区块浏览器;NOT NULL" json:"block_explorer_url"`
	TxHashRoute        string `gorm:"column:tx_hash_route;type:varchar(255);comment:区块浏览器路由地址拼接;NOT NULL" json:"tx_hash_route"`
	ChainIconUrl       string `gorm:"column:chain_icon_url;type:varchar(255);comment:链图标地址;NOT NULL" json:"chain_icon_url"`
	MarketContractAddr string `gorm:"column:market_contract_addr;type:varchar(255);comment:租赁合约地址;NOT NULL" json:"market_contract_addr"`
	DonftContractAddr  string `gorm:"column:donft_contract_addr;type:varchar(255);comment:赎回合约地址;NOT NULL" json:"donft_contract_addr"`
}

var ChainConfigModelFields chainConfigModelFields

// chainConfigModelFields
type chainConfigModelFields struct {
	Id                   string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	ChainId              string
	ChainName            string
	TokenName            string
	TokenSymbol          string
	Decimals             string
	TokenIcon            string
	RpcUrl               string
	BlockExplorerUrl     string
	TxHashRoute          string
	ChangeCertificateUrl string
	ChainIconUrl         string
	MarketContractAddr   string
	DonftContractAddr    string
}

func init() {
	// init model filed value
	ChainConfigModelFields = InitModelFields(TableChainConfig, ChainConfigModelFields).(chainConfigModelFields)
}
