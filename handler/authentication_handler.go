package handler

import (
	"fmt"
	"net/http"

	"github.com/NUS-EVCHARGE/ev-provider-service/controller/authentication"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoginHandler(c *gin.Context) {
	var (
		credentials dto.Credentials
	)

	err := c.BindJSON(&credentials)
	if err != nil {
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	resp, err := authentication.AuthenticationControllerObj.LoginUser(credentials)
	if err != nil {
		logrus.WithField("err", err).Error("error authenticating user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func SignUpHandler(c *gin.Context) {
	var (
		credentials dto.Credentials
	)

	err := c.BindJSON(&credentials)
	if err != nil {
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	err = authentication.AuthenticationControllerObj.RegisterUser(credentials)
	if err != nil {
		logrus.WithField("err", err).Error("error registering user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("Proceed to confirm user"))
	return
}

func ConfirmUserHandler(c *gin.Context) {
	var (
		userInfo dto.ConfirmUser
	)

	err := c.BindJSON(&userInfo)
	if err != nil {
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	err = authentication.AuthenticationControllerObj.ConfirmUser(userInfo)
	if err != nil {
		logrus.WithField("err", err).Error("error confirming user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("User confirmed successfully"))
	return
}

func ResendChallengeCodeHandler(c *gin.Context) {
	var (
		resendRequest dto.SignUpResendRequest
	)

	err := c.BindJSON(&resendRequest)
	if err != nil {
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	err = authentication.AuthenticationControllerObj.ResendChallengeCode(resendRequest)
	if err != nil {
		logrus.WithField("err", err).Error("error resending confirmation")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("Confirmation code sent successfully"))
	return
}
