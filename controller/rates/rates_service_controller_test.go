package rates

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() {
	NewRateController()
}
func TestCreateRatesSuccess(t *testing.T) {
	setup()
	var (
		provider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
		actualRates = dto.Rates{
			ID:                0,
			ProviderId:        0,
			NormalRate:        2,
			PenaltyRate:       2,
			NoShowPenaltyRate: 2,
			Status:            "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{}, []dto.Rates{}, []dto.Charger{})
	err := RateControllerObj.AddRate(actualRates)
	assert.Nil(t, err)

	expectedRates, err := RateControllerObj.GetRateByProviderId(provider.ID)
	assert.Nil(t, err)
	assert.Equal(t, actualRates, expectedRates[0])
}

func TestDeleteRatesSuccess(t *testing.T) {
	setup()
	var (
		provider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
		actualRates = dto.Rates{
			ID:                0,
			ProviderId:        0,
			NormalRate:        2,
			PenaltyRate:       2,
			NoShowPenaltyRate: 2,
			Status:            "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{}, []dto.Rates{actualRates}, []dto.Charger{})

	err := RateControllerObj.DeleteRate(actualRates.ID)
	assert.Nil(t, err)

	ratesList, err := RateControllerObj.GetRateByProviderId(provider.ID)
	assert.Nil(t, err)
	assert.Equal(t, len(ratesList), 0)
}

func TestUpdateRatesSuccess(t *testing.T) {
	setup()
	var (
		provider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
		actualRates = dto.Rates{
			ID:                0,
			ProviderId:        0,
			NormalRate:        2,
			PenaltyRate:       2,
			NoShowPenaltyRate: 2,
			Status:            "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{}, []dto.Rates{actualRates}, []dto.Charger{})
	actualRates.NormalRate = 100
	err := RateControllerObj.UpdateRate(actualRates)
	assert.Nil(t, err)

	ratesList, err := RateControllerObj.GetRateByProviderId(provider.ID)
	assert.Nil(t, err)
	assert.Equal(t, actualRates, ratesList[0])
}
