package services

import (
	"qr-matrix-service/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticateUser_Success(t *testing.T) {
	request := models.LoginRequest{
		User:     "admin",
		Password: "admin",
	}

	token, err := AuthenticateUser(request)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// ❌ Test para usuario vacío
func TestAuthenticateUser_Fail_EmptyUser(t *testing.T) {
	request := models.LoginRequest{
		User:     "",
		Password: "admin",
	}

	_, err := AuthenticateUser(request)

	assert.Error(t, err)
	assert.Equal(t, ErrBadRequest, err)
}

// ❌ Test para contraseña vacía
func TestAuthenticateUser_Fail_EmptyPassword(t *testing.T) {
	request := models.LoginRequest{
		User:     "admin",
		Password: "",
	}

	_, err := AuthenticateUser(request)

	assert.Error(t, err)
	assert.Equal(t, ErrBadRequest, err)
}

// ❌ Test para credenciales incorrectas
func TestAuthenticateUser_Fail_InvalidCredentials(t *testing.T) {
	request := models.LoginRequest{
		User:     "user",
		Password: "wrongpassword",
	}

	_, err := AuthenticateUser(request)

	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
}
