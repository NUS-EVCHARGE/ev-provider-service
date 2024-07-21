package main

import (
	"flag"
	"time"

	"github.com/NUS-EVCHARGE/ev-provider-service/controller/authentication"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/license"
	"github.com/NUS-EVCHARGE/ev-provider-service/helper"

	"github.com/NUS-EVCHARGE/ev-provider-service/config"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/charger"
	"github.com/NUS-EVCHARGE/ev-provider-service/controller/provider"
	"github.com/NUS-EVCHARGE/ev-provider-service/dao"
	_ "github.com/NUS-EVCHARGE/ev-provider-service/docs"
	"github.com/NUS-EVCHARGE/ev-provider-service/handler"
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
	hostname = user + ":" + pass + "@tcp(evapp-db.cbyk62is0npt.ap-southeast-1.rds.amazonaws.com:3306)/evc?parseTime=true&charset=utf8mb4"
	//hostname = configObj.Dsn

	// init db
	err = dao.InitDB(hostname)
	if err != nil {
		logrus.WithField("config", configObj).Error("failed to connect to database")
		panic("db failed")
		return
	}

	provider.NewProviderController()
	charger.NewChargerController()
	license.NewLicenseController()
	authentication.NewAuthenticationController()
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
		AllowHeaders:     []string{"authorization", "access-control-allow-origin", "Access-Control-Allow-Headers", "Origin", "Content-Length", "Content-Type", "authentication", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
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
	{
		v1.POST("/login", handler.LoginHandler)
		v1.POST("/signup", handler.SignUpHandler)
		v1.POST("/confirm", handler.ConfirmUserHandler)
		v1.POST("/resend", handler.ResendChallengeCodeHandler)
		v1.GET("/ws", handler.WsChargerEndpoint)
		v1.GET("/charger/chargerstatus", handler.GetChargerEndpointStatus)
		v1.POST("/charger/chargerstatus", handler.SetChargerEndpointStatus)
	}

	protectedV1 := r.Group("/api/v1")
	protectedV1.Use(handler.AuthMiddlewareHandler)
	{
		// provider handler
		protectedV1.POST("/provider", handler.CreateProviderHandler)
		protectedV1.GET("/provider/:provider_email", handler.GetProviderHandler)
		protectedV1.PATCH("/provider", handler.UpdateProviderHandler)
		protectedV1.DELETE("/provider/:provider_id", handler.DeleteProviderHandler)

		// charger handler
		protectedV1.POST("/charger", handler.CreateChargerHandler)
		protectedV1.GET("/charger", handler.GetAllChargerDetailsHandler)
		protectedV1.PATCH("/charger", handler.UpdateChargerHandler)

		// charger point handler
		protectedV1.POST("/chargerpoint", handler.CreateChargerPointHandler)
		protectedV1.PATCH("/chargerpoint", handler.UpdateChargerPointHandler)

		// authentication handler
		protectedV1.POST("/logout", handler.LogoutUserHandler)

		// license handler
		protectedV1.PATCH("/license", handler.UpdateLicenseByCompanyHandler)
		protectedV1.GET("/license", handler.GetLicenseByCompanyHandler)
	}
}

func healthCheckPolling() {
	// N providers,  M number of chargers -> N * M operations for health check (performance heavy) -> N threads -> high CPU usage
	// autoscale -> scale service I * N * M -> healthcehcks operations -> each charger point to have two health request

	// event driven
	// service -> check all heart
	// all heartbeat report to the services
}
