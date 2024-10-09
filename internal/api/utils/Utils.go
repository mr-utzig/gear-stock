package utils

type response struct {
	status  bool
	message string
	data    interface{}
}

func NewResponse(status bool, message string, data interface{}) response {
	return response{status: status, message: message, data: data}
}
