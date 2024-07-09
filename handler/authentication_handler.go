package handler

import (
	"fmt"
	"net/http"

	"github.com/NUS-EVCHARGE/ev-provider-service/controller/authentication"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/provider"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SignUpRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CompanyName string `json:"company_name"`
}

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
	// get compnay name
	providerDetails, err := provider.ProviderControllerObj.GetProvider(credentials.Email)
	if err != nil {
		logrus.WithField("err", err).Error("error getting provider")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	resp.CompanyName = providerDetails.CompanyName

	c.JSON(http.StatusOK, resp)
	return
}

func SignUpHandler(c *gin.Context) {
	var (
		signUpRequest   SignUpRequest
		credentials     dto.Credentials
		providerDetails dto.Provider
	)

	err := c.BindJSON(&signUpRequest)
	if err != nil {
		logrus.WithField("err", err).Error("error params")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	// check if company name already exist
	if provider.ProviderControllerObj.IsCompanyExist(signUpRequest.CompanyName) {
		logrus.Error("error company already exist")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", "error company already exist please contact adminisrator")))
		return
	}

	credentials = dto.Credentials{
		Email:    signUpRequest.Email,
		Password: signUpRequest.Password,
	}

	err = authentication.AuthenticationControllerObj.RegisterUser(credentials)
	if err != nil {
		logrus.WithField("err", err).Error("error registering user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}

	providerDetails = dto.Provider{
		UserEmail:   signUpRequest.Email,
		CompanyName: signUpRequest.CompanyName,
		Status:      "1",
	}

	// if successful, write to db
	_, err = provider.ProviderControllerObj.CreateProvider(providerDetails)
	if err != nil {
		logrus.WithField("err", err).Error("error registering user when inserting to db")
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

func LogoutUserHandler(c *gin.Context) {
	// Retrieve the access token from the Authorization header
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		logrus.Error("Authorization header is missing")
		c.JSON(http.StatusBadRequest, CreateResponse("Authorization header is missing"))
		return
	}

	// Optional: remove "Bearer " prefix if it exists
	const bearerPrefix = "Bearer "
	if len(accessToken) > len(bearerPrefix) && accessToken[:len(bearerPrefix)] == bearerPrefix {
		accessToken = accessToken[len(bearerPrefix):]
	}

	err := authentication.AuthenticationControllerObj.LogoutUser(accessToken)
	if err != nil {
		logrus.WithField("err", err).Error("error logging out user")
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("User logged out successfully"))
	return
}
