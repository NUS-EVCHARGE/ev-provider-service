package provider

import (
	"fmt"
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	userDto "github.com/NUS-EVCHARGE/ev-user-service/dto"
)

type ProviderController interface {
	CreateProvider(provider dto.Provider, user userDto.User) (dto.Provider, error)
	GetProvider(user userDto.User) (dto.Provider, error)
	DeleteProvider(providerId uint) error
	UpdateProvider(provider dto.Provider) error
}

type ProviderControllerImpl struct {
}

func (p ProviderControllerImpl) CreateProvider(provider dto.Provider, user userDto.User) (dto.Provider, error) {
	if _, err := p.GetProvider(user); err == nil {
		return dto.Provider{}, fmt.Errorf("provider already exist")
	}
	return dao.Db.CreateProviderEntry(provider)
}

func (p ProviderControllerImpl) GetProvider(user userDto.User) (dto.Provider, error) {
	return dao.Db.GetAllProviderEntry(user.Email)
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
