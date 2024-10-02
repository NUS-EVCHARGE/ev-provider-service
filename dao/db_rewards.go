package dao

import "github.com/NUS-EVCHARGE/ev-provider-service/dto"

func (db *dbImpl) CreateCoinPolicy(coinPolicy dto.CoinPolicy) (dto.CoinPolicy, error) {
	result := db.DbController.Create(&coinPolicy)
	return coinPolicy, result.Error
}

func (db *dbImpl) UpdateCoinPolicy(coinPolicy dto.CoinPolicy) (dto.CoinPolicy, error) {
	result := db.DbController.Updates(&coinPolicy)
	return coinPolicy, result.Error
}

func (db *dbImpl) GetCoinPolicy(providerId int) (dto.CoinPolicy, error) {
	var coinPolicy dto.CoinPolicy
	result := db.DbController.Where("provider_id = ?", providerId).Find(&coinPolicy)
	return coinPolicy, result.Error
}

func (db *dbImpl) CreateVouchers(voucher dto.Vouchers) (dto.Vouchers, error) {
	result := db.DbController.Create(&voucher)
	return voucher, result.Error
}

func (db *dbImpl) GetAllVouchers(providerId int) ([]dto.Vouchers, error) {
	var voucherList []dto.Vouchers
	result := db.DbController.Where("provider_id = ?", providerId).Find(&voucherList)
	return voucherList, result.Error
}

func (db *dbImpl) GetVoucher(voucherId int) (dto.Vouchers, error) {
	var voucher dto.Vouchers
	result := db.DbController.Where("id = ?", voucherId).Find(&voucher)
	return voucher, result.Error
}

func (db *dbImpl) UpdateVoucher(voucher dto.Vouchers) (dto.Vouchers, error) {
	result := db.DbController.Updates(&voucher)
	return voucher, result.Error
}

func (db *dbImpl) SetVouchersToBeExpired(currentTime int64) error {
	result := db.DbController.Where("end_date <= ?", currentTime).Table("voucher_tab").Update("status", 0)
	return result.Error
}
