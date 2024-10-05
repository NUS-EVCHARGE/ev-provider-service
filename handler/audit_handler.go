package handler

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/helper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	defaultUrl = "https://jsx85ddz0a.execute-api.ap-southeast-1.amazonaws.com/evcharger-gateway/audit/api/v1/audit"
)

func PushAuditEvent(c *gin.Context) {
	providerObj, _ := c.Get("provider")
	action, _ := c.Get("action")
	description, _ := c.Get("description")
	// var descriptionInt map[string]interface{}
	// json.Unmarshal([]byte(description.(string)), &descriptionInt)
	accessToken := c.GetHeader("Authorization")

	res, err := helper.LaunchHttpRequest("POST", defaultUrl, map[string]string{}, map[string]interface{}{
		"Action":      action,
		"ChangedBy":   providerObj,
		"Description": description,
	}, helper.WithBearerToken(accessToken))
	var resp map[string]interface{}
	helper.HttpResponseDefaultParser(res, resp)
	logrus.WithField("res", res).WithField("err", err).Info("audit_service_status")
	return
}
