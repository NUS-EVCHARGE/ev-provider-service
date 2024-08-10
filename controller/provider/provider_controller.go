package provider

import (
	"fmt"

	controller "github.com/NUS-EVCHARGE/ev-provider-service/controller/rewards"
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/sirupsen/logrus"
)

type ProviderController interface {
	CreateProvider(provider dto.Provider) (dto.Provider, error)
	GetProvider(email string) (dto.Provider, error)
	IsCompanyExist(companyName string) bool
	GetAllProvider() ([]dto.Provider, error)
	DeleteProvider(providerId uint) error
	UpdateProvider(provider dto.Provider) error
}

type ProviderControllerImpl struct {
}

func (p ProviderControllerImpl) CreateProvider(provider dto.Provider) (dto.Provider, error) {
	if _, err := p.GetProvider(provider.UserEmail); err == nil {
		return dto.Provider{}, fmt.Errorf("provider already exist")
	}

	// create provider entry
	provider, err := dao.Db.CreateProviderEntry(provider)
	if err != nil {
		return dto.Provider{}, err
	}

	// create base license
	_, err = dao.Db.CreateLicense(dto.License{
		CompanyId: int(provider.ID),
		Standard:  0,
		Starter:   0,
		Premium:   0,
	})
	if err != nil {
		return dto.Provider{}, err
	}

	// init coin policy
	var status = false
	err = controller.RewardsControllerObj.CreateCoinPolicy(dto.CoinPolicy{Status: &status, ProviderId: int(provider.ID)})
	if err != nil {
		return dto.Provider{}, err
	}

	return provider, nil
}

func (p ProviderControllerImpl) GetProvider(email string) (dto.Provider, error) {
	return dao.Db.GetProviderEntry(email)
}

func (p ProviderControllerImpl) IsCompanyExist(companyName string) bool {
	_, err := dao.Db.GetProviderEntryByCompany(companyName)
	if err != nil {
		logrus.WithField("err", err).Error("error getting provider by company")
		return false
	}
	return true
}

func (p ProviderControllerImpl) GetAllProvider() ([]dto.Provider, error) {
	return dao.Db.GetAllProviderEntry()
}

func (p ProviderControllerImpl) DeleteProvider(providerId uint) error {
	return dao.Db.DeleteProviderEntry(dto.Provider{ID: providerId})
}

func (p ProviderControllerImpl) UpdateProvider(provider dto.Provider) error {
	return dao.Db.UpdateProviderEntry(provider)
}

var (
	ProviderControllerObj ProviderController
)

func NewProviderController() {
	ProviderControllerObj = &ProviderControllerImpl{}
}
