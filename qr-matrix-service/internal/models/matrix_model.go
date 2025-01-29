package models

// Note: Se usa float64 para mayor presicion en caluclos matematicos, asumiendo un contexto en donde los resultados
// son criticos para la operacion de la empresa.
type MatrixRequest struct {
	Values [][]float64 `json:"values"`
}

type MatrixResult struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}
