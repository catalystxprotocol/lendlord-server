package dto

type CollectionByUserList struct {
	ID   string `gorm:"primarykey"` // t_nft_collections table id
	Name string `gorm:"column:name;type:varchar(100);comment:collection 名称;NOT NULL" json:"name"`
}

type NftListByUser struct {
	NftId                uint   `json:"nft_id"` // t_nfts table id
	TokenId              string `json:"token_id"`
	NftAddr              string `json:"nft_addr"`
	NftOwner             string `gorm:"column:owner_addr;type:varchar(255);comment:nft 拥有者地址;NOT NULL" json:"owner_addr"`
	ChainId              string `gorm:"column:chain_id;type:varchar(100);comment:链 chain id;NOT NULL" json:"chain_id"`
	ShelfStatus          uint   `gorm:"column:shelf_status;type:tinyint(4) unsigned;default:0;comment:上架状态（1：List 2：Delist 3：Redeem）;NOT NULL" json:"shelf_status"` //TODO  uint or string ?
	CollectionName       string `gorm:"column:collection_name;type:varchar(100);comment:collection 名称;NOT NULL" json:"collection_name"`
	LenderAddr           string `gorm:"column:lender_addr;type:varchar(255);comment:nft 出租人地址;NOT NULL" json:"lender_addr"`
	LenderPrice          string `gorm:"column:lender_price;type:varchar(255);comment:租赁 nft 价格;NOT NULL" json:"lender_price"`
	RentMinDuration      uint   `gorm:"column:renter_duration;type:int(11) unsigned;default:0;comment:承租者租赁最小持续时间;NOT NULL" json:"renter_duration"`
	LenderStartTimestamp int64  `gorm:"column:lender_start_timestamp;type:int(11) unsigned;default:0;comment:租赁开始时间戳;NOT NULL" json:"lender_start_timestamp"`
	LenderEndTimestamp   int64  `gorm:"column:lender_end_timestamp;type:int(11) unsigned;default:0;comment:租赁结束时间戳;NOT NULL" json:"lender_end_timestamp"`
	UseLink              string `gorm:"column:use_link;type:varchar(255);comment:use 链接;NOT NULL" json:"use_link"`
	NftImage             string `gorm:"column:image;type:varchar(255);comment:nft 图片地址;NOT NULL" json:"image"`
	NftName              string `gorm:"column:nft_name;type:varchar(100);comment:nft 名称;NOT NULL" json:"nft_name"`
	ONftId               string `gorm:"column:onft_id;type:varchar(255);comment:onft token id;NOT NULL" json:"onft_id"`
	//ChainId  string `json:"chain_id"`
	Symbol   string `gorm:"column:token_symbol;type:varchar(100);comment:币符号;NOT NULL" json:"token_symbol"`
	Decimals uint   `gorm:"column:decimals;type:tinyint(4) unsigned;default:0;comment:币位数;NOT NULL" json:"decimals"`
	Icon     string `gorm:"column:token_icon;type:varchar(100);comment:币图标;NOT NULL" json:"token_icon"`

	OrderId              uint   `json:"order_id"` // t_nft_orders table id
	RenterAddr           string `gorm:"column:renter_addr;type:varchar(255);comment:承租人地址;NOT NULL" json:"renter_addr"`
	RenterPrice          string `gorm:"column:renter_price;type:varchar(255);comment:承租价格;NOT NULL" json:"renter_price"`
	RenterStartTimestamp int64  `gorm:"column:renter_start_timestamp;type:int(11) unsigned;default:0;comment:承租者开始时间戳;NOT NULL" json:"renter_start_timestamp"`
	RenterEndTimestamp   int64  `gorm:"column:renter_end_timestamp;type:int(11) unsigned;default:0;comment:承租者结束时间戳;NOT NULL" json:"renter_end_timestamp"`
}

type NftDetail struct {
	ChainId              string `gorm:"column:chain_id;type:varchar(100);comment:链 chain id;NOT NULL" json:"chain_id"`
	NftAddr              string `json:"nft_addr"`
	NftId                uint   `json:"nft_id"` // t_nfts table id
	TokenId              string `json:"token_id"`
	NftOwner             string `gorm:"column:owner_addr;type:varchar(255);comment:nft 拥有者地址;NOT NULL" json:"owner_addr"`
	RenterAddr           string `json:"renter_addr"` //listed 且 成交赋值给 nft_user
	LenderAddr           string `gorm:"column:lender_addr;type:varchar(255);comment:nft 出租人地址;NOT NULL" json:"lender_addr"`
	ONftId               string `gorm:"column:onft_id;type:varchar(255);comment:onft token id;NOT NULL" json:"onft_id"`
	ONftAddr             string `gorm:"column:onft_addr;type:varchar(255);comment:onft token id;NOT NULL" json:"onft_addr"`
	DonftAddr            string `gorm:"column:donft_addr;type:varchar(255);comment:donft 合约地址;NOT NULL" json:"donft_addr"`
	Price                string `gorm:"column:lender_price;type:varchar(255);comment:租赁 nft 价格;NOT NULL" json:"lender_price"`
	LenderMinDuration    uint   `gorm:"column:lender_min_duration;type:int(11) unsigned;default:0;comment:租赁最小持续时间;NOT NULL" json:"lender_min_duration"`
	LenderStartTimestamp int64  `gorm:"column:lender_start_timestamp;type:int(11) unsigned;default:0;comment:租赁开始时间戳;NOT NULL" json:"lender_start_timestamp"`
	LenderEndTimestamp   int64  `gorm:"column:lender_end_timestamp;type:int(11) unsigned;default:0;comment:租赁结束时间戳;NOT NULL" json:"lender_end_timestamp"`
	ShelfStatus          uint   `gorm:"column:shelf_status;type:tinyint(4) unsigned;default:0;comment:上架状态（1：List 2：Delist 3：Redeem）;NOT NULL" json:"shelf_status"`
	NftImage             string `gorm:"column:image;type:varchar(255);comment:nft 图片地址;NOT NULL" json:"image"`
	NftName              string `gorm:"column:nft_name;type:varchar(100);comment:nft 名称;NOT NULL" json:"nft_name"`
	CollectionName       string `gorm:"column:collection_name;type:varchar(100);comment:collection 名称;NOT NULL" json:"collection_name"`
	Symbol               string `gorm:"column:token_symbol;type:varchar(100);comment:币符号;NOT NULL" json:"token_symbol"`
	Decimals             uint   `gorm:"column:decimals;type:tinyint(4) unsigned;default:0;comment:币位数;NOT NULL" json:"decimals"`
	Icon                 string `gorm:"column:token_icon;type:varchar(100);comment:币图标;NOT NULL" json:"token_icon"`
	RenterStartTimestamp int64  `gorm:"column:renter_start_timestamp;type:int(11) unsigned;default:0;comment:承租者开始时间戳;NOT NULL" json:"renter_start_timestamp"`
	RenterEndTimestamp   int64  `gorm:"column:renter_end_timestamp;type:int(11) unsigned;default:0;comment:承租者结束时间戳;NOT NULL" json:"renter_end_timestamp"`
	OrderId              uint   `json:"order_id"` // t_nft_orders table id
}

type NftActivityList struct {
	Event          uint   `gorm:"column:status;type:tinyint(4) unsigned;default:0;comment:活动状态（1：List 2：Delist 3：Rent, 4：Redeem）;NOT NULL" json:"status"`
	LenderAddr     string `json:"lender_addr"`
	RenterAddr     string `json:"renter_addr"`
	Price          string `json:"price"`
	Duration       uint   `json:"duration"`
	EndTimestamp   uint   `json:"end_timestamp"`
	StartTimestamp uint   `json:"start_timestamp"`
	ChainId        string `gorm:"column:chain_id;type:varchar(100);comment:链 chain id;NOT NULL" json:"chain_id"`
	Symbol         string `gorm:"column:token_symbol;type:varchar(100);comment:币符号;NOT NULL" json:"token_symbol"`
	Decimals       uint   `gorm:"column:decimals;type:tinyint(4) unsigned;default:0;comment:币位数;NOT NULL" json:"decimals"`
	Icon           string `gorm:"column:token_icon;type:varchar(100);comment:币图标;NOT NULL" json:"token_icon"`
}

type NftLendInfo struct {
	CollectionName    string `gorm:"column:collection_name;type:varchar(100);comment:collection 名称;NOT NULL" json:"collection_name"`
	TokenId           string `json:"token_id"`
	NftImage          string `gorm:"column:image;type:varchar(255);comment:nft 图片地址;NOT NULL" json:"image"`
	NftOwner          string `gorm:"column:owner_addr;type:varchar(255);comment:nft 拥有者地址;NOT NULL" json:"owner_addr"`
	Icon              string `gorm:"column:token_icon;type:varchar(100);comment:币图标;NOT NULL" json:"token_icon"`
	Symbol            string `gorm:"column:token_symbol;type:varchar(100);comment:币符号;NOT NULL" json:"token_symbol"`
	Decimals          uint   `gorm:"column:decimals;type:tinyint(4) unsigned;default:0;comment:币位数;NOT NULL" json:"decimals"`
	TokenName         string `gorm:"column:token_name;type:varchar(100);comment:币名称;NOT NULL" json:"token_name"`
	NftAddr           string `gorm:"column:nft_addr;type:varchar(255);comment:nft 合约地址;NOT NULL" json:"nft_addr"`
	ChainId           string `gorm:"column:chain_id;type:varchar(100);comment:链 chain id;NOT NULL" json:"chain_id"`
	DonftContractAddr string `gorm:"column:donft_contract_addr;type:varchar(255);comment:赎回合约地址;NOT NULL" json:"donft_contract_addr"`
	ShelfStatus       uint   `gorm:"column:shelf_status;type:tinyint(4) unsigned;default:0;comment:上架状态（1：List 2：Delist 3：Redeem）;NOT NULL" json:"shelf_status"`
	LenderAddr        string `gorm:"column:lender_addr;type:varchar(255);comment:nft 出租人地址;NOT NULL" json:"lender_addr"`
	DonftId           string `gorm:"column:donft_id;type:varchar(255);comment:donft token id;NOT NULL" json:"donft_id"`
}
