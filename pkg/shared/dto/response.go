package dto

type Response struct {
	Name    string      `json:"name"`
	Err     bool        `json:"error"`
	Payload interface{} `json:"payload"`
}
