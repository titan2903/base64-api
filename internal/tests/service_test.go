package tests

import (
	"base64-api/internal/services"
	"encoding/base64"
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	originalData := "Hello, World!"

	// Encode
	expectedEncodedData := base64.StdEncoding.EncodeToString([]byte(originalData))
	encodedData, err := services.NewServices().EncodeBase64(originalData)
	if err != nil {
		t.Errorf("Unexpected error during encoding: %v", err)
	}

	// Check if the result matches the expected value
	if encodedData != expectedEncodedData {
		t.Errorf("EncodeBase64() returned incorrect result. Expected: %s, Got: %s", expectedEncodedData, encodedData)
	}
}

func TestDecodeBase64(t *testing.T) {
	originalData := "Hello, World!"

	// Encode original data
	encodedData := base64.StdEncoding.EncodeToString([]byte(originalData))

	// Decode
	expectedDecodedData := originalData
	decodedData, err := services.NewServices().DecodeBase64(encodedData)
	if err != nil {
		t.Errorf("Unexpected error during decoding: %v", err)
	}

	// Check if the result matches the expected value
	if decodedData != expectedDecodedData {
		t.Errorf("DecodeBase64() returned incorrect result. Expected: %s, Got: %s", expectedDecodedData, decodedData)
	}
}
