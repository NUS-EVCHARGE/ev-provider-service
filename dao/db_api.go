package dao

import "github.com/NUS-EVCHARGE/ev-provider-service/dto"

func (d *dbImpl) GetChargerPointApi(id uint) (dto.ChargerPoint, error) {
	var existingCharger dto.ChargerPoint

	results := d.DbController.Find(&existingCharger, "id = ?", id)
	return existingCharger, results.Error
}
