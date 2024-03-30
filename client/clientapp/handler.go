package clientapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Fabriciope/cli_chat/client/clientapp/cui"
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

func (handler *handler) newClientResponseHandler(response shared.Response) {
	chatLine := cui.ChatLine{
		Info:      "[insert time]",
		InfoColor: cui.Green,
		Text:      strings.Trim(response.Payload.(string), " "),
	}
	handler.client.cui.DrawNewLineInChat(&chatLine)
}

func (handler *handler) sendMessageInChat(message string) error {
	if !handler.client.loggedIn {
		return errors.New("you must be logged in to send messages in chat")
	}

    // TODO: verificar se a mensagem esta vazia
	err := handler.sender.sendRequest(shared.Request{
		Name:    shared.SendMessageActionName,
		Payload: strings.Trim(message, " "),
	})
	if err != nil {
		return err
	}

	return nil
}

func (handler *handler) sendMessageInChatResponse(response shared.Response) {
	if response.Err {
		// TODO: tratar erro
	}
}

// TODO: resolver problema quando a escrita do usuario ultrapassa o consoleWidth
// TODO: definir uma cor para cada client
func (handler *handler) newMessageReceivedHandler(response shared.Response) {
	var textMessage shared.TextMessage
	json.Unmarshal([]byte(response.Payload.(string)), &textMessage)
	chatLine := cui.ChatLine{
		Info:      fmt.Sprintf("[insert time] From %s:", textMessage.Username),
		InfoColor: cui.BrightYellow,
		Text:      textMessage.Message,
	}
	handler.client.cui.DrawNewLineInChat(&chatLine)
}
