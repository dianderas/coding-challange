package handlers

import (
	"qr-matrix-service/internal/models"
	"qr-matrix-service/internal/services"
	"qr-matrix-service/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// Login maneja la autenticación HTTP y devuelve un token JWT si las credenciales son válidas.
//
// Recibe un JSON con "user" y "password".
// Responde con un token JWT si las credenciales son correctas, o un error HTTP correspondiente.
//
// Códigos de estado:
// - 200 OK: Autenticación exitosa.
// - 400 Bad Request: Formato de solicitud inválido.
// - 401 Unauthorized: Credenciales incorrectas.
// - 500 Internal Server Error: Error interno al generar el token.
func Login(c *fiber.Ctx) error {
	var request models.LoginRequest

	// Validacion (Aplicacion) de la entrada
	if err := c.BodyParser(&request); err != nil {
		return utils.JSONError(c, fiber.StatusBadRequest, "Invalid input format")
	}

	token, err := services.AuthenticateUser(request)

	if err != nil {
		return utils.HandleServiceError(c, err)
	}

	return utils.JSONSuccess(c, fiber.Map{"token": token})
}
