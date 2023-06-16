package services

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/lendlord/lendlord-server/internal/app/error"
	"github.com/lendlord/lendlord-server/internal/app/models/vo"
	"github.com/lendlord/lendlord-server/internal/app/repository"
	"gorm.io/gorm"
)

type IChainConfigService interface {
	QueryChainConfigs() (*vo.QueryChainConfigVO, error.Error)
}

type ChainConfigService struct {
	chainConfigRepo repository.IChainConfigRepo
	log             *log.Logger
}

func NewChainConfigService(chainConfigRepo repository.IChainConfigRepo, log *log.Logger) *ChainConfigService {
	return &ChainConfigService{
		chainConfigRepo: chainConfigRepo,
		log:             log,
	}
}

func (c *ChainConfigService) QueryChainConfigs() (*vo.QueryChainConfigVO, error.Error) {
	var chainConfigs = make([]vo.ChainConfigList, 0)
	logFields := log.Fields{}
	logFields["model"] = "ChainConfigService"
	logFields["func"] = "GetChainConfigs"

	resp, err := c.chainConfigRepo.GetChainConfigs()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.log.WithFields(logFields).Errorf("the chain config query failed: %v", err.Error())
		return nil, error.WrapErrorDetail(error.CodeInternalError, error.InternalServerError)
	}

	for _, chainConfig := range resp {
		cc := vo.ChainConfigList{
			ChainId:   chainConfig.ChainId,
			ChainName: chainConfig.ChainName,
			NativeCurrency: vo.NativeCurrency{
				Name:     chainConfig.TokenName,
				Symbol:   chainConfig.TokenSymbol,
				Decimals: chainConfig.Decimals,
				Icon:     chainConfig.TokenIcon,
			},
			RPCUrl:           chainConfig.RpcUrl,
			BlockExplorerUrl: chainConfig.BlockExplorerUrl,
			TxHashRoute:      chainConfig.TxHashRoute,
			IconUrl:          chainConfig.ChainIconUrl,
			ContractAddress: vo.ContractAddress{
				MarketContractAddress: chainConfig.MarketContractAddr,
				DonftContractAddress:  chainConfig.DonftContractAddr,
			},
		}

		chainConfigs = append(chainConfigs, cc)
	}
	return &vo.QueryChainConfigVO{ChainConfigList: chainConfigs}, nil
}
