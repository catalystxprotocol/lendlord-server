package vo

type (
	QueryChainConfigVO struct {
		ChainConfigList []ChainConfigList `json:"list"`
	}

	NativeCurrency struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals uint   `json:"decimals"`
		Icon     string `json:"icon"`
	}

	ContractAddress struct {
		MarketContractAddress string `json:"market_contract_address"`
		DonftContractAddress  string `json:"donft_contract_address"`
	}

	ChainConfigList struct {
		ChainId          string          `json:"chain_id"`
		ChainName        string          `json:"chain_name"`
		NativeCurrency   NativeCurrency  `json:"native_currency"`
		RPCUrl           string          `json:"rpc_url"`
		BlockExplorerUrl string          `json:"block_explorer_url"`
		TxHashRoute      string          `json:"tx_hash_route"`
		IconUrl          string          `json:"icon_url"`
		ContractAddress  ContractAddress `json:"contract_address"`
	}
)
