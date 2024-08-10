package charger

import (
	"fmt"

	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	third_party "github.com/NUS-EVCHARGE/ev-provider-service/third_party/google"
)

type ChargerController interface {
	SearchChargerPoint(providerId int, placeId string) (dto.ChargerPoint, error)
	CreateChargerPoint(charger *dto.ChargerPoint) error
	UpdateChargerPoint(charger dto.ChargerPoint) error

	CreateCharger(charger dto.Charger) error
	UpdateCharger(charger dto.Charger) error

	GetAllCharger() ([]dto.ChargerFullDetails, error)
	GetAllChargerByCompanyName(companyName string) ([]dto.ChargerFullDetails, error)
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
		// todo: need to revamp ui where they click in to view more details?
		chargerFullDetailList = append(chargerFullDetailList, dto.ChargerFullDetails{
			Address:     chargerPoint.Address,
			ChargerList: chargerList,
		})
	}
	return chargerFullDetailList, nil
}

// get all charger by company name

func (c *ChargerImpl) GetAllChargerByCompanyName(companyName string) ([]dto.ChargerFullDetails, error) {
	var (
		chargerPointList      []dto.ChargerPoint
		chargerFullDetailList = []dto.ChargerFullDetails{}
		err                   error
	)

	/*
		Get Provider Id
	*/
	providerObj, err := dao.Db.GetProviderEntryByCompany(companyName)
	if err != nil {
		return chargerFullDetailList, err
	}

	/*
		Get all charger points
		for each charger points, get charger by charging point ids
	*/

	chargerPointList, err = dao.Db.GetAllChargerPointEntryByProviderID(int(providerObj.ID))
	if err != nil {
		return chargerFullDetailList, err
	}

	for chargerPointIndex, chargerPoint := range chargerPointList {
		chargerList, err := dao.Db.GetChargerByChargerPointId(chargerPoint.ID)
		if err != nil {
			return chargerFullDetailList, err
		}
		for chargerIndex, c := range chargerList {
			chargerFullDetailList = append(chargerFullDetailList, dto.ChargerFullDetails{
				Key:            fmt.Sprintf("%v_%v", chargerPointIndex, chargerIndex),
				Address:        chargerPoint.Address,
				UID:            c.UID,
				ChargerType:    c.ChargerType,
				Status:         c.Status,
				Details:        c.Details,
				Power:          c.Power,
				PowerType:      c.PowerType,
				ChargerPointID: c.ChargerPointID,
				ID:             c.ID,
				Lat:            chargerPoint.Lat,
				Lng:            chargerPoint.Lng,
			})
		}

		// chargerFullDetailList = append(chargerFullDetailList, dto.ChargerFullDetails{
		// 	Lat:         chargerPoint.Lat,
		// 	Lng:         chargerPoint.Lng,
		// 	Address:     chargerPoint.Address,
		// 	ChargerList: chargerList,
		// })

	}
	return chargerFullDetailList, nil
}

func (c *ChargerImpl) SearchChargerPoint(providerId int, placeId string) (dto.ChargerPoint, error) {
	return dao.Db.GetChargerPointByLocation(providerId, placeId)
}

// charging point
func (c *ChargerImpl) CreateChargerPoint(charger *dto.ChargerPoint) error {
	googleClient := third_party.NewGoogleClient()
	// get information and set lat and lng and address
	placeDetails, err := googleClient.GetPlaceDetails(charger.PlaceId)
	if err != nil {
		return err
	}
	charger.Lat = placeDetails.Geometry.Location.Lat
	charger.Lng = placeDetails.Geometry.Location.Lng
	charger.Address = placeDetails.FormattedAddress

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
