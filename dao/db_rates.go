package dao

import (
	"fmt"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

func (d *dbImpl) UpdateRatesEntry(Rates dto.Rates) error {
	results := d.DbController.Model(Rates).Updates(Rates)
	if results.RowsAffected == 0 {
		return fmt.Errorf("Rates not found")
	}
	return results.Error
}

func (d *dbImpl) CreateRatesEntry(Rates *dto.Rates) error {
	result := d.DbController.Create(&Rates)
	return result.Error
}

func (d *dbImpl) DeleteRatesEntry(Rates dto.Rates) error {
	results := d.DbController.Delete(&Rates)
	if results.RowsAffected == 0 {
		return fmt.Errorf("Rates not found")
	}
	return results.Error
}

func (d *dbImpl) GetRatesByProviderId(providerId uint) ([]dto.Rates, error) {
	var ratesList []dto.Rates
	results := d.DbController.Find(&ratesList, "provider_id = ?", providerId)
	return ratesList, results.Error
}

func (d *dbImpl) GetRatesByRateId(rateId uint) (dto.Rates, error) {
	var rates dto.Rates
	results := d.DbController.Find(&rates, "id = ?", rateId)
	return rates, results.Error
}
