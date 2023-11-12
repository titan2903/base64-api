package handlers

import (
	"base64-api/internal/dto"
	"base64-api/pkg/constant"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (h *handlers) EncodeBase64(e echo.Context) error {
	// Use a pointer to EncodeBase64Request instead of new
	data := new(dto.EncodeBase64Request)

	// Bind request data to the EncodeBase64Request struct
	if err := e.Bind(data); err != nil {
		response := dto.ApiResponse("Failed to bind request data", http.StatusBadRequest, constant.StatusError, err.Error())
		return e.JSON(http.StatusBadRequest, response)
	}

	// Call the EncodeBase64 method with data.Data directly
	result, err := h.svc.EncodeBase64(data.Data)
	if err != nil {
		response := dto.ApiResponse("Failed to encode data", http.StatusBadRequest, constant.StatusError, err.Error())
		return e.JSON(http.StatusBadRequest, response)
	}

	response := dto.ApiResponse("Data encoded", http.StatusOK, constant.StatusSuccess, result)
	return e.JSON(http.StatusOK, response)
}

func (h *handlers) DecodeBase64(e echo.Context) error {
	// Use a pointer to DecodeBase64Request instead of new
	data := new(dto.DecodeBase64Request)

	// Bind request data to the DecodeBase64Request struct
	if err := e.Bind(data); err != nil {
		response := dto.ApiResponse("Failed to bind request data", http.StatusBadRequest, constant.StatusError, err.Error())
		return e.JSON(http.StatusBadRequest, response)
	}

	// Call the DecodeBase64 method with data.Data directly
	result, err := h.svc.DecodeBase64(data.Data)
	if err != nil {
		response := dto.ApiResponse("Failed to decode data", http.StatusBadRequest, constant.StatusError, err.Error())
		return e.JSON(http.StatusBadRequest, response)
	}

	response := dto.ApiResponse("Data decoded", http.StatusOK, constant.StatusSuccess, result)
	return e.JSON(http.StatusOK, response)
}

func (h *handlers) HealthCheck(c echo.Context) error {
	var (
		dataResponse dto.Response
	)

	// Define a starting time (e.g., the beginning of your health check operation).
	startTime := time.Now()

	// Simulate a health check operation (you can replace this with your own logic).
	// For this example, we'll just sleep for a few seconds.
	time.Sleep(5 * time.Second)

	// Calculate the duration of the health check operation.
	duration := time.Since(startTime)

	if duration.Seconds() > 10 {
		dataResponse = dto.ApiResponse("Running App failed", http.StatusOK, constant.StatusSuccess, fmt.Sprintf("error: %v", duration.Seconds()))
	} else {
		dataResponse = dto.ApiResponse("Running well", http.StatusOK, constant.StatusSuccess, nil)
	}

	return c.JSON(http.StatusOK, dataResponse)
}
