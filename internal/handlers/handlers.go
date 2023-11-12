package handlers

import (
	"base64-api/internal/services"

	"github.com/labstack/echo"
)

type handlers struct {
	svc services.Services
}

type Handlers interface {
	EncodeBase64(e echo.Context) error
	DecodeBase64(e echo.Context) error
	HealthCheck(c echo.Context) error
}

func NewHandlers(service services.Services) Handlers {
	return &handlers{svc: service}
}
