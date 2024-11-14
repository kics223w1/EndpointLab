package utils

type HTTPResponse struct {
	Args    map[string]string `json:"args"`
	Data    string             `json:"data"`
	Files   map[string]string  `json:"files"`
	Form    map[string]string  `json:"form"`
	Headers map[string]string `json:"headers"`
	JSON    interface{}        `json:"json"`
	Origin  string             `json:"origin"`
	URL     string             `json:"url"`
	Method  string             `json:"method"`
}