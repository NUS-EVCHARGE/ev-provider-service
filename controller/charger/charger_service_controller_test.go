package charger

import (
	"ev-provider-service/dao"
	"ev-provider-service/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() {
	NewChargerController()
}
func TestCreateChargerSuccess(t *testing.T) {
	setup()
	
	var (
		provider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
		actualCharger = dto.Charger{
			ID:         0,
			ProviderId: 0,
			RatesId:    0,
			Lat:        0,
			Lng:        0,
			Status:     "",
		}
	)
	
	dao.Db = dao.NewMockDatabase([]dto.Provider{}, []dto.Rates{}, []dto.Charger{})

	err := ChargerControllerObj.CreateCharger(actualCharger)
	assert.Nil(t, err)

	expectedCharger, err := ChargerControllerObj.GetCharger(provider.ID)
	assert.Nil(t, err)
	assert.Equal(t, actualCharger, expectedCharger[0])
}

func TestDeleteChargerSuccess(t *testing.T) {
	setup()
	var (
		provider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
		actualCharger = dto.Charger{
			ID:         0,
			ProviderId: 0,
			RatesId:    0,
			Lat:        0,
			Lng:        0,
			Status:     "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{}, []dto.Rates{}, []dto.Charger{actualCharger})

	err := ChargerControllerObj.DeleteCharger(actualCharger.ID)
	assert.Nil(t, err)

	chargerList, err := ChargerControllerObj.GetCharger(provider.ID)
	assert.Nil(t, err)
	assert.Equal(t, len(chargerList), 0)
}
