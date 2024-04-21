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

	conn := ctx.Value("connection").(*net.TCPConn)
	service.sender.propagateMessage(conn, dto.Response{
		Name:    dto.NewClientNotificationName,
		Err:     false,
		Payload: fmt.Sprintf("%s joined the chat", username),
	})

	return nil
}

func (service *Service) logout(ctx context.Context) error {
	conn := ctx.Value("connection").(*net.TCPConn)
	client, err := service.server.userByRemoteAddr(conn.RemoteAddr().String())
	if err != nil {
		return err
	}

	username := client.username
	err = service.server.removeClient(client.RemoteAddr())
	if err != nil {
		return err
	}

	service.sender.propagateMessage(conn, dto.Response{
		Name:    dto.ClientDisconnectedActionName,
		Err:     false,
		Payload: fmt.Sprintf("user %s has logged out", username),
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

	return service.sender.propagateMessage(conn, dto.Response{
		Name:    dto.NewMessageNotificationName,
		Err:     false,
		Payload: string(textMessage),
	})
}

func (service *Service) getUsers(ctx context.Context) (users []map[string]string) {
	conn := ctx.Value("connection").(*net.TCPConn)

	for addr, client := range service.server.clients {
		if addr != conn.RemoteAddr().String() {
			users = append(users, map[string]string{
				"name":  client.username,
				"color": string(client.color),
			})
		}
	}

	return
}

func (service *Service) getUsersCount() int {
	return len(service.server.clients) - 1
}
