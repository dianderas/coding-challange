package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"qr-matrix-service/internal/models"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

// Note: por ahora este test integracion no esta hecho para funciona dentro de la red de containers sino es externo.
// por ello si los containers estan levantados localmente apuntamos a localhost.
func TestFullProcessMatrixIntegration(t *testing.T) {
	// Inicializar la app Fiber
	app := fiber.New()

	// Asegurar que `STATISTICS_SERVICE_URL` está configurado
	os.Setenv("STATISTICS_SERVICE_URL", "http://localhost:3000")

	// Registrar el handler en Fiber
	app.Post("/api/fullmatrix-process", FullProcessMatrix)

	// Crear payload válido
	payload := models.MatrixRequest{
		Values: [][]float64{
			{1, 2},
			{3, 4},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	assert.NoError(t, err)

	// Simular una solicitud HTTP real a `backend-go`
	req := httptest.NewRequest("POST", "/api/fullmatrix-process", bytes.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json") // ✅ Asegurar que el Content-Type es correcto
	req.Header.Set("Authorization", "Bearer "+generateTestJWT())

	// Ejecutar la solicitud en Fiber
	resp, err := app.Test(req, int(10*time.Second))
	if err != nil {
		t.Fatalf("❌ Error en la solicitud: %v", err)
	}

	if resp == nil {
		t.Fatalf("❌ Respuesta es `nil`, backend-go no respondió.")
	}

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Verificar respuesta HTTP 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Imprimir resultado para depuración
	log.Println("Respuesta de backend-go:", string(body))

	// Verificar que la respuesta tiene los campos esperados
	assert.Contains(t, string(body), "qrResult")
	assert.Contains(t, string(body), "statistics")
}

func generateTestJWT() string {
	secret := []byte("my-secret") // Usa el mismo `JWT_SECRET` que `backend-go`
	claims := jwt.MapClaims{
		"name":  "admin",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Expira en 3 días
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(secret)
	return signedToken
}
