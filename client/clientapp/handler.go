package clientapp

import (
	"errors"

	"github.com/Fabriciope/cli_chat/shared"
)

type handler struct {
	client *Client
	sender sender
}

func newHandler(client *Client) *handler {
	return &handler{
		client: client,
		sender: newRequestSender(client.connection),
	}
}

func (handler *handler) loginHandler(username string) error {
	request := shared.Request{Name: shared.LoginActionName, Payload: username}
	err := handler.sender.sendRequest(request)
	if err != nil {
		return err
	}

	response, err := handler.client.awaitResponseFromServer()
	if err != nil {
		return err
	}

	return handler.loginResponseHandler(response)
}

func (handler *handler) loginResponseHandler(response shared.Response) error {
	if response.Err {
		return errors.New(response.Payload.(string))
	}

	handler.client.loggedIn = true
    handler.client.cui.SetLoggedAs(true)
	return nil
}

func (handler *handler) sendMessageInChat(message string) error {
	// TODO: do this
	return nil
}
