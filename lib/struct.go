package lib

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"result,omitempty"`
}
