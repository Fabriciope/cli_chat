package clientapp

import (
	"fmt"

	"github.com/Fabriciope/cli_chat/shared"
)

type handler struct {
	client *Client
	sender sender
}

func newHandlers(client *Client) *handler {
	return &handler{
		client: client,
		sender: client.requestSender,
	}
}

func (handler *handler) loginInputHandler(username string) {
	handler.sender.sendRequest(shared.Request{
		Name:    shared.LoginActionName,
		Payload: username,
	})
}

func (handlers *handler) loginResponseHandler(response shared.Response) {
	fmt.Println(response)
}