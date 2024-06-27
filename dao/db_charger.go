package dao

import (
	"fmt"

	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

// Charging Point
func (d *dbImpl) CreateChargerPointEntry(chargerPoint *dto.ChargerPoint) error {
	result := d.DbController.Create(&chargerPoint)
	return result.Error
}

func (d *dbImpl) GetChargerPointEntryByID(id uint) (dto.ChargerPoint, error) {
	var existingCharger dto.ChargerPoint

	results := d.DbController.Find(&existingCharger, "id = ?", id)
	return existingCharger, results.Error
}

func (d *dbImpl) GetChargerPointEntryByProviderID(providerId uint) ([]dto.ChargerPoint, error) {
	var existingCharger []dto.ChargerPoint

	results := d.DbController.Find(&existingCharger, "provider_id = ?", providerId)
	return existingCharger, results.Error
}

func (d *dbImpl) GetAllChargerPointEntry() ([]dto.ChargerPoint, error) {
	var existingCharger []dto.ChargerPoint

	results := d.DbController.Find(&existingCharger)
	return existingCharger, results.Error
}

func (d *dbImpl) GetAllChargerPointEntryByProviderID(providerId int) ([]dto.ChargerPoint, error) {
	var existingCharger []dto.ChargerPoint

	results := d.DbController.Where("provider_id = ?", providerId).Find(&existingCharger)
	return existingCharger, results.Error
}

func (d *dbImpl) UpdateChargerPointEntry(chargerPoint dto.ChargerPoint) error {
	results := d.DbController.Model(chargerPoint).Updates(chargerPoint)
	if results.RowsAffected == 0 {
		return fmt.Errorf("Charger not found")
	}
	return results.Error
}

func (d *dbImpl) GetChargerPointByLocation(providerId int, lat, lng float64) (dto.ChargerPoint, error) {
	var existingCharger dto.ChargerPoint

	results := d.DbController.Where("lat = ? and lng = ? and provider_id = ?", lat, lng, providerId).Find(&existingCharger)
	return existingCharger, results.Error
}

// Charger
func (d *dbImpl) CreateChargerEntry(charger dto.Charger) error {
	result := d.DbController.Create(&charger)
	return result.Error
}

func (d *dbImpl) GetChargerByChargerPointId(chargerPointId uint) ([]dto.Charger, error) {
	var existingCharger []dto.Charger

	results := d.DbController.Find(&existingCharger, "charger_point_id = ?", chargerPointId)
	return existingCharger, results.Error
}

func (d *dbImpl) GetChargerById(chargerId uint) (dto.Charger, error) {
	var existingCharger dto.Charger

	results := d.DbController.Find(&existingCharger, "id = ?", chargerId)
	return existingCharger, results.Error
}

func (d *dbImpl) UpdateChargerEntry(Charger dto.Charger) error {
	results := d.DbController.Model(Charger).Updates(Charger)
	if results.RowsAffected == 0 {
		return fmt.Errorf("Charger not found")
	}
	return results.Error
}
