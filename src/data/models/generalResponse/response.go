package generalResponse

const (
	Error   = "error"
	Message = "successful"
)

func NewResponse(messageType string, message string, count string, data interface{}) GeneralResponse {
	return GeneralResponse{
		MessageType: messageType,
		Message:     message,
		Count:       count,
		Data:        data,
	}
}
