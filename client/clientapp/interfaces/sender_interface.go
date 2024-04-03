package interfaces

import "github.com/Fabriciope/cli_chat/shared"

type Sender interface {
	SendRequest(shared.Request) error
}
