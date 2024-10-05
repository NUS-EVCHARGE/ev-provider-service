package handler

import (
	"fmt"
	"net/http"

	"github.com/NUS-EVCHARGE/ev-provider-service/controller/charger"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateChargerRequest struct {
	Email              string  `json:"email"`
	PlaceId            string  `json:"place_id"`
	Address            string  `gorm:"colummn:address" json:"address"`
	ProviderName       string  `gorm:"provider_name" json:"provider_name"`
	UID                string  `gorm:"column:uid" json:"uid"`
	EVRegistrationMark string  `json:"ev_registration_mark"`
	Details            string  `gorm:"column:details" json:"details"` // json struct
	PowerType          string  `gorm:"column:power_type" json:"power_type"`
	ChargerType        string  `gorm:"column:charger_type" json:"charger_type"`
	Power              float64 `gorm:"colum:power" json:"power"`
}

// @Summary		Create Charger by provider
// @Description	create Charger by provider
// @Tags			Charger
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.ChargerReq	"returns a sucess"
// @Router			/charger [post]
func CreateChargerHandler(c *gin.Context) {
	var (
		chargerReq CreateChargerRequest
	)

	err := c.BindJSON(&chargerReq)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		c.Abort()
		return
	}
	// get company id
	providerObj, _ := c.Get("provider")

	// check for charger point
	chargerPoint, err := charger.ChargerControllerObj.SearchChargerPoint(int(providerObj.(dto.Provider).ID), chargerReq.PlaceId)
	if err != nil {
		logrus.WithField("err", err).Error("get charger point error")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		c.Abort()
		return
	}

	// there is no charger point
	if chargerPoint.ID == 0 {
		res, err := GoogleClient.GetPlaceDetails(chargerReq.PlaceId)
		if err != nil {
			logrus.WithField("err", err).Error("get charger point details error from google")
			c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
			c.Abort()
			return
		}

		chargerPoint = dto.ChargerPoint{
			ProviderId: providerObj.(dto.Provider).ID,
			PlaceId:    chargerReq.PlaceId,
			Lat:        res.Geometry.Location.Lat,
			Lng:        res.Geometry.Location.Lng,
		}
		err = charger.ChargerControllerObj.CreateChargerPoint(&chargerPoint)
		if err != nil {
			logrus.WithField("err", err).Error("error creating charger point")
			c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
			c.Abort()
			return
		}
	}

	chargerObj := dto.Charger{
		ChargerPointID: chargerPoint.ID,
		UID:            chargerReq.UID,
		Details:        chargerReq.Details,
		PowerType:      chargerReq.PowerType,
		ChargerType:    chargerReq.ChargerType,
		Power:          chargerReq.Power,
	}

	err = charger.ChargerControllerObj.CreateCharger(chargerObj)

	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.Set("action", "create_charger")
	c.Set("description", fmt.Sprintf("create req: %v", chargerObj))
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
	providerObj, _ := c.Get("provider")
	companeName := providerObj.(dto.Provider).CompanyName

	if companeName != "" {
		chargerResult, err := charger.ChargerControllerObj.GetAllChargerByCompanyName(companeName)
		if err != nil {
			// todo: change to common library
			logrus.WithField("err", err).Error("error getting all charger")
			c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
			return
		}

		c.JSON(http.StatusOK, chargerResult)
		return
	}
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
		c.Abort()
		return
	}
	err = charger.ChargerControllerObj.UpdateCharger(chargerRequest)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Charger", chargerRequest).WithField("err", err).Error("error update Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		c.Abort()
		return
	}
	c.Set("action", "update_charger")
	c.Set("description", fmt.Sprintf("create req: %v", chargerRequest))
	c.JSON(http.StatusOK, CreateResponse("success"))
	return
}
