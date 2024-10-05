package handler

import (
	"github.com/NUS-EVCHARGE/ev-provider-service/helper"
	"github.com/gin-gonic/gin"
)

const (
	defaultUrl = "https://jsx85ddz0a.execute-api.ap-southeast-1.amazonaws.com/evcharger-gateway/audit/api/v1/audit"
)

func PushAuditEvent(c *gin.Context) {
	providerObj, _ := c.Get("provider")
	action, _ := c.Get("action")
	accessToken := c.GetHeader("Authorization")

	helper.LaunchHttpRequest("POST", defaultUrl, map[string]string{}, map[string]interface{}{
		"Action":    action,
		"ChangedBy": providerObj,
	}, helper.WithBearerToken(accessToken))
	return
}
