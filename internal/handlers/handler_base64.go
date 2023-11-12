package handlers

import (
	"base64-api/internal/dto"

	"github.com/labstack/echo"
)

func (h *handlers) EncodeBase64(e echo.Context) error {
	var request dto.EncodeBase64Request

	return nil
}

func (h *handlers) DecodeBase64(e echo.Context) error {
	var request dto.DecodeBase64Request

	return nil
}
