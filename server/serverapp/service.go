package serverapp

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/Fabriciope/cli_chat/shared"
)

type Service struct {
	server *Server
	sender *Sender
}

func newService(server *Server) *Service {
	return &Service{
		server: server,
		sender: newSender(server),
	}
}

func (service *Service) hasClient(username string) bool {
	for i := range service.server.clients {
		if service.server.clients[i].username == username {
			return true
		}
	}

	return false
}

func (service *Service) addClient(ctx context.Context, username string) error {
	if service.hasClient(username) {
		return errors.New("user already exists")
	}

	var conn *net.TCPConn = ctx.Value("connection").(*net.TCPConn)
	var client *Client = newClient(conn, username)

	service.server.lock()
	service.server.clients[conn.RemoteAddr()] = client
	service.server.unlock()

	return nil
}

func (service *Service) login(ctx context.Context, username string) (bool, string) {
	if err := service.addClient(ctx, username); err != nil {
		return false, err.Error()
	}

	service.sender.propagateMessage(ctx, shared.Response{
		Name:    shared.NewClientNotificationName,
		Err:     false,
		Payload: fmt.Sprintf("%s joined the chat", username),
	})

	return true, fmt.Sprintf("User %s logged", username)
}
