package handler

import (
	"ev-provider-service/config"
	"ev-provider-service/controller/rates"
	"ev-provider-service/dto"
	"ev-provider-service/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

//	@Summary		Create Rates by user
//	@Description	create Provider by user
//	@Tags			Rates
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Rates	"returns a Provider object"
//	@Router			/provider/${provider_id}/rates [post]
//	@Param			authentication	header	string	yes	"jwtToken of the user"
//	@Param			provider_id				path	int		true	"provider id"
func CreateRatesHandler(c *gin.Context) {
	var (
		ratesObj dto.Rates
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

	err = c.BindJSON(&ratesObj)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	providerIdInt, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	ratesObj.ProviderId = uint(providerIdInt)
	err = rates.RateControllerObj.AddRate(ratesObj)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error adding rates")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, ratesObj)
	return
}

//	@Summary		Get Rates by Provider
//	@Description	get Rates by Provider
//	@Tags			Rates
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.Rates	"returns a []dot.Rates object"
//	@Router			/provider/${provider_id}/rates [get]
//	@Param			authentication	header	string	yes	"jwtToken of the user"
//	@Param			provider_id				path	int		true	"provider id"
func GetRatesHandler(c *gin.Context) {
	tokenStr := c.GetHeader("Authentication")

	// Get User information
	_, err := helper.GetUser(config.GetUserUrl, tokenStr)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	providerIdInt, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	ratesList, err := rates.RateControllerObj.GetRateByProviderId(uint(providerIdInt))

	c.JSON(http.StatusOK, ratesList)
	return
}

//	@Summary		update rates by provider
//	@Description	update rates by provider
//	@Tags			Rates
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Rates	"returns a Rates object"
//	@Router			/provider/${provider_id}/rates [patch]
//	@Param			authentication	header	string	yes	"jwtToken of the user"
//	@Param			provider_id				path	int		true	"provider id"
func UpdateRatesHandler(c *gin.Context) {
	var (
		Rates dto.Rates
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
	providerIdInt, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	err = c.BindJSON(&Rates)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	Rates.ProviderId = uint(providerIdInt)
	err = rates.RateControllerObj.UpdateRate(Rates)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Rates", Rates).WithField("err", err).Error("error update rates")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, Rates)
	return
}

//	@Summary		delete rates by rates id
//	@Description	delete rates by rates id
//	@Tags			Rates
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Rates	"returns true/false"
//	@Router			/provider/{provider_id}/rates/{rates_id} [delete]
//	@Param			authentication	header	string	yes		"jwtToken of the user"
//	@Param			id				path	int		true	"Provider id"
//	@Param			provider_id				path	int		true	"provider id"
//	@Param			rates_id				path	int		true	"rates id"
func DeleteRatesHandler(c *gin.Context) {
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

	ratesId, err := strconv.Atoi(c.Param("rates)d"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("id but be an integer"))
	}
	err = rates.RateControllerObj.DeleteRate(uint(ratesId))
	if err != nil {
		// todo: change to common library
		logrus.WithField("Provider", Provider).WithField("err", err).Error("error update Provider")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, CreateResponse("Provider deletion success"))
	return
}
