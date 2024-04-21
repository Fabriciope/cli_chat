package server

import (
	"net"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

// TODO: trocar esquema de cores do sistema
var availableColors = []escapecode.ColorCode{
	//	escapecode.Black,
	//	escapecode.Red,
	//	escapecode.Green,
	//	escapecode.Yellow,
	//	escapecode.Blue,
	escapecode.Magenta,
	escapecode.Cyan,
	escapecode.LightGray,
	escapecode.DarkGray,
	escapecode.BrightRed,
	escapecode.BrightGreen,
	escapecode.BrightYellow,
	escapecode.BrightBlue,
	escapecode.BrightMagenta,
	escapecode.BrightCyan,
	// escapecode.White,
}

// TODO: adaptar para json
type client struct {
	connection *net.TCPConn
	username   string
	color      escapecode.ColorCode
}

func newClient(conn *net.TCPConn, username string, color escapecode.ColorCode) *client {
	return &client{conn, username, color}
}

func (client *client) RemoteAddr() string {
	return client.connection.RemoteAddr().String()
}

func (client *client) Conn() *net.TCPConn {
	return client.connection
}
