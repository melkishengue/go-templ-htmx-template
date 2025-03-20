package handlers

import (
	"errors"
	"fmt"
	"myapp/logger"
	"myapp/utils"
	"myapp/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := "Une erreur s'est produite"

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	log.Infof("Error with status code %d has occurred", code)

	// we need to send 200 otherwise htmx will not swap content and nothing will happen on the frontend
	c.Status(fiber.StatusOK)

	var errorIndex templ.Component

	logger.Show(e.Message)

	switch code {
	case 401:
		errorIndex = views.Error401(message)
	case 404:
		errorIndex = views.Error404(message)
	default:
		errorIndex = views.Error500(code, e.Message)
	}

	if utils.IsHTMXRequest(c) {
		err := utils.RenderComponent(c, errorIndex)

		if err != nil {
			c.Status(fiber.StatusOK).SendString("An error occurred. Please retry")
		}

		return nil
	}

	pageTitle := fmt.Sprintf(" | Error %d", code)
	page := views.ErrorPage(pageTitle, errorIndex)

	handler := adaptor.HTTPHandler(templ.Handler(page))

	return handler(c)
}
