package utils

type response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(status bool, message string, data interface{}) *response {
	return &response{Status: status, Message: message, Data: data}
}
