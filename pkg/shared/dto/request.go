package dto

type Request struct {
	Name    string `json:"name"`
	Payload string `json:"payload,omitempty"`
}
