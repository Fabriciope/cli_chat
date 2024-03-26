package serverapp

import (
	"context"
	"fmt"

	"github.com/Fabriciope/cli_chat/shared"
)

type Service struct {
	server *Server
	sender *responseSender
}

func newService(server *Server) *Service {
	return &Service{
		server: server,
		sender: newResponseSender(server),
	}
}

func (service *Service) login(ctx context.Context, username string) (bool, string) {
	if err := service.server.addClient(ctx, username); err != nil {
		return false, err.Error()
	}

	service.sender.propagateMessage(ctx, shared.Response{
		Name:    shared.NewClientNotificationName,
		Err:     false,
		Payload: fmt.Sprintf("%s joined the chat", username),
	})

	return true, fmt.Sprintf("User %s logged", username)
}

