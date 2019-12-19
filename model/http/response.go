package http

type ResponseModel struct {
	MessageCode int         `json:"message_code" example:"400"`
	Data        interface{} `json:"data" example:""`
}

type ResponseSuccess struct {
	MessageCode int         `json:"message_code" example:"200"`
	Data        interface{} `json:"data" example:""`
}
