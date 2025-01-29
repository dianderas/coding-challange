package services

import (
	"os"
	"qr-matrix-service/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// AuthenticateUser valida las credenciales y genera un token JWT.
//
// Recibe:
// - request: Estructura con "user" y "password".
//
// Retorna:
// - string: Token JWT si la autenticación es exitosa.
// - error: Si ocurre algún problema (credenciales incorrectas, error interno, etc.).
func AuthenticateUser(loginReq models.LoginRequest) (string, error) {
	// Valida que los inputs requeridos no estén vacíos
	if loginReq.User == "" || loginReq.Password == "" {
		return "", ErrBadRequest
	}

	// Verifica credenciales (esto puede moverse a una base de datos en el futuro)
	if loginReq.User != "admin" || loginReq.Password != "admin" {
		return "", ErrInvalidCredentials
	}

	// Genera un token JWT
	token, err := generateJWT(loginReq.User)
	if err != nil {
		return "", ErrInternalError
	}

	return token, nil
}

func generateJWT(username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"name":  username,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
