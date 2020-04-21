package structs

type HttpResponseModel struct {
	Errors []string     `json:"errors"`
	Result interface{} `json:"result"`
}
