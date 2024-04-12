package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
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

func (service *Service) login(ctx context.Context, username string) error {
	if err := service.server.addClient(ctx, username); err != nil {
		return err
	}

	service.sender.propagateMessage(ctx, dto.Response{
		Name:    dto.NewClientNotificationName,
		Err:     false,
		Payload: fmt.Sprintf("%s joined the chat", username),
	})

	return nil
}

func (service *Service) sendMessageToEveryone(ctx context.Context, message string) error {
	conn := ctx.Value("connection").(*net.TCPConn)

	client, err := service.server.userByRemoteAddr(conn.RemoteAddr().String())
	if err != nil {
		return err
	}

	textMessage, err := json.Marshal(dto.TextMessage{
		Username:  client.username,
		UserColor: client.color,
		Message:   message,
	})
	if err != nil {
		return err
	}

	return service.sender.propagateMessage(ctx, dto.Response{
		Name:    dto.NewMessageNotificationName,
		Err:     false,
		Payload: string(textMessage),
	})
}
