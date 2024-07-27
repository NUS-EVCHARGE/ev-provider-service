package charger

import (
	"testing"

	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/stretchr/testify/assert"
)

var (
	mockProvider = dto.Provider{
		ID:          0,
		CompanyName: "testing",
	}
	mockChargerPoint = dto.ChargerPoint{
		ID:         0,
		Address:    "test address",
		ProviderId: 0,
	}
	mockChargerList = []dto.Charger{
		{
			ID:      0,
			UID:     "wleihflwe",
			Status:  "available",
			Details: `{"charger_type":"AC","rates":"$0.123"}`,
		},
		{
			ID:      1,
			UID:     "ewleihflwe",
			Status:  "available",
			Details: `{"charger_type":"DC","rates":"$0.123"}`,
		},
	}
)

func setup() {
	NewChargerController()
	dao.Db = dao.NewMockDatabase([]dto.Provider{
		mockProvider,
	}, []*dto.ChargerPoint{
		&mockChargerPoint,
	}, mockChargerList)
}

func TestCreateChargingPointSucess(t *testing.T) {
	setup()
	newChargingPoint := dto.ChargerPoint{
		ID:         1,
		Address:    "new address",
		ProviderId: 0,
	}
	expectedChargerDetails := []dto.ChargerFullDetails{
		{
			Address:     mockChargerPoint.Address,
			ChargerList: mockChargerList,
		},
		{
			Address: "new address",
		},
	}
	err := ChargerControllerObj.CreateChargerPoint(&newChargingPoint)
	assert.Nil(t, err)
	chargerDetails, err := ChargerControllerObj.GetAllCharger()
	assert.Nil(t, err)
	assert.Equal(t, chargerDetails, expectedChargerDetails)
}

func TestCreateChargerSucess(t *testing.T) {
	setup()
	newCharger := dto.Charger{
		ID:      2,
		UID:     "eewfwefwleihflwe",
		Status:  "available",
		Details: `{"charger_type":"DC","rates":"$0.123"}`,
	}
	expectedChargerDetails := []dto.ChargerFullDetails{
		{
			Address:     mockChargerPoint.Address,
			ChargerList: append(mockChargerList, newCharger),
		},
	}
	err := ChargerControllerObj.CreateCharger(newCharger)
	assert.Nil(t, err)
	chargerDetails, err := ChargerControllerObj.GetAllCharger()
	assert.Nil(t, err)
	assert.Equal(t, chargerDetails, expectedChargerDetails)
}

func TestUpdateChargerSucess(t *testing.T) {
	setup()
	newCharger := dto.Charger{
		ID:      1,
		UID:     "eewfwefwleihflwe",
		Status:  "available",
		Details: `{"charger_type":"DC","rates":"$0.123"}`,
	}
	newChargerList := mockChargerList
	newChargerList[1] = newCharger

	expectedChargerDetails := []dto.ChargerFullDetails{
		{
			Address:     mockChargerPoint.Address,
			ChargerList: newChargerList,
		},
	}

	err := ChargerControllerObj.UpdateCharger(newCharger)
	assert.Nil(t, err)
	chargerDetails, err := ChargerControllerObj.GetAllCharger()
	assert.Nil(t, err)
	assert.Equal(t, chargerDetails, expectedChargerDetails)
}

func TestUpdateChargerPointSucess(t *testing.T) {
	setup()
	newChargingPoint := dto.ChargerPoint{
		ID:         0,
		Address:    "new address",
		ProviderId: 0,
	}

	expectedChargerDetails := []dto.ChargerFullDetails{
		{
			Address:     "new address",
			ChargerList: mockChargerList,
		},
	}

	err := ChargerControllerObj.UpdateChargerPoint(newChargingPoint)
	assert.Nil(t, err)
	chargerDetails, err := ChargerControllerObj.GetAllCharger()
	assert.Nil(t, err)
	assert.Equal(t, chargerDetails, expectedChargerDetails)
}
