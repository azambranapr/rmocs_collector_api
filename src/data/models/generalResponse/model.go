package generalResponse

type model struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Count       string      `json:"count"`
	Data        interface{} `json:"data"`
}

type Models []*model

type GeneralResponse model
