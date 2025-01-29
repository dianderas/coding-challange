package services

import (
	"qr-matrix-service/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ✅ Test para una matriz cuadrada (3x3)
func TestFactorizeQr_Success_SquareMatrix(t *testing.T) {
	request := models.MatrixRequest{
		Values: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	result, err := FactorizeQr(request)

	assert.NoError(t, err)
	assert.NotEmpty(t, result.Q)
	assert.NotEmpty(t, result.R)
}

// ✅ Test para una matriz rectangular válida (4x2)
func TestFactorizeQr_Success_RectangularMatrix(t *testing.T) {
	request := models.MatrixRequest{
		Values: [][]float64{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
		},
	}

	result, err := FactorizeQr(request)

	assert.NoError(t, err)
	assert.NotEmpty(t, result.Q)
	assert.NotEmpty(t, result.R)
}

// ❌ Test para una matriz no rectangular (mal formada)
func TestFactorizeQr_Fail_MalformedMatrix(t *testing.T) {
	request := models.MatrixRequest{
		Values: [][]float64{
			{1, 2, 3},
			{4, 5},
			{6, 7, 8, 9},
		},
	}

	_, err := FactorizeQr(request)

	assert.Error(t, err)
	assert.Equal(t, ErrBadRequest, err)
}

// ❌ Test para una matriz con más columnas que filas (m < n)
func TestFactorizeQr_Fail_MLessThanN(t *testing.T) {
	request := models.MatrixRequest{
		Values: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	_, err := FactorizeQr(request)

	assert.Error(t, err)
	assert.Equal(t, ErrBadRequest, err)
}

// ❌ Test para una matriz vacía
func TestFactorizeQr_Fail_EmptyMatrix(t *testing.T) {
	request := models.MatrixRequest{
		Values: [][]float64{},
	}

	_, err := FactorizeQr(request)

	assert.Error(t, err)
	assert.Equal(t, ErrBadRequest, err)
}
