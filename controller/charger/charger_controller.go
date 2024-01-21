package charger

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
)

type ChargerController interface {
	CreateChargerPoint(charger dto.ChargerPoint) error
	UpdateChargerPoint(charger dto.ChargerPoint) error

	CreateCharger(charger dto.Charger) error
	UpdateCharger(charger dto.Charger) error

	GetAllCharger() ([]dto.ChargerFullDetails, error)
}

type ChargerImpl struct {
}

// all
func (c *ChargerImpl) GetAllCharger() ([]dto.ChargerFullDetails, error) {
	var (
		chargerPointList      []dto.ChargerPoint
		chargerFullDetailList = []dto.ChargerFullDetails{}
		err                   error
	)
	/*
		Get all charger points
		for each charger points, get charger by charging point ids
	*/

	chargerPointList, err = dao.Db.GetAllChargerPointEntry()
	if err != nil {
		return chargerFullDetailList, err
	}

	for _, chargerPoint := range chargerPointList {
		chargerList, err := dao.Db.GetChargerByChargerPointId(chargerPoint.ID)
		if err != nil {
			return chargerFullDetailList, err
		}
		chargerFullDetailList = append(chargerFullDetailList, dto.ChargerFullDetails{
			ProviderName: chargerPoint.ProviderName,
			Lat:          chargerPoint.Lat,
			Lng:          chargerPoint.Lng,
			Address:      chargerPoint.Address,
			ChargerList:  chargerList,
			Status:       chargerPoint.Status,
		})
	}
	return chargerFullDetailList, nil
}

// charging point
func (c *ChargerImpl) CreateChargerPoint(charger dto.ChargerPoint) error {
	return dao.Db.CreateChargerPointEntry(charger)
}

func (c *ChargerImpl) UpdateChargerPoint(charger dto.ChargerPoint) error {
	return dao.Db.UpdateChargerPointEntry(charger)
}

// charger
func (c *ChargerImpl) CreateCharger(charger dto.Charger) error {
	return dao.Db.CreateChargerEntry(charger)
}

func (c *ChargerImpl) UpdateCharger(charger dto.Charger) error {
	return dao.Db.UpdateChargerEntry(charger)
}

var (
	ChargerControllerObj ChargerController
)

func NewChargerController() {
	ChargerControllerObj = &ChargerImpl{}
}
