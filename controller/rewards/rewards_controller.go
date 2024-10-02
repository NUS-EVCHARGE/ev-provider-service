package controller

import (
	"time"

	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/sirupsen/logrus"
)

type RewardsController interface {
	CreateCoinPolicy(coin dto.CoinPolicy) error
	UpdateCoinPolicy(coin dto.CoinPolicy) error
	GetCoinPolicy(providerId int) (dto.CoinPolicy, error)
	CreateVoucher(voucher dto.Vouchers) error
	UpdateVoucher(voucher dto.Vouchers) error
	GetAllVouchers(providerId int) ([]dto.Vouchers, error)
	GetVouchers(voucherId int) (dto.Vouchers, error)
	ExpireVouchers()
}

type RewardsControllerImpl struct {
}

var RewardsControllerObj RewardsController

func NewRewardController() {
	RewardsControllerObj = &RewardsControllerImpl{}
}

func (r *RewardsControllerImpl) CreateCoinPolicy(coin dto.CoinPolicy) error {
	_, err := dao.Db.CreateCoinPolicy(coin)
	return err
}

func (r *RewardsControllerImpl) UpdateCoinPolicy(coin dto.CoinPolicy) error {
	_, err := dao.Db.UpdateCoinPolicy(coin)
	return err
}

func (r *RewardsControllerImpl) GetCoinPolicy(providerId int) (dto.CoinPolicy, error) {
	return dao.Db.GetCoinPolicy(providerId)
}

func (r *RewardsControllerImpl) CreateVoucher(voucher dto.Vouchers) error {
	startDateObj, err := time.Parse(time.RFC3339, voucher.StartDate)
	if err != nil {
		return err
	}
	endDateObj, err := time.Parse(time.RFC3339, voucher.EndDate)
	if err != nil {
		return err
	}
	voucher.StartDateInUnix = startDateObj.UnixNano()
	voucher.EndDateInUnix = endDateObj.UnixNano()
	voucher.Status = true
	_, err = dao.Db.CreateVouchers(voucher)
	return err
}

func (r *RewardsControllerImpl) UpdateVoucher(voucher dto.Vouchers) error {
	startDateObj, err := time.Parse(time.RFC3339, voucher.StartDate)
	if err != nil {
		return err
	}
	endDateObj, err := time.Parse(time.RFC3339, voucher.EndDate)
	if err != nil {
		return err
	}
	voucher.StartDateInUnix = startDateObj.UnixNano()
	voucher.EndDateInUnix = endDateObj.UnixNano()
	_, err = dao.Db.UpdateVoucher(voucher)
	return err
}

func (r *RewardsControllerImpl) GetAllVouchers(providerId int) ([]dto.Vouchers, error) {
	voucherList, err := dao.Db.GetAllVouchers(providerId)
	if err != nil {
		return nil, err
	}
	for index, v := range voucherList {
		voucherList[index].Key = index
		voucherList[index].StartDate = time.Unix(0, v.StartDateInUnix).Format("2006-01-02 15:04:00")
		voucherList[index].EndDate = time.Unix(0, v.EndDateInUnix).Format("2006-01-02 15:04:00")
	}
	return voucherList, nil
}

func (r *RewardsControllerImpl) GetVouchers(voucherId int) (dto.Vouchers, error) {
	voucher, err := dao.Db.GetVoucher(voucherId)
	if err != nil {
		return dto.Vouchers{}, err
	}

	voucher.StartDate = time.Unix(0, voucher.StartDateInUnix).Format("2006-01-02 15:04:00")
	voucher.EndDate = time.Unix(0, voucher.EndDateInUnix).Format("2006-01-02 15:04:00")
	return voucher, nil
}

// this is a goroutine that sets voucher to be expired
func (r *RewardsControllerImpl) ExpireVouchers() {
	// update DB for every 10 seconds when a voucher has expired
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			// get current time
			currentTime := time.Now().UnixNano()
			if err := dao.Db.SetVouchersToBeExpired(currentTime); err != nil {
				logrus.WithField("err", err).Error("error_getting_expired_vouchers")
			}
		}
	}
}
