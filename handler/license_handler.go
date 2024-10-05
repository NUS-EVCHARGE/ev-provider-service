package handler

import (
	"fmt"
	"net/http"

	"github.com/NUS-EVCHARGE/ev-provider-service/controller/license"
	"github.com/NUS-EVCHARGE/ev-provider-service/dto"
	"github.com/gin-gonic/gin"
)

func GetLicenseByCompanyHandler(c *gin.Context) {
	providerObj, _ := c.Get("provider")

	companyName := providerObj.(dto.Provider).CompanyName

	if companyName == "" {
		c.JSON(http.StatusBadRequest, CreateResponse("company_name undefined"))
		return
	}
	licenseObj, err := license.LicenseControllerObj.GetLicenseByCompany(companyName)
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		return
	}
	c.JSON(http.StatusOK, licenseObj)
	return
}

func UpdateLicenseByCompanyHandler(c *gin.Context) {
	var licenseObj dto.License
	companyName := c.Query("company_name")
	if companyName == "" {
		c.JSON(http.StatusBadRequest, CreateResponse("company_name undefined"))
		c.Abort()
		return
	}
	c.BindJSON(&licenseObj)
	if err := license.LicenseControllerObj.UpdateLicense(licenseObj, companyName); err != nil {
		c.JSON(http.StatusBadRequest, CreateResponse(fmt.Sprintf("%v", err)))
		c.Abort()
		return
	}
	c.Set("action", "update_license")
	c.Set("description", licenseObj)
	c.JSON(http.StatusOK, CreateResponse("success"))
	return
}
