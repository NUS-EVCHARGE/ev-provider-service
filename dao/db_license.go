package dao

import (
	"fmt"

	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

func (d *dbImpl) CreateLicense(license dto.License) (dto.License, error) {
	results := d.DbController.Create(license)
	return license, results.Error
}

func (d *dbImpl) GetLicenseByCompanyId(companyId int) (dto.License, error) {
	var companyLicense dto.License
	results := d.DbController.Where("company_id = ?", companyId).First(&companyLicense)
	return companyLicense, results.Error
}

func (d *dbImpl) UpdateLicense(license dto.License) error {
	results := d.DbController.Where("company_id = ?", license.CompanyId).Updates(license)
	if results.RowsAffected == 0 {
		return fmt.Errorf("license not found or there is no change in data")
	}
	return results.Error
}
