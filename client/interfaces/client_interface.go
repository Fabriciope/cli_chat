package interfaces

import (
	"net"

	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

type Client interface {
	Conn() *net.TCPConn
	LoggedIn() bool
	SetLoggedInAs(bool)
	CUI() CUI
	AwaitResponseFromServer() (dto.Response, error)
	// TODO: estudar injecao ee depedencia e inversao de dependencia
}
