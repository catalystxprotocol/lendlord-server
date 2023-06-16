package vo

type (
	PaymentTokenInfo struct {
		ChainId  string `json:"chain_id"`
		Symbol   string `json:"symbol"`
		Decimals uint   `json:"decimals"`
		Icon     string `json:"icon"`
	}

	PageInfo struct {
		PageSize         int   `json:"page_size"`
		PageNumber       int   `json:"page_number"`
		TotalRecordCount int64 `json:"total_record_count"`
		TotalPageCount   int64 `json:"total_page_count"`
	}

	BaseResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	Success struct {
		Data interface{} `json:"data"`
	}

	Fail struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	FailResp struct {
		Error Fail `json:"error"`
	}

	PaymentToken struct {
		Icon             string `json:"icon"`
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		Decimals         uint   `json:"decimals"`
		IsOrigin         bool   `json:"is_origin"`
		PaymentTokenAddr string `json:"payment_token_addr"`
	}
)
