package helpers

type MetaFormat struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type APIResponseFormat struct {
	Meta MetaFormat `json:"meta"`
	Data any        `json:"data"`
}

func APIResponse(code int, status, message string, data any) APIResponseFormat {
	return APIResponseFormat{
		Meta: MetaFormat{
			Code:    code,
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}
