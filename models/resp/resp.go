package resp

type RespModel struct {
	IsSuccess bool   `json:"is_success"`
	Error     string `json:"error"`
	Message   string `json:"message"`
}
