package provider

import (
	"fmt"

	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

type ProviderController interface {
	CreateProvider(provider dto.Provider) (dto.Provider, error)
	GetProvider(email string) (dto.Provider, error)
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
	return dao.Db.CreateProviderEntry(provider)
}

func (p ProviderControllerImpl) GetProvider(email string) (dto.Provider, error) {
	return dao.Db.GetProviderEntry(email)
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
