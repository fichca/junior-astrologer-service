package main

import (
	"fmt"
	"github.com/fichca/junior-astrologer-service/internal/client"
	"github.com/fichca/junior-astrologer-service/internal/config"
	"github.com/fichca/junior-astrologer-service/internal/repository"
	"github.com/fichca/junior-astrologer-service/internal/server"
	"github.com/fichca/junior-astrologer-service/internal/service"
	"github.com/fichca/junior-astrologer-service/internal/worker"
	"github.com/fichca/junior-astrologer-service/pkg/db"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title           Junior astrologer API
// @version         1.0
// @description     Junior astrologer service

// @BasePath  /api
func main() {
	logger := logrus.New()
	app := initApp()
	cfg := config.Config{}
	cfg.Parse()
	minioConnection := db.InitMinioConnection(logger, cfg.DB.Minio)
	dbConnection := db.InitConnection(cfg.DB.Postgre, logger)

	apodRepo := repository.NewAPODRepo(dbConnection)
	minioRepo := repository.NewMinio(minioConnection, cfg.DB.Minio.Bucket)

	if err := db.RunMigrations(dbConnection.DB, cfg.DB.Postgre); err != nil {
		logger.Warning(err)
	}

	cl := client.NewClient(cfg.App.Client)
	srv := service.NewAPODService(logger, cl, minioRepo, apodRepo)

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
	swaggerCfg := swagger.Config{
		FilePath: "./docs/swagger.json",
		Path:     "astrologer/swagger",
		CacheAge: 10,
	}
	app.Use(recover.New())
	app.Use(swagger.New(swaggerCfg))
	app.Use(cors.New())
	return app
}
