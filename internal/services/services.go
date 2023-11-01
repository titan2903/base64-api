package services

import "base64-api/internal/dto"

type services struct {
}

type Services interface {
	EncodeBase64(req dto.EncodeBase64Request) (res dto.EncodeBase64Response, err error)
	DecodeBase64(req dto.DecodeBase64Request) (res dto.DecodeBase64Response, err error)
}

func NewServices() Services {
	return &services{}
}
