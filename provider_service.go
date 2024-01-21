package main

import (
	"flag"
	"time"

	"github.com/NUS-EVCHARGE/ev-provider-service/config"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/charger"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/provider"
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
	hostname = user + ":" + pass + "@tcp(evapp-db.c3i0qsy82gn1.ap-southeast-1.rds.amazonaws.com:3306)/evc?parseTime=true&charset=utf8mb4"

	// init db
	err = dao.InitDB(hostname)
	if err != nil {
		logrus.WithField("config", configObj).Error("failed to connect to database")
		return
	}

	provider.NewProviderController()
	charger.NewChargerController()
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

	// provider handler
	v1.POST("/provider", handler.CreateProviderHandler)
	v1.GET("/provider/:provider_email", handler.GetProviderHandler)
	v1.PATCH("/provider", handler.UpdateProviderHandler)
	v1.DELETE("/provider/:provider_id", handler.DeleteProviderHandler)

	// charger handler
	v1.POST("/charger", handler.CreateChargerHandler)
	v1.GET("/charger", handler.GetAllChargerDetailsHandler)
	v1.PATCH("/charger", handler.UpdateChargerHandler)

	// charger point handler
	v1.POST("/chargerpoint", handler.CreateChargerPointHandler)
	v1.PATCH("/chargerpoint", handler.UpdateChargerPointHandler)
}
