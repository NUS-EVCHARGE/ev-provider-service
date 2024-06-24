package dao

import (
	"fmt"

	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

func (d *dbImpl) UpdateProviderEntry(Provider dto.Provider) error {
	results := d.DbController.Model(Provider).Updates(Provider)
	if results.RowsAffected == 0 {
		return fmt.Errorf("Provider not found")
	}
	return results.Error
}

func (d *dbImpl) CreateProviderEntry(Provider dto.Provider) (dto.Provider, error) {
	result := d.DbController.Create(&Provider)
	return Provider, result.Error
}

func (d *dbImpl) DeleteProviderEntry(Provider dto.Provider) error {
	results := d.DbController.Delete(&Provider)
	if results.RowsAffected == 0 {
		return fmt.Errorf("Provider not found")
	}
	return results.Error
}

func (d *dbImpl) GetProviderEntry(email string) (dto.Provider, error) {
	var existingProvider dto.Provider

	results := d.DbController.First(&existingProvider, "user_email = ?", email)
	return existingProvider, results.Error
}

func (d *dbImpl) GetProviderEntryByCompany(companyName string) (dto.Provider, error) {
	var existingProvider dto.Provider

	results := d.DbController.Where("company_name = ?", companyName).First(&existingProvider)
	return existingProvider, results.Error
}

func (d *dbImpl) GetAllProviderEntry() ([]dto.Provider, error) {
	var existingProvider []dto.Provider

	results := d.DbController.Find(&existingProvider)
	return existingProvider, results.Error
}
