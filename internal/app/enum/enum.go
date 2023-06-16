package enum

const (
	DefaultPage = 1
	DefaultSize = 10
)

// nft列表状态 status
const (
	Unknown = iota
	Idle
	Listed
	ListDeal // 挂单成交
	Rented
	IdleOwner     // NFT状态为Idle且Owner是自己
	OwnerLendRent // 自己的NFT自己租
)

// nft_order 表 shelf_status状态
const (
	Zero   = iota
	List   // 上架
	Delist // 下架
	Redeem // 赎回
)
