package dto

type EncodeBase64Request struct {
}

type EncodeBase64Response struct {
	Data string `json:"data"`
}

type DecodeBase64Request struct {
}

type DecodeBase64Response struct {
	Data string `json:"data"`
}
