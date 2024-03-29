package handler

import (
	"fmt"
	"net/http"

	"github.com/NUS-EVCHARGE/ev-provider-service/controller/charger"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary		Create Charger by provider
// @Description	create Charger by provider
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Charger	"returns a sucess"
// @Router			/charger [post]
func CreateChargerHandler(c *gin.Context) {
	var (
		chargerPoint dto.Charger
	)

	err := c.BindJSON(&chargerPoint)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	err = charger.ChargerControllerObj.CreateCharger(chargerPoint)

	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, CreateResponse("success"))
	return
}

// @Summary		Get All Charger
// @Description	get Charger by provider id or charger id
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]dto.ChargerDetails	"returns a list of Charger Details object"
// @Router			/charger [get]
func GetAllChargerDetailsHandler(c *gin.Context) {
	chargerResult, err := charger.ChargerControllerObj.GetAllCharger()
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting all charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, chargerResult)
	return
}

// @Summary		Update Charger by id
// @Description	Update Charger by provider
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Charger	"returns a Charger object"
// @Router			/charger [patch]
func UpdateChargerHandler(c *gin.Context) {
	var (
		chargerRequest dto.Charger
	)

	err := c.BindJSON(&chargerRequest)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	err = charger.ChargerControllerObj.UpdateCharger(chargerRequest)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Charger", chargerRequest).WithField("err", err).Error("error update Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("success"))
	return
}
