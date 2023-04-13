package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamolegga/enviper"
	"github.com/kwahome/cards-deck-api/config"
	"github.com/kwahome/cards-deck-api/internal/api/healthcheck"
	"github.com/kwahome/cards-deck-api/internal/api/v1"
	"github.com/kwahome/cards-deck-api/pkg/http/middlewares"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var conf config.Config

func main() {
	initConfig()

	err := RunHttpServer(conf)
	if err != nil {
		logrus.Fatal("An error has occurred while starting the web server: ", err)
	}
}

func initConfig() {
	v := enviper.New(viper.New())

	v.AddConfigPath(".")
	v.SetConfigName(".app")

	if err := v.Unmarshal(&conf); err != nil {
		logrus.Fatal("An error has occurred while parsing config file: ", err)
	}

	logrus.Info(fmt.Sprintf("using the config file: %s", v.ConfigFileUsed()))
}

// RunHttpServer starts the http server.
func RunHttpServer(config config.Config) error {
	gin.SetMode(gin.ReleaseMode)
	if config.HTTP.Debug {
		gin.SetMode(gin.DebugMode)
	}

	/* ---------------------------  Create router  --------------------------- */
	router := gin.Default()

	/* ------------------------  Recovery middleware  ------------------------ */
	router.Use(gin.Recovery())

	/* ------------------------  RequestId middleware  ----------------------- */
	router.Use(middlewares.RequestIdMiddleware())

	/* ---------------------------  Public routes  --------------------------- */
	healthcheck.RegisterRoutes(router)

	/* --------------------------  Auth middleware  -------------------------- */
	router.Use(middlewares.TokenAuthMiddleware(config.AuthToken))

	/* ---------------------------  Private routes  --------------------------- */
	v1.RegisterRoutes(router)

	address := net.JoinHostPort(config.HTTP.Host, config.HTTP.Port)
	server := http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	logrus.Info(fmt.Sprintf("HTTP server is listening on port: %s", conf.HTTP.Port))

	return server.ListenAndServe()
}
