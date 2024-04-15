package sender

import "github.com/Fabriciope/cli_chat/pkg/shared/dto"

type SenderInterface interface {
	SendRequest(dto.Request) error
}
