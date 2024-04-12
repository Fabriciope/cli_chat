package interfaces

import "github.com/Fabriciope/cli_chat/pkg/shared/dto"

type Sender interface {
	SendRequest(dto.Request) error
}
