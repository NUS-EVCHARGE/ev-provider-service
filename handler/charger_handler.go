package handler

import (
	"fmt"
	"github.com/NUS-EVCHARGE/ev-provider-service/config"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/charger"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/rates"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/NUS-EVCHARGE/ev-provider-service/helper"
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
//	@Param			authentication	header	string	yes		"jwtToken of the user"
//	@Param			provider_id		path	int		true	"Provider id"
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
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	providerId, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	Charger.ProviderId = uint(providerId)
	err = charger.ChargerControllerObj.CreateCharger(&Charger)
	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error creating Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, Charger)
	return
}

//	@Summary		Get Charger by provider
//	@Summary		Get Charger by charger id
//	@Description	get Charger by provider id or charger id
//	@Tags			Charger
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.Charger	"returns a list of Charger object"
//	@Router			/provider/{provider_id}/charger [get]
//	@Router			/provider/charger/{charger_id} [get]
//	@Param			authentication	header	string	yes		"jwtToken of the user"
//	@Param			provider_id		path	int		true	"Provider id"
//	@Param			charger_id		path	int		true	"Charger id"
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
	chargerId, err := strconv.Atoi(c.Param("charger_id"))

	if chargerId != 0 {
		chargerId, err := strconv.Atoi(c.Param("charger_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, CreateResponse("charger id must be an integer"))
		}
		chargerResult, err := charger.ChargerControllerObj.GetChargerById(uint(chargerId))
		chargerList = append(chargerList, chargerResult)
	} else if providerId != 0 {
		chargerList, err = charger.ChargerControllerObj.GetChargerByProvider(uint(providerId))
	}

	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, chargerList)
	return
}

//	@Summary		Get all Charger
//	@Summary		Get all Charger
//	@Description	get Charger by provider id or charger id
//	@Tags			Charger
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.Charger	"returns a list of Charger object"
//	@Router			/provider/{provider_id}/charger [get]
//	@Router			/provider/charger/{charger_id} [get]
//	@Param			authentication	header	string	yes		"jwtToken of the user"
//	@Param			provider_id		path	int		true	"Provider id"
//	@Param			charger_id		path	int		true	"Charger id"
func GetAllChargerHandler(c *gin.Context) {
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

	chargerList, err = charger.ChargerControllerObj.GetAllCharger()

	if err != nil {
		// todo: change to common library
		logrus.WithField("err", err).Error("error getting Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, chargerList)
	return
}

//	@Summary		Get Charger and Rate by provider id
//	@Description	get Charger and Rate by provider id
//	@Tags			ChargerRate
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.ChargerRate	"returns a list of Charger object"
//	@Router			/provider/{provider_id}/chargerandrate [get]
//	@Param			authentication	header	string	yes		"jwtToken of the user"
//	@Param			provider_id		path	int		true	"Provider id"
func GetChargerAndRateHandler(c *gin.Context) {
	var (
		chargerRateList []dto.ChargerRate
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

	if providerId != 0 {
		chargerList, err := charger.ChargerControllerObj.GetChargerByProvider(uint(providerId))
		if err != nil {
			c.JSON(http.StatusBadRequest, CreateResponse("provider id must be a valid integer"))
		}
		for _, charger := range chargerList {
			chargerRate, err := rates.RateControllerObj.GetRateByRateId(charger.RatesId)
			if err != nil {
				c.JSON(http.StatusBadRequest, CreateResponse("provider id must be a valid integer"))
			}
			item := dto.ChargerRate{
				ID:         charger.ID,
				ProviderId: charger.ProviderId,
				Address:    charger.Address,
				Lat:        charger.Lat,
				Lng:        charger.Lng,
				Status:     charger.Status,
				Rates:      chargerRate,
			}

			chargerRateList = append(chargerRateList, item)
		}
	}

	c.JSON(http.StatusOK, chargerRateList)
}

//	@summary		Create Charger and Rate by provider id
//	@description	Create Charger and Rate by provider id
//	@tags			ChargerRate
//	@accept			json
//	@produce		json
//	@success		200	{object}	dto.ChargerRate	"returns a ChargerRate object"
//	@router			/provider/{provider_id}/chargerandrate [post]
//	@param			authentication	header	string	yes	"jwtToken of the user"
func CreateChargerAndRateHandlerByProviderId(c *gin.Context) {
	var (
		chargerRate dto.ChargerRate
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

	err = c.BindJSON(&chargerRate)
	if err != nil {
		logrus.WithField("err", err).Error("error binding request")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	providerId, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	newRate := dto.Rates{
		ProviderId:        uint(providerId),
		NormalRate:        chargerRate.Rates.NormalRate,
		PenaltyRate:       chargerRate.Rates.PenaltyRate,
		NoShowPenaltyRate: chargerRate.Rates.NoShowPenaltyRate,
		Status:            chargerRate.Rates.Status,
	}

	err = rates.RateControllerObj.AddRate(&newRate)

	if err != nil {
		logrus.WithField("err", err).Error("error creating Rate")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	newCharger := dto.Charger{
		ProviderId: uint(providerId),
		RatesId:    newRate.ID,
		Address:    chargerRate.Address,
		Lat:        chargerRate.Lat,
		Lng:        chargerRate.Lng,
		Status:     chargerRate.Status,
	}

	err = charger.ChargerControllerObj.CreateCharger(&newCharger)

	if err != nil {
		logrus.WithField("err", err).Error("error creating Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	chargerRate.ID = newCharger.ID
	chargerRate.Rates.ID = newRate.ID

	c.JSON(http.StatusOK, chargerRate)
	return
}

//	@Summary		Update Charger and Rate by provider Id
//	@Description	Update Charger and Rate by provider Id
//	@Tags			ChargerRate
//	@Accept			json
//	@Produce		json
//	@success		200	{object}	dto.ChargerRate	"returns a ChargerRate object"
//	@router			/provider/{provider_id}/chargerandrate [patch]
//	@Param			authentication	header	string	yes		"jwtToken of the user"
//	@param			provider_id		path	int		true	"Provider id"
func UpdateChargerAndRateHandlerByProviderId(c *gin.Context) {
	var (
		chargerRate dto.ChargerRate
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

	err = c.BindJSON(&chargerRate)
	if err != nil {
		logrus.WithField("err", err).Error("error getting request")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	providerId, err := strconv.Atoi(c.Param("provider_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse("provider id must be an integer"))
	}

	newRate := dto.Rates{
		ID:                chargerRate.Rates.ID,
		ProviderId:        uint(providerId),
		NormalRate:        chargerRate.Rates.NormalRate,
		PenaltyRate:       chargerRate.Rates.PenaltyRate,
		NoShowPenaltyRate: chargerRate.Rates.NoShowPenaltyRate,
		Status:            chargerRate.Rates.Status,
	}

	err = rates.RateControllerObj.UpdateRate(newRate)
	if err != nil {
		logrus.WithField("err", err).Error("error updating Rate")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	newCharger := dto.Charger{
		ID:         chargerRate.ID,
		ProviderId: uint(providerId),
		RatesId:    newRate.ID,
		Address:    chargerRate.Address,
		Lat:        chargerRate.Lat,
		Lng:        chargerRate.Lng,
		Status:     chargerRate.Status,
	}

	err = charger.ChargerControllerObj.UpdateCharger(newCharger)
	if err != nil {
		logrus.WithField("err", err).Error("error updating Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, chargerRate)

}

//	@Summary		Update Charger by provider
//	@Description	Update Charger by provider
//	@Tags			Charger
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.Charger	"returns a Charger object"
//	@Router			/provider/{provider_id}/charger [patch]
//	@Param			authentication	header	string	yes		"jwtToken of the user"
//	@Param			provider_id		path	int		true	"Provider id"
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
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	err = charger.ChargerControllerObj.UpdateCharger(Charger)
	if err != nil {
		// todo: change to common library
		logrus.WithField("Charger", Charger).WithField("err", err).Error("error update Charger")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
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
//	@Param			provider_id		path	int		true	"Provider id"
//	@Param			charger_id		path	int		true	"Charger id"
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
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("Charger deletion success"))
	return
}
