package dao

import (
	"fmt"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

func (d *dbImpl) UpdateChargerEntry(Charger dto.Charger) error {
	results := d.DbController.Model(Charger).Updates(Charger)
	if results.RowsAffected == 0 {
		return fmt.Errorf("Charger not found")
	}
	return results.Error
}

func (d *dbImpl) CreateChargerEntry(Charger dto.Charger) error {
	result := d.DbController.Create(&Charger)
	return result.Error
}

func (d *dbImpl) DeleteChargerEntry(Charger dto.Charger) error {
	results := d.DbController.Delete(&Charger)
	if results.RowsAffected == 0 {
		return fmt.Errorf("Charger not found")
	}
	return results.Error
}

func (d *dbImpl) GetChargerEntryByProvider(providerId uint) ([]dto.Charger, error) {
	var existingCharger []dto.Charger

	results := d.DbController.Find(&existingCharger, "provider_id = ?", providerId)
	return existingCharger, results.Error
}

func (d *dbImpl) GetAllCharger() ([]dto.Charger, error) {
	var existingCharger []dto.Charger

	results := d.DbController.Find(&existingCharger)
	return existingCharger, results.Error
}
