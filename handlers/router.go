package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type Router struct {
	App *fiber.App

	HomeHandlers *HomeHandlers
}

func NewRouter(app *fiber.App) *Router {
	return &Router{
		App:          app,
		HomeHandlers: NewHomeHandlers(),
	}
}

func (r *Router) Setup() {
	r.App.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/.well-known/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/.well-known/ready",
	}))

	r.App.Use(helmet.New(helmet.Config{CrossOriginEmbedderPolicy: "false"}))
	r.App.Use(limiter.New(limiter.Config{
		Max: 30, // per minute
	}))

	root := r.App.Group("/")

	root.Get("/", r.HomeHandlers.HandleViewHome)
	root.Get("/test", r.HomeHandlers.HandleViewHome)

	/* ↓ Not Found Management - Fallback Page ↓ */
	r.App.Get("/*", func(c *fiber.Ctx) error {
		return fiber.NewError(
			fiber.StatusNotFound,
			"error 404: not found",
		)
	})
}
