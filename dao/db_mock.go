package dao

import (
	"fmt"

	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

type mockDbImpl struct {
	providerList     []dto.Provider
	chargerPointList []*dto.ChargerPoint
	chargerList      []dto.Charger
}

// GetAllChargerPointEntryByProviderID implements Database.
func (*mockDbImpl) GetAllChargerPointEntryByProviderID(providerId int) ([]dto.ChargerPoint, error) {
	panic("unimplemented")
}

// GetProviderEntryByCompany implements Database.
func (*mockDbImpl) GetProviderEntryByCompany(companyName string) (dto.Provider, error) {
	panic("unimplemented")
}

// GetChargerPointByLocation implements Database.
func (*mockDbImpl) GetChargerPointByLocation(providerId int, lat float64, lng float64) (dto.ChargerPoint, error) {
	panic("unimplemented")
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
func (m *mockDbImpl) GetAllProviderEntry() ([]dto.Provider, error) {
	return m.providerList, nil
}

func (m *mockDbImpl) GetProviderEntry(email string) (dto.Provider, error) {
	for _, p := range m.providerList {
		if p.UserEmail == email {
			return p, nil
		}
	}
	return dto.Provider{}, fmt.Errorf("provider not found")
}

// charger points
// CreateChargerPointEntry implements Database.
func (m *mockDbImpl) CreateChargerPointEntry(chargerPoint *dto.ChargerPoint) error {
	m.chargerPointList = append(m.chargerPointList, chargerPoint)
	return nil
}

// GetAllChargerPointEntry implements Database.
func (m *mockDbImpl) GetAllChargerPointEntry() ([]dto.ChargerPoint, error) {
	charger := []dto.ChargerPoint{}
	for _, c := range m.chargerPointList {
		charger = append(charger, *c)
	}
	return charger, nil
}

// GetChargerPointEntryByID implements Database.
func (m *mockDbImpl) GetChargerPointEntryByID(chargerId uint) (dto.ChargerPoint, error) {
	if len(m.chargerPointList) <= int(chargerId) {
		return dto.ChargerPoint{}, fmt.Errorf("charger not found")
	}
	return *m.chargerPointList[int(chargerId)], nil
}

// GetChargerPointEntryByProviderID implements Database.
func (m *mockDbImpl) GetChargerPointEntryByProviderID(providerId uint) ([]dto.ChargerPoint, error) {
	var cpList = []dto.ChargerPoint{}
	for _, cp := range m.chargerPointList {
		if cp.ProviderId == providerId {
			cpList = append(cpList, *cp)
		}
	}
	return cpList, nil
}

// UpdateChargerPointEntry implements Database.
func (m *mockDbImpl) UpdateChargerPointEntry(chargerPoint dto.ChargerPoint) error {
	if len(m.chargerPointList) <= int(chargerPoint.ID) {
		return fmt.Errorf("charger not found")
	}
	m.chargerPointList[int(chargerPoint.ID)] = &chargerPoint
	return nil
}

// charger
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

// GetChargerByChargerPointId implements Database.
func (m *mockDbImpl) GetChargerByChargerPointId(chargerPointId uint) ([]dto.Charger, error) {
	var chargerList []dto.Charger

	for _, c := range m.chargerList {
		if c.ChargerPointID == chargerPointId {
			chargerList = append(chargerList, c)
		}
	}
	return chargerList, nil
}
func (m *mockDbImpl) GetChargerById(chargerId uint) (dto.Charger, error) {
	if len(m.chargerList) <= int(chargerId) {
		return dto.Charger{}, fmt.Errorf("charger not found")
	}
	return m.chargerList[chargerId], nil
}

func NewMockDatabase(providerList []dto.Provider, chargerPointList []*dto.ChargerPoint, chargerList []dto.Charger) Database {
	return &mockDbImpl{
		providerList:     providerList,
		chargerList:      chargerList,
		chargerPointList: chargerPointList,
	}
}
