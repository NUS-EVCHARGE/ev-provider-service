package handler

import (
	"fmt"
	"net/http"

	google "github.com/NUS-EVCHARGE/ev-provider-service/third_party/google"
	"github.com/gin-gonic/gin"
)

var (
	GoogleClient = google.NewGoogleClient()
)

func RecommendationHandler(c *gin.Context) {
	input := c.Query("input")
	if input == "" {
		c.JSON(http.StatusBadRequest, CreateResponse("input is not defined"))
		return
	}
	res, err := GoogleClient.GetRecommendation(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
