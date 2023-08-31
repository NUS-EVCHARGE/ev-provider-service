package main

import (
	"ev-provider-service/config"
	"ev-provider-service/controller/charger"
	"ev-provider-service/controller/provider"
	Rate "ev-provider-service/controller/rates"
	"ev-provider-service/dao"
	_ "ev-provider-service/docs"
	"ev-provider-service/handler"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	r *gin.Engine
)

func main() {
	var (
		configFile string
	)
	flag.StringVar(&configFile, "config", "config.yaml", "configuration file of this service")
	flag.Parse()

	// init configurations
	configObj, err := config.ParseConfig(configFile)
	if err != nil {
		logrus.WithField("error", err).WithField("filename", configFile).Error("failed to init configurations")
		return
	}

	// init db
	err = dao.InitDB(configObj.Dsn)
	if err != nil {
		logrus.WithField("config", configObj).Error("failed to connect to database")
		//return
	}

	provider.NewProviderController()
	charger.NewChargerController()
	Rate.NewRateController()
	InitHttpServer(configObj.HttpAddress)
}

func InitHttpServer(httpAddress string) {
	r = gin.Default()
	registerHandler()

	if err := r.Run(httpAddress); err != nil {
		logrus.WithField("error", err).Errorf("http server failed to start")
	}
}

func registerHandler() {
	// use to generate swagger ui
	//	@BasePath	/api/v1
	//	@title		Provider Service API
	//	@version	1.0
	//	@schemes	http
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// api versioning
	v1 := r.Group("/api/v1")
	providerGroup := v1.Group("provider")
	providerGroup.POST("/", handler.CreateProviderHandler)
	providerGroup.GET("/", handler.GetProviderHandler)
	providerGroup.PATCH("/", handler.UpdateProviderHandler)
	providerGroup.DELETE("/:provider_id", handler.DeleteProviderHandler)
	
	providerGroup.POST("/:provider_id/charger", handler.CreateChargerHandler)
	providerGroup.GET("/:provider_id/charger", handler.GetChargerHandler)
	providerGroup.PATCH("/:provider_id/charger", handler.UpdateChargerHandler)
	providerGroup.DELETE("/:provider_id/charger/:charger_id", handler.DeleteChargerHandler)

	providerGroup.POST("/:provider_id/rates", handler.CreateRatesHandler)
	providerGroup.GET("/:provider_id/rates", handler.GetRatesHandler)
	providerGroup.PATCH("/:provider_id/rates", handler.UpdateRatesHandler)
	providerGroup.DELETE("/:provider_id/rates/:rates_id", handler.DeleteRatesHandler)
}
