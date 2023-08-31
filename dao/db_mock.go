package dao

import (
	"ev-provider-service/dto"
	"fmt"
)

type mockDbImpl struct {
	providerList []dto.Provider
	ratesList []dto.Rates
	chargerList []dto.Charger
}

func (m *mockDbImpl) CreateProviderEntry(provider dto.Provider) error {
	m.providerList = append(m.providerList, provider)
	return nil
}

func (m *mockDbImpl) UpdateProviderEntry(provider dto.Provider) error {
	if len(m.providerList) <= int(provider.ID) {
		return fmt.Errorf("provider not found")
	}
	m.providerList[int(provider.ID)] = provider
	return nil
}

func (m *mockDbImpl) DeleteProviderEntry(provider dto.Provider) error {
	if len(m.providerList) == 1 {
		m.providerList = []dto.Provider{}
		return nil
	}
	m.providerList = append(m.providerList[:int(provider.ID)],m.providerList[int(provider.ID) + 1:]...)
	return nil
}

func (m *mockDbImpl) GetAllProviderEntry(email string) (dto.Provider, error) {
	for _, p := range m.providerList {
		if p.UserEmail == email {
			return p, nil
		}
	}
	return dto.Provider{}, fmt.Errorf("provider not found")
}

func (m *mockDbImpl) CreateChargerEntry(charger dto.Charger) error {
	m.chargerList = append(m.chargerList, charger)
	return nil
}

func (m *mockDbImpl) UpdateChargerEntry(charger dto.Charger) error {
	if len(m.chargerList) <= int(charger.ID) {
		return fmt.Errorf("charger not found")
	}
	m.chargerList[int(charger.ID)] = charger
	return nil
}

func (m *mockDbImpl) DeleteChargerEntry(charger dto.Charger) error {
	if len(m.chargerList) == 1 {
		m.chargerList = []dto.Charger{}
		return nil
	}
	m.chargerList = append(m.chargerList[:int(charger.ID)],m.chargerList[int(charger.ID) + 1:]...)
	return nil
}

func (m *mockDbImpl) GetAllChargerEntry(providerId uint) ([]dto.Charger, error) {
	var chargerList []dto.Charger

	for  _, c := range m.chargerList {
		if c.ProviderId == providerId {
			chargerList = append(chargerList, c)
		} 
	}
	
	return chargerList, nil
}

func (m *mockDbImpl) CreateRatesEntry(rates dto.Rates) error {
	m.ratesList = append(m.ratesList, rates)
	return nil
}

func (m *mockDbImpl) UpdateRatesEntry(rates dto.Rates) error {
	if len(m.ratesList) <= int(rates.ID) {
		return fmt.Errorf("rates not found")
	}
	return nil
}

func (m *mockDbImpl) DeleteRatesEntry(rates dto.Rates) error {
	if len(m.ratesList) == 1 {
		m.ratesList = []dto.Rates{}
		return nil
	}
	m.ratesList = append(m.ratesList[:int(rates.ID)],m.ratesList[int(rates.ID) + 1:]...)
	return nil
}

func (m *mockDbImpl) GetRatesByProviderId(providerId uint) ([]dto.Rates, error) {
	var ratesList = []dto.Rates{}
	for _, r := range m.ratesList {
		if r.ProviderId == providerId {
			ratesList = append(ratesList, r)
		}
	}
	return ratesList, nil
}

func NewMockDatabase(providerList []dto.Provider, ratesList []dto.Rates, chargerList []dto.Charger) Database {
	return &mockDbImpl{
		providerList: providerList,
		ratesList: ratesList,
		chargerList: chargerList,
	}
}
