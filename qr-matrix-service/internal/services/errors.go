package services

import "errors"

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")    // 401 Unauthorized
	ErrInvalidFormatRequest = errors.New("invalid format request") // 400 Bad Request
	ErrBadRequest           = errors.New("bad request")            // 400 Bad Request
	ErrInternalError        = errors.New("internal server error")  // 500 Internal Server Error
)
