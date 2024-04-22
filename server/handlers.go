package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

type handlersMap map[string]func(context.Context, dto.Request) dto.Response

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

func (rh *RequestHandlers) loginHandler(ctx context.Context, request dto.Request) dto.Response {
	username := strings.Trim(request.Payload, " ")
	err := rh.service.login(ctx, username)
	if err != nil {
		errStr := err.Error()
		return dto.Response{
			Name:    request.Name,
			Err:     true,
			Payload: errStr,
		}
	}

	log.Printf("client %s logged\n", username)
	return dto.Response{
		Name:    request.Name,
		Err:     false,
		Payload: fmt.Sprintf("User %s logged", username),
	}
}

func (rh *RequestHandlers) clientLogout(ctx context.Context, request dto.Request) dto.Response {
	err := rh.service.logout(ctx)
	if err != nil {
		return dto.Response{
			Name:    request.Name,
			Err:     true,
			Payload: err.Error(),
		}
	}

	return dto.Response{
		Name:    request.Name,
		Err:     false,
		Payload: "you have been successfully logged out",
	}
}

func (rh *RequestHandlers) sendMessageInChat(ctx context.Context, request dto.Request) dto.Response {
	message := strings.Trim(request.Payload, " ")
	err := rh.service.sendMessageToEveryone(ctx, message)
	if err != nil {
		return dto.Response{
			Name:    request.Name,
			Err:     true,
			Payload: err.Error(),
		}
	}

	return dto.Response{
		Name:    request.Name,
		Err:     false,
		Payload: "message sent successfully",
	}
}

func (rh *RequestHandlers) getUsers(ctx context.Context, request dto.Request) dto.Response {
	users := rh.service.getUsers(ctx)
	if len(users) == 0 {
		return dto.Response{
			Name:    request.Name,
			Err:     false,
			Payload: "there are no users in this room",
		}
	}

	usersStr, err := json.Marshal(users)
	if err != nil {
		return dto.Response{
			Name:    request.Name,
			Err:     true,
			Payload: "error from server",
		}

	}

	return dto.Response{
		Name:    request.Name,
		Err:     false,
		Payload: string(usersStr),
	}
}

func (rh *RequestHandlers) getUsersCount(ctx context.Context, request dto.Request) dto.Response {
	return dto.Response{
		Name:    request.Name,
		Err:     false,
		Payload: strconv.Itoa(rh.service.getUsersCount()),
	}
}
