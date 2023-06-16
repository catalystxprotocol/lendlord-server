package form

type (
	PageQueryNftListByUserForm struct {
		PageSize     int    `json:"page_size" form:"page_size"`
		PageNumber   int    `json:"page_number" form:"page_number"`
		Status       int    `json:"status" form:"status" binding:"omitempty,oneof= 1 2 4"` // 1-idle; 2-list; 4-rent
		CollectionId int64  `json:"collection_id" form:"collection_id" binding:"omitempty,gte=1"`
		UserAddr     string `form:"user_addr" binding:"required"`
	}

	PageQueryNftActivityListForm struct {
		NftId      int `json:"nft_id" form:"nft_id" binding:"required,gte=1"`
		PageSize   int `json:"page_size" form:"page_size"`
		PageNumber int `json:"page_number" form:"page_number"`
		EventType  int `json:"event_type" form:"event_type" binding:"omitempty,oneof=1 3"` // 1-list; 3-rent
	}

	QueryCollectionListByUserForm struct {
		UserAddr string `form:"user_addr" binding:"required"`
	}

	QueryNftDetailForm struct {
		NftId    int    `form:"nft_id" binding:"required,gte=1"`
		UserAddr string `form:"user_addr"`
	}

	QueryNftLendInfoForm struct {
		NftId int `form:"nft_id" binding:"required,gte=1"`
	}
)
