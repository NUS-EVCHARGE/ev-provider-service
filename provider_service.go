package main

import (
	"flag"
	"time"

	"github.com/NUS-EVCHARGE/ev-provider-service/config"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/charger"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/provider"
	Rate "github.com/NUS-EVCHARGE/ev-provider-service/controller/rates"
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	_ "github.com/NUS-EVCHARGE/ev-provider-service/docs"
	"github.com/NUS-EVCHARGE/ev-provider-service/handler"
	"github.com/NUS-EVCHARGE/ev-provider-service/helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	r *gin.Engine
)

type DatabaseSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

	var hostname string
	user, pass := helper.GetDatabaseSecrets()

	if pass != "" {
		hostname = user + ":" + pass + "@tcp(evapp-db.c3i0qsy82gn1.ap-southeast-1.rds.amazonaws.com:3306)/evc?parseTime=true&charset=utf8mb4"
	} else {
		hostname = configObj.Dsn // localhost
	}

	// init db
	err = dao.InitDB(hostname)
	if err != nil {
		logrus.WithField("config", configObj).Error("failed to connect to database")
		return
	}

	provider.NewProviderController()
	charger.NewChargerController()
	Rate.NewRateController()
	InitHttpServer(configObj.HttpAddress)
}

func GetDatabaseSecrets() {
	panic("unimplemented")
}

func InitHttpServer(httpAddress string) {
	r = gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authentication"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
	r.GET("provider/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/provider/home", handler.GetBookingHealthCheckHandler)

	// api versioning
	v1 := r.Group("/api/v1")
	v1.POST("/provider", handler.CreateProviderHandler)
	v1.GET("/provider", handler.GetProviderHandler)
	v1.PATCH("/provider", handler.UpdateProviderHandler)
	v1.DELETE("/provider/:provider_id", handler.DeleteProviderHandler)

	v1.POST("/provider/:provider_id/charger", handler.CreateChargerHandler)
	v1.GET("/provider/:provider_id/charger", handler.GetChargerHandler)
	v1.GET("/provider/charger/:charger_id", handler.GetChargerHandler)
	v1.PATCH("/provider/:provider_id/charger", handler.UpdateChargerHandler)
	v1.DELETE("/provider/:provider_id/charger/:charger_id", handler.DeleteChargerHandler)

	v1.POST("/provider/:provider_id/chargerandrate", handler.CreateChargerAndRateHandlerByProviderId)
	v1.GET("/provider/:provider_id/chargerandrate", handler.GetChargerAndRateHandler)
	v1.PATCH("/provider/:provider_id/chargerandrate", handler.UpdateChargerAndRateHandlerByProviderId)

	v1.POST("/provider/:provider_id/rates", handler.CreateRatesHandler)
	v1.GET("/provider/:provider_id/rates", handler.GetRatesHandler)
	v1.GET("/provider/rates/:rates_id", handler.GetRatesHandler)
	v1.PATCH("/provider/:provider_id/rates", handler.UpdateRatesHandler)
	v1.DELETE("/provider/:provider_id/rates/:rates_id", handler.DeleteRatesHandler)
	v1.GET("/charger", handler.GetAllChargerHandler)
}
