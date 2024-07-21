package license

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

type LicenseController interface {
	UpdateLicense(license dto.License, companyName string) error
	GetLicenseByCompany(companyName string) (dto.License, error)
}

type LicenseControllerImpl struct {
}

// GetLicenseByCompany implements LicenseController.
func (*LicenseControllerImpl) GetLicenseByCompany(companyName string) (dto.License, error) {
	provider, err := dao.Db.GetProviderEntryByCompany(companyName)
	if err != nil {
		return dto.License{}, err
	}
	license, err := dao.Db.GetLicenseByCompanyId(int(provider.ID))
	if err != nil {
		return dto.License{}, err
	}
	return license, nil
}

// UpdateLicense implements LicenseController.
func (*LicenseControllerImpl) UpdateLicense(license dto.License, companyName string) error {
	provider, err := dao.Db.GetProviderEntryByCompany(companyName)
	if err != nil {
		return err
	}
	license.CompanyId = int(provider.ID)
	if err = dao.Db.UpdateLicense(license); err != nil {
		return err
	}

	return nil
}

var (
	LicenseControllerObj LicenseController
)

func NewLicenseController() {
	LicenseControllerObj = &LicenseControllerImpl{}
}
