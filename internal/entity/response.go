package entity

type ResponseDTO struct {
	StatusCode int
	Message    string
	Data       interface{}
}

func HandleResponseDTO(statusCode int, message string, data interface{}) *ResponseDTO {
	return &ResponseDTO{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
