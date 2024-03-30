package serverapp

import (
	"context"
	"log"
	"strings"

	"github.com/Fabriciope/cli_chat/shared"
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
	// TODO: trocar a lógica de retorno do servico
	username := strings.Trim(request.Payload, " ")
	loggedIn, payload := rh.service.login(ctx, username)
	if loggedIn {
		log.Printf("client %s logged\n\n", username)
	} else {
		log.Printf("cannot log in %s: %s\n\n", username, payload)
	}

	return &shared.Response{
		Name:    request.Name,
		Err:     !loggedIn,
		Payload: payload,
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
