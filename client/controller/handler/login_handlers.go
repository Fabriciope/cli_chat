package handler

import (
	"errors"

	"github.com/Fabriciope/cli_chat/pkg/shared"
)

func (handler *Handler) LoginHandler(username string) error {
	request := shared.Request{Name: shared.LoginActionName, Payload: username}
	err := handler.sender.SendRequest(request)
	if err != nil {
		return err
	}

	response, err := handler.user.AwaitResponseFromServer()
	if err != nil {
		return err
	}

	// TODO: trocar esta logica de fazer login
	return handler.LoginResponseHandler(response)
}

func (handler *Handler) LoginResponseHandler(response shared.Response) error {
	if response.Err {
		return errors.New(response.Payload.(string))
	}

	handler.user.SetLoggedInAs(true)
	return nil
}
