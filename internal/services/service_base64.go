package services

import (
	"encoding/base64"
)

func (s *services) EncodeBase64(data string) (string, error) {
	// Encode data to base64
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedData, nil
}

func (s *services) DecodeBase64(encodedData string) (string, error) {
	// Decode base64 data
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}
	return string(decodedData), nil

}
