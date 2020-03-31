package server

type ResponseInfo struct {
	Success bool        `json:"success"`
	Info    string      `json:"info"`
	Result  interface{} `json:"result"`
}

type BasicPluginInfo struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Author      string   `json:"author"`
	Core        string   `json:"core"`
	Console     string   `json:"console"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Commands    []string `json:"commands"`
}
