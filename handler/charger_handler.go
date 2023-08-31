package handler

import (
	"ev-provider-service/config"
	"ev-provider-service/controller/charger"
	"ev-provider-service/dto"
	"ev-provider-service/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

//	@Summary		Create Charger by provider
//	@Description	create Charger by provider
//	@Tags			Charger
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Charger	"returns a Charger object"
//	@Router			/provider/{provider_id}/charger [post]
//	@Param			authentication	header	string	yes	"jwtToken of the user"
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
		c.JSON(http.StatusBadRequest, err)
		return
	}

	providerId , err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	Charger.ProviderId = uint(providerId)
	err = charger.ChargerControllerObj.CreateCharger(Charger)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating Charger")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, Charger)
	return
}

//	@Summary		Get Charger by provider
//	@Description	get Charger by provider
//	@Tags			Charger
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.Charger	"returns a list of Charger object"
//	@Router			/provider/{provider_id}/charger [get]
//	@Param			authentication	header	string	yes	"jwtToken of the user"
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
	providerId , err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	chargerList, err = charger.ChargerControllerObj.GetCharger(uint(providerId))
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting Charger")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, chargerList)
	return
}

//	@Summary		Update Charger by provider
//	@Description	Update Charger by provider
//	@Tags			Charger
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Charger	"returns a Charger object"
//	@Router			/Charger [patch]
//	@Param			authentication	header	string	yes	"jwtToken of the user"
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
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = charger.ChargerControllerObj.UpdateCharger(Charger)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Charger", Charger).WithField("err", err).Error("error update Charger")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, Charger)
	return
}

//	@Summary		Delete Charger by charger id
//	@Description	Delete Charger by charger id
//	@Tags			Charger
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Charger	"returns a Charger object"
//	@Router			/provider/{provider_id}/charger/{charger_id} [delete]
//	@Param			authentication	header	string	yes		"jwtToken of the user"
//	@Param			id				path	int		true	"Charger id"
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
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, CreateResponse("Charger deletion success"))
	return
}
