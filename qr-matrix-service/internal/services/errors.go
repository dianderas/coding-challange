package services

import "errors"

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrInvalidFormatRequest = errors.New("invalid format request")
	ErrBadRequest           = errors.New("bad request")
	ErrInternalError        = errors.New("internal server error")
)
