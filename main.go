package main

import (
	"myapp/config"
	controllers "myapp/handlers"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.GetConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler: controllers.CustomErrorHandler,
	})
	router := controllers.NewRouter(app)

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Static("/", "./assets")

	app.Use(logger.New(logger.Config{DisableColors: false}))

	router.Setup()

	host := ""
	if cfg.Environment == "development" {
		host = "0.0.0.0"
	}

	go func() {
		err := app.Listen(host + ":" + cfg.Port)
		if err != nil {
			log.Errorf("Error while starting the server %s", err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})

	<-stop
	log.Info("Server is shutting down...")

	if err := app.Shutdown(); err != nil {
		log.Errorf("Server Shutdown Failed:%+v", err)
	}

	close(done)

	<-done
	log.Info("Server stopped gracefully")
}
