package utils

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ErrorStruct struct {
	Message string `json:"message"`
}

func ErrorToJson(err error) (*[]byte, error) {
	errorStruct := ErrorStruct{
		Message: err.Error(),
	}

	u, err := json.Marshal(errorStruct)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func RespondError(c *fiber.Ctx, status int, errorMessage string) error {
	payload, _ := ErrorToJson(errors.New(errorMessage))

	return c.Status(status).SendString(string(*payload))
}

func ShowErrorPage(status int, errorMessage string) error {
	return fiber.NewError(status, errorMessage)
}
