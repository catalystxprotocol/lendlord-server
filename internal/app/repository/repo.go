package repository

import (
	log "github.com/sirupsen/logrus"
)

type Repos struct {
	ChainConfigRepo   IChainConfigRepo
	NftRepo           INftRepo
	NftActivityRepo   INftActivityRepo
	NftCollectionRepo INftCollectionRepo
	NftOrderRepo      INftOrderRepo
	Log               *log.Logger
}
