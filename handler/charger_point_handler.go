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
// @Router			/chargerpoint [post]
func CreateChargerPointHandler(c *gin.Context) {
	var (
		chargerPoint dto.ChargerPoint
	)

	err := c.BindJSON(&chargerPoint)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	err = charger.ChargerControllerObj.CreateChargerPoint(&chargerPoint)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, CreateResponse("success"))
	return
}

// @Summary		Update charging point by Id
// @Description	Update charging point by Id
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.Charger	"returns a Charger object"
// @Router			/charger/{charger_id} [patch]
func UpdateChargerPointHandler(c *gin.Context) {
	var (
		chargerRequest dto.ChargerPoint
	)

	err := c.BindJSON(&chargerRequest)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	err = charger.ChargerControllerObj.UpdateChargerPoint(chargerRequest)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Charger", chargerRequest).WithField("err", err).Error("error update Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("success"))
	return
}
