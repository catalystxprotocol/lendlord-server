package services

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/lendlord/lendlord-server/internal/app/constant"
	"github.com/lendlord/lendlord-server/internal/app/enum"
	"github.com/lendlord/lendlord-server/internal/app/error"
	"github.com/lendlord/lendlord-server/internal/app/models/form"
	"github.com/lendlord/lendlord-server/internal/app/models/vo"
	"github.com/lendlord/lendlord-server/internal/app/repository"
	"gorm.io/gorm"
	"strings"
)

type INftService interface {
	QueryNftCollectionListByUser(params *form.QueryCollectionListByUserForm) (*vo.QueryCollectionByUserVo, error.Error)
	PageQueryNftListByUser(params *form.PageQueryNftListByUserForm) (*vo.PageQueryNftByUserVo, error.Error)
	QueryNftDetail(params *form.QueryNftDetailForm) (*vo.QueryNftDetailVo, error.Error)
	PageQueryNftActivityList(params *form.PageQueryNftActivityListForm) (*vo.PageQueryNftActivityVo, error.Error)
	QueryNftLendInfo(params *form.QueryNftLendInfoForm) (*vo.QueryNftLendInfoVo, error.Error)
	QueryNftCollectionSum() (*vo.QueryNftCollectionSumVo, error.Error)
	//QueryNFTDetailByID() (*vo.QueryChainConfigVO, error.Error)
	//QueryNFTActivities() (*vo.QueryChainConfigVO, error.Error)
	//QueryNFTLendInfo() (*vo.QueryChainConfigVO, error.Error)
}

type NftService struct {
	nftRepo           repository.INftRepo
	nftActivityRepo   repository.INftActivityRepo
	nftCollectionRepo repository.INftCollectionRepo
	nftOrderRepo      repository.INftOrderRepo
	log               *log.Logger
}

func NewNftService(nftRepo repository.INftRepo, nftActivityRepo repository.INftActivityRepo,
	nftCollectionRepo repository.INftCollectionRepo, nftOrderRepo repository.INftOrderRepo, log *log.Logger) *NftService {
	return &NftService{
		nftRepo:           nftRepo,
		nftActivityRepo:   nftActivityRepo,
		nftCollectionRepo: nftCollectionRepo,
		nftOrderRepo:      nftOrderRepo,
		log:               log,
	}
}

func (n NftService) QueryNftCollectionListByUser(params *form.QueryCollectionListByUserForm) (*vo.QueryCollectionByUserVo, error.Error) {
	var collections = make([]vo.CollectionByUserList, 0)
	logFields := log.Fields{}
	logFields["model"] = "NftService"
	logFields["func"] = "QueryNftCollectionList"

	resp, err := n.nftCollectionRepo.GetNftCollectionsByUser(strings.ToLower(params.UserAddr))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		n.log.WithFields(logFields).Errorf("nft collection list query failed: %v", err.Error())
		return nil, error.WrapErrorDetail(error.CodeInternalError, error.InternalServerError)
	}

	for _, collection := range resp {
		collections = append(collections, vo.CollectionByUserList{
			Id:   collection.ID,
			Name: collection.Name,
		})
	}

	return &vo.QueryCollectionByUserVo{CollectionByUserList: collections}, nil
}

func (n NftService) PageQueryNftListByUser(params *form.PageQueryNftListByUserForm) (*vo.PageQueryNftByUserVo, error.Error) {
	nfts := make([]vo.PageQueryNftByUserList, 0)
	logFields := log.Fields{}
	logFields["model"] = "NftService"
	logFields["func"] = "PageQueryNftListByUser"

	if params.PageNumber <= 0 {
		params.PageNumber = enum.DefaultPage
	}

	if params.PageSize <= 0 {
		params.PageSize = enum.DefaultSize
	}

	params.UserAddr = strings.ToLower(params.UserAddr)

	total, totalPage, resp, err := n.nftRepo.GetNftListByUser(params)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		n.log.WithFields(logFields).Errorf("page query nft list failed: %v", err.Error())
		return nil, error.WrapErrorDetail(error.CodeInternalError, error.InternalServerError)
	}
	pageInfo := vo.PageInfo{
		PageSize:         params.PageSize,
		PageNumber:       params.PageNumber,
		TotalRecordCount: total,
		TotalPageCount:   totalPage,
	}

	for _, item := range resp {
		status := convertStatus(params.Status, params.UserAddr, item.NftOwner, item.RenterAddr, item.LenderAddr, item.OrderId,
			item.ShelfStatus, item.RenterStartTimestamp, item.RenterEndTimestamp, item.LenderStartTimestamp, item.LenderEndTimestamp)
		nft := vo.PageQueryNftByUserList{
			NftId:            item.NftId,
			NftTokenId:       item.TokenId,
			NftAddr:          item.NftAddr,
			NftOwner:         item.NftOwner,
			ChainId:          item.ChainId,
			NftStatus:        status,
			CollectionName:   item.CollectionName,
			EndTimestamp:     item.RenterEndTimestamp,
			LenderPrice:      item.LenderPrice,
			RentMinDuration:  item.RentMinDuration,
			LendEndTimestamp: item.LenderEndTimestamp,
			UseLink:          item.UseLink,
			NftImage:         item.NftImage,
			NftName:          item.NftName,
			ONftId:           item.ONftId,
			PaymentTokenInfo: vo.PaymentTokenInfo{
				ChainId:  item.ChainId,
				Symbol:   item.Symbol,
				Decimals: item.Decimals,
				Icon:     item.Icon,
			},
		}
		nfts = append(nfts, nft)
	}

	return &vo.PageQueryNftByUserVo{
		PageInfo:               pageInfo,
		PageQueryNftByUserList: nfts,
	}, nil
}

func (n NftService) QueryNftDetail(params *form.QueryNftDetailForm) (*vo.QueryNftDetailVo, error.Error) {
	logFields := log.Fields{}
	logFields["model"] = "NftService"
	logFields["func"] = "QueryNftDetail"

	params.UserAddr = strings.ToLower(params.UserAddr)

	resp, err := n.nftRepo.QueryNftDetail(params.NftId)
	if err != nil {
		n.log.WithFields(logFields).Errorf("nft detail query failed: %v", err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.WrapErrorDetail(error.CodeNotFound, error.ResourceNotFound)
		} else {
			return nil, error.WrapErrorDetail(error.CodeInternalError, error.InternalServerError)
		}

	}

	status := detailConvertStatus(params.UserAddr, resp.NftOwner, resp.RenterAddr, resp.LenderAddr, resp.OrderId, resp.ShelfStatus,
		resp.RenterStartTimestamp, resp.RenterEndTimestamp, resp.LenderStartTimestamp, resp.LenderEndTimestamp)

	owner := ""
	switch status {
	case enum.Unknown, enum.Idle, enum.IdleOwner:
		owner = resp.NftOwner
	default:
		owner = resp.DonftAddr
	}

	user := ""
	if status == enum.ListDeal {
		user = resp.RenterAddr
	}

	return &vo.QueryNftDetailVo{
		ChainId:          resp.ChainId,
		NftAddr:          resp.NftAddr,
		NftId:            resp.NftId,
		NftTokenId:       resp.TokenId,
		NftOwner:         owner,
		LenderAddr:       resp.LenderAddr,
		NftUser:          user,
		ONftId:           resp.ONftId,
		ONftAddr:         resp.ONftAddr,
		EndTimestamp:     resp.RenterEndTimestamp,
		Price:            resp.Price,
		RentMinDuration:  resp.LenderMinDuration,
		LendEndTimestamp: resp.LenderEndTimestamp,
		NftStatus:        status,
		NftImage:         resp.NftImage,
		NftName:          resp.NftName,
		CollectionName:   resp.CollectionName,
		PaymentTokenInfo: vo.PaymentTokenInfo{
			ChainId:  resp.ChainId,
			Symbol:   resp.Symbol,
			Decimals: resp.Decimals,
			Icon:     resp.Icon,
		},
	}, nil
}

func (n NftService) PageQueryNftActivityList(params *form.PageQueryNftActivityListForm) (*vo.PageQueryNftActivityVo, error.Error) {
	activities := make([]vo.PageQueryNftActivityList, 0)
	logFields := log.Fields{}
	logFields["model"] = "NftService"
	logFields["func"] = "PageQueryNftActivityList"

	//if params.EventType != "" && params.EventType != enum.Listed && params.EventType != enum.Rented {
	//	n.log.WithFields(logFields).Errorf("params illegal, the params is %v\n", params.EventType)
	//	return nil, error.WrapErrorDetail(error.CodeInvalidParam, error.ParamsIllegal)
	//}

	if params.PageNumber <= 0 {
		params.PageNumber = enum.DefaultPage
	}

	if params.PageSize <= 0 {
		params.PageSize = enum.DefaultSize
	}

	total, totalPage, resp, err := n.nftActivityRepo.GetNftActivityList(params)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		n.log.WithFields(logFields).Errorf("page query nft activity list failed: %v", err.Error())
		return nil, error.WrapErrorDetail(error.CodeInternalError, error.InternalServerError)
	}
	pageInfo := vo.PageInfo{
		PageSize:         params.PageSize,
		PageNumber:       params.PageNumber,
		TotalRecordCount: total,
		TotalPageCount:   totalPage,
	}

	for _, item := range resp {
		activity := vo.PageQueryNftActivityList{
			Event:            item.Event,
			Lender:           item.LenderAddr,
			Renter:           item.RenterAddr,
			Price:            item.Price,
			RentMinDuration:  item.Duration,
			LendEndTimestamp: item.EndTimestamp,
			StartTimestamp:   item.StartTimestamp,
			PaymentTokenInfo: vo.PaymentTokenInfo{
				ChainId:  item.ChainId,
				Symbol:   item.Symbol,
				Decimals: item.Decimals,
				Icon:     item.Icon,
			},
		}
		activities = append(activities, activity)
	}

	return &vo.PageQueryNftActivityVo{
		PageInfo:                 pageInfo,
		PageQueryNftActivityList: activities,
	}, nil
}

func (n NftService) QueryNftLendInfo(params *form.QueryNftLendInfoForm) (*vo.QueryNftLendInfoVo, error.Error) {
	logFields := log.Fields{}
	logFields["model"] = "NftService"
	logFields["func"] = "QueryNftLendInfo"

	resp, err := n.nftRepo.GetNftLendInfo(params.NftId)
	if err != nil {
		n.log.WithFields(logFields).Errorf("query nft lendInfo failed: %v", err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.WrapErrorDetail(error.CodeNotFound, error.ResourceNotFound)
		} else {
			return nil, error.WrapErrorDetail(error.CodeInternalError, error.InternalServerError)
		}

	}

	//nftOwner := ""
	//if resp.ShelfStatus == enum.Zero || resp.ShelfStatus == enum.Redeem {
	//	nftOwner = resp.NftOwner
	//} else {
	//	nftOwner = resp.DonftContractAddr
	//}

	return &vo.QueryNftLendInfoVo{
		CollectionName: resp.CollectionName,
		NftTokenId:     resp.TokenId,
		NftImage:       resp.NftImage,
		NftOwner:       resp.NftOwner,
		ONftId:         resp.TokenId,
		ONftAddr:       resp.NftAddr,
		ChainID:        resp.ChainId,
		DonftId:        resp.DonftId,
		DoNftAddr:      resp.DonftContractAddr,
		LenderAddr:     resp.LenderAddr,
		PaymentTokenList: []vo.PaymentToken{{
			Icon:             resp.Icon,
			Symbol:           resp.Symbol,
			Name:             resp.TokenName,
			Decimals:         resp.Decimals,
			IsOrigin:         true,
			PaymentTokenAddr: constant.OriginPaymentToken,
		}},
	}, nil
}

func (n NftService) QueryNftCollectionSum() (*vo.QueryNftCollectionSumVo, error.Error) {
	logFields := log.Fields{}
	logFields["model"] = "NftService"
	logFields["func"] = "QueryNftCollectionSum"

	count, err := n.nftCollectionRepo.GetNftCollectionSum()
	if err != nil {
		n.log.WithFields(logFields).Errorf("query nft collection query failed: %v", err.Error())
		return nil, error.WrapErrorDetail(error.CodeInternalError, error.InternalServerError)
	}
	return &vo.QueryNftCollectionSumVo{TotalSum: count}, nil
}
