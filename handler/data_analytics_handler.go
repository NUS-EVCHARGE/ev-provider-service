package handler

import (
	"net/http"
	"os"

	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func ChangerLocationHeatmapHandler(c *gin.Context) {
	clientsFile, err := os.OpenFile("handler/Singapore_EV_Charger_Recommendations_-_Top_200.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer clientsFile.Close()

	data := []*dto.HeatMapData{}
	if err := gocsv.UnmarshalFile(clientsFile, &data); err != nil { // Load clients from file
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}
