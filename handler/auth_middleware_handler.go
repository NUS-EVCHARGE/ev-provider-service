package handler

import (
	"fmt"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/authentication"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func AuthMiddlewareHandler(c *gin.Context) {
	// Retrieve the access token from the Authorization header
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		logrus.Error("Authorization header is missing")
		c.JSON(http.StatusBadRequest, CreateResponse("Authorization header is missing"))
		c.Abort()
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(accessToken, bearerPrefix) {
		logrus.Error("Invalid authorization header format")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
		c.Abort()
		return
	}

	// Optional: remove "Bearer " prefix if it exists
	if len(accessToken) > len(bearerPrefix) && accessToken[:len(bearerPrefix)] == bearerPrefix {
		accessToken = accessToken[len(bearerPrefix):]
	}

	err := authentication.AuthenticationControllerObj.GetUserInfo(accessToken)
	if err != nil {
		logrus.WithField("err", err).Error("error validating token")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, CreateResponse("Token is valid"))
	return
}
