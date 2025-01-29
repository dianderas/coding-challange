package handlers

import (
	"log"
	"os"
	"qr-matrix-service/internal/models"
	"qr-matrix-service/internal/services"
	"qr-matrix-service/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// FullProcessMatrix maneja la solicitud de factorización QR e integración con un servicio externo.
//
// Recibe:
// - c *fiber.Ctx: Contexto de Fiber con la solicitud HTTP.
//
// Flujo:
// 1. Extrae la URL del servicio externo desde `STATISTICS_SERVICE_URL`.
// 2. Obtiene el token JWT desde la cabecera `Authorization`.
// 3. Parsea y valida la matriz de entrada (`models.MatrixRequest`).
// 4. Llama al servicio `FactorizeQr` para realizar la factorización QR.
// 5. Envía los datos factorizados a la API externa con autenticación JWT.
// 6. Retorna los resultados combinados de la factorización y del servicio externo.
//
// Retorna:
// - 200 OK con los resultados de `FactorizeQr` y del servicio externo.
// - 400 Bad Request si la entrada es inválida.
// - 500 Internal Server Error si ocurre un error inesperado.
func FullProcessMatrix(c *fiber.Ctx) error {
	extUrl := os.Getenv("STATISTICS_SERVICE_URL")
	fullUrl := strings.TrimRight(extUrl, "/") + "/api/statistics-matrix-process"

	var request models.MatrixRequest
	userToken := c.Get("Authorization")
	log.Println("TOKEN:", userToken)

	// Parsear la solicitud
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato inválido"})
	}

	qrResult, err := services.FactorizeQr(request)
	if err != nil {
		return utils.HandleServiceError(c, err)
	}

	// Llamada a la API externa con autenticación
	client := utils.NewHTTPClient()
	nodeResult, err := client.MakeAuthenticatedRequest("POST", fullUrl, userToken, qrResult)
	if err != nil {
		log.Println("❌ Error en la solicitud HTTP a backend-node:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al comunicarse con el servicio externo"})
	}

	return c.JSON(fiber.Map{
		"qrResult":   qrResult,
		"statistics": nodeResult,
	})
}
