package handler

import (
	"fmt"
	"net/http"

	controller "github.com/NUS-EVCHARGE/ev-provider-service/controller/rewards"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/gin-gonic/gin"
)

func CreateVoucher(c *gin.Context) {
	var req dto.Vouchers

	c.BindJSON(&req)
	if err := controller.RewardsControllerObj.CreateVoucher(req); err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("success"))
}
