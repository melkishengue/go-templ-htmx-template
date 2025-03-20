package handlers

import (
	"myapp/logger"
	"myapp/utils"
	"myapp/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type HomeHandlers struct {
}

func NewHomeHandlers() *HomeHandlers {
	return &HomeHandlers{}
}

func (h *HomeHandlers) HandleViewHome(c *fiber.Ctx) error {
	logger.Show("lfjsldfs fasjdfölasdjfs")
	pageContent := views.HomeIndex()

	if utils.IsHTMXRequest(c) {
		err := utils.RenderComponent(c, pageContent)

		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString("An error occurred. Please retry")
		}

		return nil
	}

	page := views.Home("", pageContent)

	handler := adaptor.HTTPHandler(templ.Handler(page))

	return handler(c)
}
