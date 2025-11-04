package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/melkishengue/gotemplate/cmd/backend/city"
	db "github.com/melkishengue/gotemplate/internal/database"
	"github.com/melkishengue/gotemplate/internal/logger"
	middleware "github.com/melkishengue/gotemplate/internal/middlewares"
	"github.com/melkishengue/gotemplate/pkg/utils"
	_ "github.com/melkishengue/gotemplate/spec"
)

//	@title			gotemplate-app
//	@version		1.0
//	@description	A well-rounded template with everything you need to get started in golang, htmx and templ
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	gotemplate-app Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	gotemplate-app@my-app.com

// @Host      localhost:3020
// @BasePath  /api/v1

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	_ = godotenv.Load()

	logger.SetUp(utils.GetEnvOrDie("ENVIRONMENT"))
	db.InitDatabase(&gorm.Config{})

	defer func() {
		sqlDB, err := db.Connection.DB()
		if err != nil {
			log.Fatalf("Failed to get underlying *sql.DB: %v", err)
		}

		sqlDB.Close()
	}()

	g := gin.New()
	g.Use(
		middleware.SlogMiddleware(),
		gin.Recovery(),
	)

	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3010", "http://localhost:4010"}, // original backend host and proxy
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	v1 := g.Group("/api/v1")
	{
		v1.Use(middleware.CacheMiddleware())
		v1.GET("/countries/:country_code/cities", city.CitiesHandler)
		v1.GET("/countries/:country_code/cities/:id", city.CityHandler)
	}

	server := &http.Server{
		Addr:    ":3020",
		Handler: g,
	}

	stopped := make(chan struct{})

	go func() {
		defer close(stopped)

		slog.Info("starting server", slog.String("address", server.Addr))

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed to start or run unexpectedly", slog.Any("err", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-quit:
		slog.Info("Shutting down gracefully", "signal", sig)
	case <-stopped:
		slog.Error("Server stopped unexpectedly, initiating shutdown...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("HTTP Server forced to shutdown", "err", err)
	} else {
		slog.Info("HTTP Server successfully shut down")
	}
}
