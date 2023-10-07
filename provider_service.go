package main

import (
	"encoding/json"
	"flag"
	"github.com/NUS-EVCHARGE/ev-provider-service/config"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/charger"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/provider"
	Rate "github.com/NUS-EVCHARGE/ev-provider-service/controller/rates"
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	_ "github.com/NUS-EVCHARGE/ev-provider-service/docs"
	"github.com/NUS-EVCHARGE/ev-provider-service/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"time"
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
	var database DatabaseSecret
	secret := os.Getenv("MYSQL_PASSWORD")
	if secret != "" {
		// Parse the JSON data into the struct
		if err := json.Unmarshal([]byte(secret), &database); err != nil {
			logrus.WithField("decodeSecretManager", database).Error("failed to decode value from secret manager")
			return
		}
	}

	if database.Password != "" {
		hostname = database.Username + ":" + database.Password + "@tcp(ev-charger-mysql-db.cdklkqeyoz4a.ap-southeast-1.rds.amazonaws.com:3306)/evc?parseTime=true&charset=utf8mb4"
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
