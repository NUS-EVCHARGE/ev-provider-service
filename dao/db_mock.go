package dao

import (
	"fmt"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

type mockDbImpl struct {
	providerList []dto.Provider
	ratesList    []dto.Rates
	chargerList  []dto.Charger
}

func (m *mockDbImpl) GetAllCharger() ([]dto.Charger, error) {
	return m.chargerList, nil
}

func (m *mockDbImpl) CreateProviderEntry(provider dto.Provider) (dto.Provider, error) {
	m.providerList = append(m.providerList, provider)
	provider.ID = uint(len(m.providerList) - 1)
	return provider, nil
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
	m.providerList = append(m.providerList[:int(provider.ID)], m.providerList[int(provider.ID)+1:]...)
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

func (m *mockDbImpl) CreateChargerEntry(charger *dto.Charger) error {
	m.chargerList = append(m.chargerList, *charger)
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
	m.chargerList = append(m.chargerList[:int(charger.ID)], m.chargerList[int(charger.ID)+1:]...)
	return nil
}

func (m *mockDbImpl) GetChargerEntryByProvider(providerId uint) ([]dto.Charger, error) {
	var chargerList []dto.Charger

	for _, c := range m.chargerList {
		if c.ProviderId == providerId {
			chargerList = append(chargerList, c)
		}
	}

	return chargerList, nil
}

func (m *mockDbImpl) CreateRatesEntry(rates *dto.Rates) error {
	m.ratesList = append(m.ratesList, *rates)
	return nil
}

func (m *mockDbImpl) UpdateRatesEntry(rates dto.Rates) error {
	if len(m.ratesList) <= int(rates.ID) {
		return fmt.Errorf("rates not found")
	}
	m.ratesList[rates.ID] = rates
	return nil
}

func (m *mockDbImpl) DeleteRatesEntry(rates dto.Rates) error {
	if len(m.ratesList) == 1 {
		m.ratesList = []dto.Rates{}
		return nil
	}
	m.ratesList = append(m.ratesList[:int(rates.ID)], m.ratesList[int(rates.ID)+1:]...)
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

func (m *mockDbImpl) GetRatesByRateId(rateId uint) (dto.Rates, error) {
	if len(m.ratesList) <= int(rateId) {
		return dto.Rates{}, fmt.Errorf("rates not found")
	}
	return m.ratesList[rateId], nil
}

func (m *mockDbImpl) GetChargerById(chargerId uint) (dto.Charger, error) {
	if len(m.chargerList) <= int(chargerId) {
		return dto.Charger{}, fmt.Errorf("charger not found")
	}
	return m.chargerList[chargerId], nil
}

func NewMockDatabase(providerList []dto.Provider, ratesList []dto.Rates, chargerList []dto.Charger) Database {
	return &mockDbImpl{
		providerList: providerList,
		ratesList:    ratesList,
		chargerList:  chargerList,
	}
}
