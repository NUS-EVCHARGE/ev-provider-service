package provider

import (
	"fmt"
	"testing"

	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/stretchr/testify/assert"
)

func setup() {
	NewProviderController()
}

func TestCreateProviderSuccess(t *testing.T) {
	setup()
	var (
		actualProvider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{}, []dto.ChargerPoint{}, []dto.Charger{})

	_, err := ProviderControllerObj.CreateProvider(actualProvider)
	assert.Nil(t, err)

	expectedProvider, err := ProviderControllerObj.GetProvider(actualProvider.UserEmail)
	assert.Nil(t, err)
	assert.Equal(t, actualProvider, expectedProvider)
}

func TestCreateProviderIfExist(t *testing.T) {
	setup()
	var (
		actualProvider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{actualProvider}, []dto.ChargerPoint{}, []dto.Charger{})

	_, err := ProviderControllerObj.CreateProvider(actualProvider)
	assert.Equal(t, err, fmt.Errorf("provider already exist"))
}

func TestDeleteProviderSuccess(t *testing.T) {
	setup()
	var (
		actualProvider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{actualProvider}, []dto.ChargerPoint{}, []dto.Charger{})

	err := ProviderControllerObj.DeleteProvider(actualProvider.ID)
	assert.Nil(t, err)

	_, err = ProviderControllerObj.GetProvider(actualProvider.UserEmail)
	assert.Equal(t, err, fmt.Errorf("provider not found"))
}

func TestUpdateProviderSuccess(t *testing.T) {
	setup()
	var (
		actualProvider = dto.Provider{
			ID:          0,
			UserEmail:   "example@example.com",
			CompanyName: "example",
			Description: "example",
			Status:      "",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{actualProvider}, []dto.ChargerPoint{}, []dto.Charger{})

	actualProvider.CompanyName = "example2"
	err := ProviderControllerObj.UpdateProvider(actualProvider)
	assert.Nil(t, err)

	expectedProvider, err := ProviderControllerObj.GetProvider(actualProvider.UserEmail)
	assert.Nil(t, err)
	assert.Equal(t, actualProvider, expectedProvider)
}
