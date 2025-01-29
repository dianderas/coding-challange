package handlers

import (
	"qr-matrix-service/internal/models"
	"qr-matrix-service/internal/services"
	"qr-matrix-service/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// ProcessMatrix maneja la solicitud de factorización QR.
//
// Recibe:
// - JSON con una matriz rectangular en models.MatrixRequest.
//
// Retorna:
// - 200 OK con la matriz factorizada (Q, R).
// - 400 Bad Request si la matriz es inválida.
// - 500 Internal Server Error si ocurre un error inesperado.
func ProcessMatrix(c *fiber.Ctx) error {
	var request models.MatrixRequest

	// Validacion (Aplicacion) de la entrada
	if err := c.BodyParser(&request); err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid input format")
	}

	result, err := services.FactorizeQr(request)
	if err != nil {
		return utils.HandleServiceError(c, err)
	}

	return utils.JSONSuccess(c, result)
}
