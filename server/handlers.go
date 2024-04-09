package server

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Fabriciope/cli_chat/pkg/shared"
)

type handlersMap map[string]func(context.Context, shared.Request) *shared.Response

type RequestHandlers struct {
	server  *Server
	service *Service
}

func newRequestHandlers(server *Server) *RequestHandlers {
	return &RequestHandlers{
		server:  server,
		service: newService(server),
	}
}

func (rh *RequestHandlers) loginHandler(ctx context.Context, request shared.Request) *shared.Response {
	username := strings.Trim(request.Payload, " ")
	err := rh.service.login(ctx, username)
	if err != nil {
		errStr := err.Error()
		log.Printf("cannot log in %s: %s\n\n", username, errStr)
		return &shared.Response{
			Name:    request.Name,
			Err:     true,
			Payload: errStr,
		}
	}

	log.Printf("client %s logged\n\n", username)
	return &shared.Response{
		Name:    request.Name,
		Err:     false,
		Payload: fmt.Sprintf("User %s logged", username),
	}
}

func (rh *RequestHandlers) sendMessageInChat(ctx context.Context, request shared.Request) *shared.Response {
	message := strings.Trim(request.Payload, " ")
	err := rh.service.sendMessageToEveryone(ctx, message)
	if err != nil {
		return &shared.Response{
			Name:    request.Name,
			Err:     true,
			Payload: err.Error(),
		}
	}

	return &shared.Response{
		Name:    request.Name,
		Err:     false,
		Payload: "message sent successfully",
	}
}
