package utils

import (
	"qr-matrix-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

// HandleServiceError maneja errores de servicios y devuelve una respuesta HTTP adecuada.
func HandleServiceError(c *fiber.Ctx, err error) error {
	switch err {
	case services.ErrBadRequest:
		return JSONError(c, fiber.StatusBadRequest, err.Error())

	case services.ErrInternalError:
		return JSONError(c, fiber.StatusInternalServerError, "Internal Server Error")

	default:
		return JSONError(c, fiber.StatusInternalServerError, "Unexpected error occurred")
	}
}
