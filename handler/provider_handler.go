package handler

import (
	"ev-provider-service/config"
	"ev-provider-service/controller/provider"
	"ev-provider-service/dto"
	"ev-provider-service/helper"
	"fmt"
	userDto "github.com/NUS-EVCHARGE/ev-user-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
// @Param			authentication	header	string	yes	"jwtToken of the user"
func CreateProviderHandler(c *gin.Context) {
	var (
		user     userDto.User
		Provider dto.Provider
	)
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	user, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	err = c.BindJSON(&Provider)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	Provider.UserEmail = user.Email
	Provider, err = provider.ProviderControllerObj.CreateProvider(Provider, user)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating Provider")
		c.JSON(http.StatusBadRequest, err)
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
// @Router			/provider [get]
// @Param			authentication	header	string	yes	"jwtToken of the user"
func GetProviderHandler(c *gin.Context) {
	var (
		user     userDto.User
		Provider dto.Provider
	)
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	user, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	Provider, err = provider.ProviderControllerObj.GetProvider(user)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting Provider")
		c.JSON(http.StatusBadRequest, err)
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
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	user, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	err = c.BindJSON(&Provider)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	Provider.UserEmail = user.Email
	err = provider.ProviderControllerObj.UpdateProvider(Provider)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Provider", Provider).WithField("err", err).Error("error update Provider")
		c.JSON(http.StatusBadRequest, err)
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
// @Param			provider_id				path	int		true	"Provider id"
func DeleteProviderHandler(c *gin.Context) {
	var (
		Provider dto.Provider
	)
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	_, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	id, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("id but be an integer"))
	}
	err = provider.ProviderControllerObj.DeleteProvider(uint(id))
	if err != nil {
		// todo: change to common library
		logrus.WithField("Provider", Provider).WithField("err", err).Error("error update Provider")
		c.JSON(http.StatusBadRequest, err)
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
