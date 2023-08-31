package provider

import (
	"ev-provider-service/dao"
	"ev-provider-service/dto"
	"fmt"
	userDto "github.com/NUS-EVCHARGE/ev-user-service/dto"
	"github.com/stretchr/testify/assert"
	"testing"
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
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{}, []dto.Rates{}, []dto.Charger{})

	err := ProviderControllerObj.CreateProvider(actualProvider, user)
	assert.Nil(t, err)

	expectedProvider, err := ProviderControllerObj.GetProvider(user)
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
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{actualProvider}, []dto.Rates{}, []dto.Charger{})

	err := ProviderControllerObj.CreateProvider(actualProvider, user)
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
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{actualProvider}, []dto.Rates{}, []dto.Charger{})

	err := ProviderControllerObj.DeleteProvider(actualProvider.ID)
	assert.Nil(t, err)

	_, err = ProviderControllerObj.GetProvider(user)
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
		user = userDto.User{
			User:  "example",
			Email: "example@example.com",
		}
	)
	dao.Db = dao.NewMockDatabase([]dto.Provider{actualProvider}, []dto.Rates{}, []dto.Charger{})

	actualProvider.CompanyName = "example2"
	err := ProviderControllerObj.UpdateProvider(actualProvider)
	assert.Nil(t, err)

	expectedProvider, err := ProviderControllerObj.GetProvider(user)
	assert.Nil(t, err)
	assert.Equal(t, actualProvider, expectedProvider)
}
