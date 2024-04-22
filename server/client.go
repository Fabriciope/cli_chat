package server

import (
	"net"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

var availableColors = []escapecode.ColorCode{
	escapecode.Cyan,
	escapecode.Magenta,
	escapecode.Green,
	escapecode.Red,
	escapecode.Yellow,
	escapecode.Blue,
}

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
