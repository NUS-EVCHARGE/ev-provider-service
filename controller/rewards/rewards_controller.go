package controller

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

type RewardsController interface {
	CreateCoinPolicy(coin dto.CoinPolicy) error
	UpdateCoinPolicy(coin dto.CoinPolicy) error
	GetCoinPolicy(providerId int) (dto.CoinPolicy, error)
	CreateVoucher(voucher dto.Vouchers) error
	UpdateVoucher(voucher dto.Vouchers) error
	GetAllVouchers(providerId int) ([]dto.Vouchers, error)
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
	_, err := dao.Db.CreateVouchers(voucher)
	return err
}

func (r *RewardsControllerImpl) UpdateVoucher(voucher dto.Vouchers) error {
	_, err := dao.Db.UpdateVoucher(voucher)
	return err
}

func (r *RewardsControllerImpl) GetAllVouchers(providerId int) ([]dto.Vouchers, error) {
	return dao.Db.GetAllVouchers(providerId)
}
