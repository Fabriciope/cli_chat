package serverapp

import (
	"net"

	"github.com/Fabriciope/cli_chat/shared"
)

var availableColors = []shared.ColorCode{
	//	shared.Black,
	//	shared.Red,
	//	shared.Green,
	//	shared.Yellow,
	//	shared.Blue,
	shared.Magenta,
	shared.Cyan,
	shared.LightGray,
	shared.DarkGray,
	shared.BrightRed,
	shared.BrightGreen,
	shared.BrightYellow,
	shared.BrightBlue,
	shared.BrightMagenta,
	shared.BrightCyan,
	// shared.White,
}

type Client struct {
	connection *net.TCPConn
	username   string
	color      shared.ColorCode
}

func newClient(conn *net.TCPConn, username string, color shared.ColorCode) *Client {
	return &Client{conn, username, color}
}

func (client *Client) RemoteAddr() string {
	return client.connection.RemoteAddr().String()
}

// TODO: adicionar o metodo Conn() para retornar a connection
