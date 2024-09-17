package main

import (
	"fmt"
	"github.com/fichca/junior-astrologer-service/internal/client"
	"github.com/fichca/junior-astrologer-service/internal/config"
	"github.com/fichca/junior-astrologer-service/internal/server"
	"github.com/fichca/junior-astrologer-service/internal/service"
	"github.com/fichca/junior-astrologer-service/internal/worker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	app := initApp()
	cfg := config.Config{}
	cfg.Parse()

	cl := client.NewClient(cfg.App.Client)
	srv := service.NewService(logger, cl)

	hdl := server.NewHandler(logger, app, srv)
	hdl.RegisterRoutes()

	w := worker.NewWorker(srv, logger)

	if err := w.Start(); err != nil {
		logger.Fatalf("Failed to start worker: %v", err)
	}
	defer w.Stop()

	logger.Info(fmt.Sprintf("✅  server is running on port %v!", cfg.App.HTTP.Port))
	if err := app.Listen(":" + cfg.App.HTTP.Port); err != nil {
		logger.Fatalf("Failed to start app: %v", err)
	}
}

func initApp() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		AppName:           "Junior astrologer service",
	})
	app.Use(recover.New())
	return app
}
