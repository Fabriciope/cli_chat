package serverapp

import "net"

type Client struct {
	connection *net.TCPConn
	username   string
}

func newClient(conn *net.TCPConn, username string) *Client {
	return &Client{conn, username}
}

func (client *Client) RemoteAddr() string {
	return client.connection.RemoteAddr().String()
}
