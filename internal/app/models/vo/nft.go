package vo

type (
	QueryCollectionByUserVo struct {
		CollectionByUserList []CollectionByUserList `json:"list"`
	}

	CollectionByUserList struct {
		Id   string `json:"id"` // t_nft_collections table id
		Name string `json:"name"`
	}

	PageQueryNftByUserVo struct {
		PageInfo               PageInfo                 `json:"page_info"`
		PageQueryNftByUserList []PageQueryNftByUserList `json:"list"`
	}

	PageQueryNftByUserList struct {
		NftId            uint             `json:"nft_id"` // t_nfts table id
		NftTokenId       string           `json:"nft_token_id"`
		NftAddr          string           `json:"nft_addr"`
		NftOwner         string           `json:"nft_owner"`
		ChainId          string           `json:"chain_id"`
		NftStatus        uint             `json:"nft_status"`
		CollectionName   string           `json:"collection_name"`
		EndTimestamp     int64            `json:"end_timestamp"`
		LenderPrice      string           `json:"lender_price"`
		RentMinDuration  uint             `json:"rent_min_duration"`
		LendEndTimestamp int64            `json:"lend_end_timestamp"`
		UseLink          string           `json:"use_link"`
		NftImage         string           `json:"nft_image"`
		NftName          string           `json:"nft_name"`
		ONftId           string           `json:"onft_id"`
		PaymentTokenInfo PaymentTokenInfo `json:"payment_token_info"`
	}

	QueryNftDetailVo struct {
		ChainId          string           `json:"chain_id"`
		NftAddr          string           `json:"nft_addr"`
		NftId            uint             `json:"nft_id"`
		NftTokenId       string           `json:"nft_token_id"`
		NftOwner         string           `json:"nft_owner"`
		LenderAddr       string           `json:"lender_addr"`
		NftUser          string           `json:"nft_user"`
		ONftId           string           `json:"onft_id"`
		ONftAddr         string           `json:"onft_addr"`
		EndTimestamp     int64            `json:"end_timestamp"`
		Price            string           `json:"price"`
		RentMinDuration  uint             `json:"rent_min_duration"`
		LendEndTimestamp int64            `json:"lend_end_timestamp"`
		NftStatus        uint             `json:"nft_status"`
		NftImage         string           `json:"nft_image"`
		NftName          string           `json:"nft_name"`
		CollectionName   string           `json:"collection_name"`
		PaymentTokenInfo PaymentTokenInfo `json:"payment_token_info"`
	}

	PageQueryNftActivityVo struct {
		PageInfo                 PageInfo                   `json:"page_info"`
		PageQueryNftActivityList []PageQueryNftActivityList `json:"list"`
	}

	PageQueryNftActivityList struct {
		Event            uint             `json:"event"`
		Lender           string           `json:"lender"`
		Renter           string           `json:"renter"`
		Price            string           `json:"price"`
		RentMinDuration  uint             `json:"rent_min_duration"`
		LendEndTimestamp uint             `json:"lend_end_timestamp"`
		StartTimestamp   uint             `json:"start_timestamp"`
		PaymentTokenInfo PaymentTokenInfo `json:"payment_token_info"`
	}

	QueryNftLendInfoVo struct {
		CollectionName   string         `json:"collection_name"`
		NftTokenId       string         `json:"nft_token_id"`
		NftImage         string         `json:"nft_image"`
		NftOwner         string         `json:"nft_owner"`
		ONftId           string         `json:"onft_id"`
		ONftAddr         string         `json:"onft_addr"`
		ChainID          string         `json:"chain_id"`
		DonftId          string         `json:"donft_id"`
		DoNftAddr        string         `json:"donft_addr"`
		LenderAddr       string         `json:"lender_addr"`
		PaymentTokenList []PaymentToken `json:"payment_token_list"`
	}
	QueryNftCollectionSumVo struct {
		TotalSum int64 `json:"total_sum"`
	}
)
