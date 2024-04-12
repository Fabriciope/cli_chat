package dto

// TODO: colocar um campo com o horario que foi feito o response
type Response struct {
	Name    string      `json:"name"`
	Err     bool        `json:"error"`
	Payload interface{} `json:"payload"`
}
