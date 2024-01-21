package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NUS-EVCHARGE/ev-provider-service/controller/provider"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	userDto "github.com/NUS-EVCHARGE/ev-user-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetBookingHealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, CreateResponse("Welcome to ev-providers-service"))
	return
}

// @Summary		Create Provider by user
// @Description	create Provider by user
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Provider	"returns a Provider object"
// @Router			/provider [post]
func CreateProviderHandler(c *gin.Context) {
	var (
		Provider dto.Provider
	)

	err := c.BindJSON(&Provider)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	// TODO: status management
	Provider.Status = "available"
	Provider, err = provider.ProviderControllerObj.CreateProvider(Provider)

	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating Provider")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, Provider)
	return
}

// @Summary		Get Provider by user
// @Description	get Provider by user
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Provider	"returns a Provider object"
// @Router			/provider/{provider_email} [get]
func GetProviderHandler(c *gin.Context) {
	var (
		Provider dto.Provider
	)
	providerEmail := c.Param("provider_email")
	if providerEmail == "" {
		providerList, err := provider.ProviderControllerObj.GetAllProvider()
		if err != nil {
			// todo: change to common library
			logrus.WithField("err", err).Error("error getting Provider")
			c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
			return
		}
		c.JSON(http.StatusOK, providerList)
		return
	}
	Provider, err := provider.ProviderControllerObj.GetProvider(providerEmail)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting Provider")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, Provider)
	return
}

// @Summary		Create Provider by user
// @Description	create Provider by user
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Provider	"returns a Provider object"
// @Router			/provider [patch]
// @Param			authentication	header	string	yes	"jwtToken of the user"
func UpdateProviderHandler(c *gin.Context) {
	var (
		user     userDto.User
		Provider dto.Provider
	)

	err := c.BindJSON(&Provider)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	Provider.UserEmail = user.Email
	err = provider.ProviderControllerObj.UpdateProvider(Provider)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Provider", Provider).WithField("err", err).Error("error update Provider")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, Provider)
	return
}

// @Summary		Create Provider by user
// @Description	create Provider by user
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Provider	"returns a Provider object"
// @Router			/provider/{provider_id} [delete]
// @Param			authentication	header	string	yes		"jwtToken of the user"
// @Param			provider_id		path	int		true	"Provider id"
func DeleteProviderHandler(c *gin.Context) {
	var (
		Provider dto.Provider
	)

	id, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("id but be an integer"))
	}
	err = provider.ProviderControllerObj.DeleteProvider(uint(id))
	if err != nil {
		// todo: change to common library
		logrus.WithField("Provider", Provider).WithField("err", err).Error("error update Provider")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("Provider deletion success"))
	return
}

func CreateResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
	}
}
