package dto

// TODO: colocar um campo com o horario que foi feito o request
type Request struct {
	Name    string `json:"name"`
	Payload string `json:"payload,omitempty"`
}
