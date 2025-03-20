package utils

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func IsHTMXRequest(c *fiber.Ctx) bool {
	if c.Get("HX-History-Restore-Request") == "true" {
		return false
	}

	return c.Get("HX-Request") == "true"
}

func RenderComponent(c *fiber.Ctx, component templ.Component) error {
	var buf bytes.Buffer

	// c.Set("Cache-Control", "no-cache, no-store, must-revalidate")

	if err := component.Render(c.Context(), &buf); err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("An error occurred. Please retry")
		return err
	}

	if _, err := c.Write([]byte(buf.String())); err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("An error occurred. Please retry")
		return err
	}

	return nil
}
