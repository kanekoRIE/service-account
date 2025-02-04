package main

import (
	"service-account/config"
	"service-account/handler"
	"service-account/helper"
	"service-account/logger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.InitLogger()
	helper.InitValidator()
	err := godotenv.Load()
	if err != nil {
		logger.LogCritical("Error loading .env file", logrus.Fields{})
	}

	db, err := config.InitDB()
	if err != nil {
		logger.LogCritical("Failed to initialize DB", logrus.Fields{"error": err})
	}
	defer db.Close()

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	// all available route
	e.POST("/daftar", handler.Daftar)
	e.POST("/tabung", handler.Tabung)
	e.POST("/tarik", handler.Tarik)
	e.GET("/saldo/:no_rekening", handler.Saldo)

	logger.LogInfo("Application running", logrus.Fields{"port": config.SERVER_PORT, "host": config.SERVER_HOST})
	e.Logger.Fatal(e.Start(config.SERVER_PORT))
}
