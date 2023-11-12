package routes

import (
	"base64-api/internal/handlers"
	"base64-api/internal/services"

	"github.com/labstack/echo"
)

func NewRoutes() *echo.Echo {
	e := echo.New()

	handler := handlers.NewHandlers(services.NewServices())

	e.POST("/v1/base64/encode", handler.EncodeBase64)
	e.POST("/v1/base64/decode", handler.DecodeBase64)

	e.GET("/v1/healthcheck", handler.HealthCheck)

	return e
}
