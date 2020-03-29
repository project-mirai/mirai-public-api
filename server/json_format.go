package server

type ResponseInfo struct {
	Success bool        `json:"success"`
	Info    string      `json:"info"`
	Result  interface{} `json:"result"`
}
