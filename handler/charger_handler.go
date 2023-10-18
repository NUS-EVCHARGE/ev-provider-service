package handler

import (
	"fmt"
	"github.com/NUS-EVCHARGE/ev-provider-service/config"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/charger"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/NUS-EVCHARGE/ev-provider-service/helper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary		Create Charger by provider
// @Description	create Charger by provider
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Charger	"returns a Charger object"
// @Router			/provider/{provider_id}/charger [post]
// @Param			authentication	header	string	yes	"jwtToken of the user"
// @Param			provider_id				path	int		true	"Provider id"
func CreateChargerHandler(c *gin.Context) {
	var (
		Charger dto.Charger
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

	err = c.BindJSON(&Charger)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v",err)))
		return
	}

	providerId, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	Charger.ProviderId = uint(providerId)
	err = charger.ChargerControllerObj.CreateCharger(Charger)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v",err)))
		return
	}
	c.JSON(http.StatusOK, Charger)
	return
}

// @Summary		Get Charger by provider
// @Description	get Charger by provider
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]dto.Charger	"returns a list of Charger object"
// @Router			/provider/{provider_id}/charger [get]
// @Param			authentication	header	string	yes	"jwtToken of the user"
// @Param			provider_id				path	int		true	"Provider id"
func GetChargerHandler(c *gin.Context) {
	var (
		chargerList []dto.Charger
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
	providerId, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		// return all chargers as provider id cannot be parsed
		chargerList, err = charger.ChargerControllerObj.GetAllCharger()
	} else {
		chargerList, err = charger.ChargerControllerObj.GetChargerByProvider(uint(providerId))
	}

	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v",err)))
		return
	}
	c.JSON(http.StatusOK, chargerList)
	return
}

// @Summary		Update Charger by provider
// @Description	Update Charger by provider
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Charger	"returns a Charger object"
// @Router			/provider/{provider_id}/charger [patch]
// @Param			authentication	header	string	yes	"jwtToken of the user"
// @Param			provider_id				path	int		true	"Provider id"
func UpdateChargerHandler(c *gin.Context) {
	var (
		Charger dto.Charger
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
	err = c.BindJSON(&Charger)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v",err)))
		return
	}
	err = charger.ChargerControllerObj.UpdateCharger(Charger)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Charger", Charger).WithField("err", err).Error("error update Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v",err)))
		return
	}
	c.JSON(http.StatusOK, Charger)
	return
}

// @Summary		Delete Charger by charger id
// @Description	Delete Charger by charger id
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Charger	"returns a Charger object"
// @Router			/provider/{provider_id}/charger/{charger_id} [delete]
// @Param			authentication	header	string	yes		"jwtToken of the user"
// @Param			provider_id				path	int		true	"Provider id"
// @Param			charger_id				path	int		true	"Charger id"
func DeleteChargerHandler(c *gin.Context) {
	var (
		Charger dto.Charger
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

	id, err := strconv.Atoi(c.Param("charger_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("id but be an integer"))
	}
	err = charger.ChargerControllerObj.DeleteCharger(uint(id))
	if err != nil {
		// todo: change to common library
		logrus.WithField("Charger", Charger).WithField("err", err).Error("error update Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v",err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("Charger deletion success"))
	return
}
