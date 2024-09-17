package main

import (
	"fmt"
	"github.com/fichca/junior-astrologer-service/internal/config"
	"github.com/fichca/junior-astrologer-service/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	app := initApp()
	cfg := config.Config{}
	cfg.Parse()
	handler := server.NewHandler(logger, app, nil)

	handler.RegisterRoutes()
	logger.Info(fmt.Sprintf("✅  server is running on port %v!", cfg.App.HTTP.Port))
	err := app.Listen(":" + cfg.App.HTTP.Port)
	if err != nil {
		logger.Fatal(err)
	}
}
func initApp() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		AppName:           "APOD service",
	})
	app.Use(recover.New())
	return app
}
