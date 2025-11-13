package helpers

type MetaStruct struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
}

type APIResponseFormatStruct struct {
	Meta MetaStruct
	Data any
}

func APIResponseFormat(code int, status, message string, data any) APIResponseFormatStruct {
	return APIResponseFormatStruct{
		Meta: MetaStruct{
			Code: code,
			Status: status,
			Message: message,
		},
		Data: data,
	}
}