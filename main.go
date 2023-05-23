package main

import (
	"os"

	handler "github.com/absormu/tokped/app/handler"
	md "github.com/absormu/tokped/app/middleware"
	cm "github.com/absormu/tokped/pkg/configuration"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func initHandlers(e *echo.Echo) {
	root := e.Group(cm.Config.RootURL)
	root.POST("/login", handler.LoginHandler)
	root.GET("/job-list", handler.GetJobListHandler)
	root.GET("/job-detail/:id", handler.GetJobDetailHandler)

	root.GET("/order-history/:id", handler.GetOrderHistoryHandler)

	// Start serverlog.Info()
	log.Info("Staring server ...")
}

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})
}

func main() {
	e := echo.New()
	initLogger()

	cm.LoadConfig()

	e.Use(md.AddLogger)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	initHandlers(e)

	err := e.Start(cm.Config.ListenPort)

	if err != nil {
		log.WithField("error", err).Error("Unable to start the server")
		os.Exit(1)
	}
}
