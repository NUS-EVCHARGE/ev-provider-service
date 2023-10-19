package charger

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

type ChargerController interface {
	CreateCharger(charger *dto.Charger) error
	UpdateCharger(charger dto.Charger) error
	GetChargerByProvider(providerId uint) ([]dto.Charger, error)
	GetAllCharger() ([]dto.Charger, error)
	GetChargerById(chargerId uint) (dto.Charger, error)

	DeleteCharger(chargerId uint) error
}

type ChargerImpl struct {
}

func (c *ChargerImpl) CreateCharger(charger *dto.Charger) error {
	return dao.Db.CreateChargerEntry(charger)
}

func (c *ChargerImpl) UpdateCharger(charger dto.Charger) error {
	return dao.Db.UpdateChargerEntry(charger)
}

func (c *ChargerImpl) GetChargerByProvider(providerId uint) ([]dto.Charger, error) {
	return dao.Db.GetChargerEntryByProvider(providerId)
}

func (c *ChargerImpl) GetAllCharger() ([]dto.Charger, error) {
	return dao.Db.GetAllCharger()
}

func (c *ChargerImpl) GetChargerById(chargerId uint) (dto.Charger, error) {
	return dao.Db.GetChargerById(chargerId)
}

func (c *ChargerImpl) DeleteCharger(chargerId uint) error {
	return dao.Db.DeleteChargerEntry(dto.Charger{ID: chargerId})
}

var (
	ChargerControllerObj ChargerController
)

func NewChargerController() {
	ChargerControllerObj = &ChargerImpl{}
}
