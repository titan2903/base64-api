package dto

type EncodeBase64Request struct {
	Data string `json:"data"`
}

type EncodeBase64Response struct {
	Data string `json:"data"`
}

type DecodeBase64Request struct {
	Data string `json:"data"`
}

type DecodeBase64Response struct {
	Data string `json:"data"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Data: data,
		Meta: meta,
	}

	return jsonResponse
}

type (
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	}
	Response struct {
		Data interface{} `json:"data"`
		Meta Meta        `json:"meta"`
	}
)
