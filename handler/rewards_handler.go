package handler

import (
	"fmt"
	"net/http"

	controller "github.com/NUS-EVCHARGE/ev-provider-service/controller/rewards"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/gin-gonic/gin"
)

// coinpolicy handler
func UpdateCoinPolicy(c *gin.Context) {
	var req dto.CoinPolicy
	c.BindJSON(&req)

	providerObj, _ := c.Get("provider")
	req.ProviderId = int(providerObj.(dto.Provider).ID)

	if err := controller.RewardsControllerObj.UpdateCoinPolicy(req); err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("success"))
}

func GetCoinPolicy(c *gin.Context) {
	providerObj, _ := c.Get("provider")
	providerId := int(providerObj.(dto.Provider).ID)
	coinPolicy, err := controller.RewardsControllerObj.GetCoinPolicy(providerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, coinPolicy)
}

// voucher handler
func CreateVoucher(c *gin.Context) {
	var req dto.Vouchers

	c.BindJSON(&req)

	providerObj, _ := c.Get("provider")
	req.ProviderId = int(providerObj.(dto.Provider).ID)

	if err := controller.RewardsControllerObj.CreateVoucher(req); err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("success"))
}

func UpdateVoucher(c *gin.Context) {
	var req dto.Vouchers

	c.BindJSON(&req)

	providerObj, _ := c.Get("provider")
	req.ProviderId = int(providerObj.(dto.Provider).ID)

	if err := controller.RewardsControllerObj.UpdateVoucher(req); err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, CreateResponse("success"))
}

func GetVoucher(c *gin.Context) {
	providerObj, _ := c.Get("provider")
	providerId := int(providerObj.(dto.Provider).ID)
	voucherList, err := controller.RewardsControllerObj.GetAllVouchers(providerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, voucherList)
}
