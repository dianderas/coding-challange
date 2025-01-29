package services

import (
	"log"
	"qr-matrix-service/internal/models"

	"gonum.org/v1/gonum/mat"
)

// FactorizeQr realiza la factorización QR de una matriz rectangular.
//
// Recibe:
// - request: models.MatrixRequest con una matriz rectangular m × n (m >= n).
//
// Retorna:
// - result: models.MatrixResult con las matrices Q y R de la factorización QR.
// - err:
//   - ErrBadRequest si la matriz no es válida (mal formada o m < n).
//   - ErrInternalError si ocurre un error inesperado.
//
// Manejo de errores:
// - Captura cualquier panic con recover() y lo convierte en ErrInternalError.
func FactorizeQr(request models.MatrixRequest) (result models.MatrixResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("❌ Panic detectado en FactorizeQr:", r)
			result = models.MatrixResult{}
			err = ErrInternalError // Maneja error no controlado
		}
	}()

	matrix := request.Values
	// Validacion (Negocio) de errores controlados
	if !isValidMatrix(matrix) {
		return models.MatrixResult{}, ErrBadRequest
	}

	// Convierte la matriz a mat.Dense
	rows := len(matrix)
	numCols := len(matrix[0])
	data := make([]float64, 0, rows*numCols)

	for _, row := range matrix {
		data = append(data, row...)
	}

	A := mat.NewDense(rows, numCols, data)

	// Factorizacion QR
	var qr mat.QR
	qr.Factorize(A)

	// Extraer Q y R
	var Q, R mat.Dense
	qr.QTo(&Q)
	qr.RTo(&R)

	// Convertir Q y R a [][]float64
	QMatrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		QMatrix[i] = make([]float64, numCols)
		for j := 0; j < numCols; j++ {
			QMatrix[i][j] = Q.At(i, j)
		}
	}

	RMatrix := make([][]float64, numCols)
	for i := 0; i < numCols; i++ {
		RMatrix[i] = make([]float64, numCols)
		for j := 0; j < numCols; j++ {
			RMatrix[i][j] = R.At(i, j)
		}
	}

	return models.MatrixResult{
		Q: QMatrix,
		R: RMatrix,
	}, nil
}

// isValidMatrix verifica si una matriz es válida para la factorización QR.
//
// Condiciones de validez:
// - Todas las filas deben tener la misma cantidad de columnas.
// - Debe cumplir m >= n (más filas que columnas o cuadrada).
//
// Retorna:
// - true si la matriz es válida.
// - false si la matriz está mal formada o no cumple m >= n.
func isValidMatrix(matrix [][]float64) bool {
	rows := len(matrix)

	// Valida que haya al menos una fila y que la primera tenga contenido
	if rows == 0 || len(matrix[0]) == 0 {
		log.Println("Error: Matriz vacía o primera fila sin columnas")
		return false
	}

	numCols := len(matrix[0])

	// Validar que todas las filas tengan el mismo número de columnas
	for i, row := range matrix {
		if len(row) != numCols {
			log.Printf("Error: Fila %d tiene %d columnas, esperado %d\n", i, len(row), numCols)
			return false
		}
	}

	// Validar que haya al menos tantas filas como columnas (m >= n)
	if rows < numCols {
		log.Printf("Error: Matriz no válida para factorizar (m=%d, n=%d)\n", rows, numCols)
		return false
	}

	return true
}
