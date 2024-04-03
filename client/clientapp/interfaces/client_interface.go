package interfaces

import (
	"net"

	"github.com/Fabriciope/cli_chat/client/clientapp/cui"
	"github.com/Fabriciope/cli_chat/shared"
)


type Client interface {
    Conn() *net.TCPConn
    LoggedIn() bool
    SetLoggedInAs(bool)
    CUI() *cui.CUI
    AwaitResponseFromServer()(shared.Response,  error)
    // TODO: estudar injecao ee depedencia e inversao de dependencia
}
