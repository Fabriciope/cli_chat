package dto

import "github.com/Fabriciope/cli_chat/pkg/escapecode"

type TextMessage struct {
	Username  string               `json:"username"`
	UserColor escapecode.ColorCode `json:"user_color"`
	Message   string               `json:"message"`
}
