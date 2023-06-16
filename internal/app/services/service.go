package services

import (
	"time"

	"github.com/lendlord/lendlord-server/internal/app/enum"
	"github.com/lendlord/lendlord-server/internal/app/models/entity"
)

type Service struct {
	ChainConfigService *ChainConfigService
	NftService         *NftService
}

// NFT列表接口状态转换函数  //inputStatus: 1-idle; 2-list; 4-rent
func convertStatus(inputStatus int, userAddr, nftOwner, renterAddr, lenderAddr string, orderId, shelfStatus uint,
	renterStartTimestamp, renterEndTimestamp, lenderStartTimestamp, lenderEndTimestamp int64) (status uint) {
	now := time.Now().Unix()

	if shelfStatus == entity.List {
		// lender日期过期
		if !(lenderStartTimestamp <= now && now <= lenderEndTimestamp) {
			return enum.Idle
		}

		if userAddr == lenderAddr && renterAddr == "" {
			return enum.Listed
		}

		if userAddr == lenderAddr && renterStartTimestamp < now && now <= renterEndTimestamp {
			if lenderAddr == renterAddr {
				if inputStatus == enum.Listed {
					return enum.ListDeal
				} else if inputStatus == enum.Rented {
					return enum.Rented
				}
			}
			return enum.ListDeal
		}

		if userAddr == renterAddr && renterStartTimestamp < now && now <= renterEndTimestamp {
			return enum.Rented
		}
		return enum.Listed
	}

	if userAddr == nftOwner && (shelfStatus == entity.DeList || shelfStatus == entity.Redeem || orderId == 0) {
		return enum.Idle
	}

	return enum.Idle
}

// NFT详情接口状态转换函数
func detailConvertStatus(userAddr, nftOwner, renterAddr, lenderAddr string, orderId, shelfStatus uint,
	renterStartTimestamp, renterEndTimestamp, lenderStartTimestamp, lenderEndTimestamp int64) (status uint) {
	now := time.Now().Unix()

	if shelfStatus == entity.List {
		// lender日期过期
		if !(lenderStartTimestamp <= now && now <= lenderEndTimestamp) {
			if userAddr == nftOwner || userAddr == lenderAddr {
				return enum.IdleOwner
			}
			return enum.Idle
		}
		if renterAddr == "" {
			return enum.Listed
		}

		if renterStartTimestamp < now && now <= renterEndTimestamp {
			return enum.ListDeal
		}
		return enum.Listed
	}

	if shelfStatus == entity.DeList || shelfStatus == entity.Redeem || orderId == 0 {
		if userAddr != "" && (userAddr == nftOwner || userAddr == lenderAddr) {
			return enum.IdleOwner
		}
		return enum.Idle
	}

	return enum.Idle
}
