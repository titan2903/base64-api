package handler

import "base64-api/internal/services"

type handler struct {
	svc services.Services
}

type Handlers interface {
}

func NewHandlers(service services.Services) Handlers {
	return &handler{svc: service}
}
