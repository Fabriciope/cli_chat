package interfaces

import "github.com/Fabriciope/cli_chat/pkg/shared"

type Sender interface {
	SendRequest(shared.Request) error
}
