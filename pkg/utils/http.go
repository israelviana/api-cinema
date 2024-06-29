package utils

import (
	"api-cinema/internal/core/domain"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func HTTPFail(ctx *fiber.Ctx, code int, err error, message string) error {
	errJson, _ := json.Marshal(err)

	result := &domain.HTTPErrorResponse{
		Error:   errJson,
		Message: message,
	}

	if err != nil {
		result.ErrorMessage = err.Error()
	}

	return ctx.Status(code).JSON(result)
}
