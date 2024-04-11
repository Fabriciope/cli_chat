package interfaces

import (
	"net"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/shared"
)

type Client interface {
	Conn() *net.TCPConn
	LoggedIn() bool
	SetLoggedInAs(bool)
	CUI() *cui.CUI
	AwaitResponseFromServer() (shared.Response, error)
	// TODO: estudar injecao ee depedencia e inversao de dependencia
}