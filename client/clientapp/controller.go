package clientapp

import (
	"errors"
	"log"
	"strings"

	"github.com/Fabriciope/cli_chat/shared"
)

type commandHandler func() error
type commandsHandlersMap map[string]commandHandler
type responseHandler func(shared.Response)
type responsesHandlersMap map[string]responseHandler

type controller struct {
	commandsHandlers  commandsHandlersMap
	responsesHandlers responsesHandlersMap
	handler           *handler

	responsesFromServer <-chan shared.Response
}

func newController(client *Client) *controller {
	controller := &controller{
		commandsHandlers:    make(commandsHandlersMap),
		responsesHandlers:   make(responsesHandlersMap),
		handler:             newHandler(client),
		responsesFromServer: client.responsesFromServer,
	}

	controller.setHandlerForEachCommand()
	controller.setHandlerForEachResponse()

	return controller
}

func (controller *controller) setHandlerForEachCommand() {
	controller.commandsHandlers = commandsHandlersMap{
		// TODO: add comands and handlers
	}
}

func (controller *controller) setHandlerForEachResponse() {
	controller.responsesHandlers = responsesHandlersMap{
		shared.LoginActionName: controller.handler.loginResponseHandler,
	}
}

func (controller *controller) commandHandler(actionName string) (commandHandler, bool) {
	handler, exists := (*controller).commandsHandlers[actionName]
	return handler, exists
}

func (controller *controller) responseHandler(actionName string) (responseHandler, bool) {
	handler, exists := (*controller).responsesHandlers[actionName]
	return handler, exists
}

func (controller *controller) loginHandler() func(string) error {
	return controller.handler.loginHandler
}

func (controller *controller) handleInput(input string) error {
	if strings.HasPrefix(input, ":") {
		return controller.findHandlerAndRun(input)
	}

	return controller.handler.sendMessageInChat(input)
}

func (controller *controller) findHandlerAndRun(command string) error {
	handler, exists := controller.commandHandler(command)
	if !exists {
		return errors.New("this command does not exist")
	}

	return handler()
}

func (controller *controller) catchResponsesAndHandle() {
	for response := range controller.responsesFromServer {
		if response.Err {
			// TODO: jogar todos os logs para um arquivo e tratar o erro de outra maneira
			log.Fatalf("error name: %s - msg: %s", response.Name, response.Payload)
			return
		}

		handler, _ := controller.responseHandler(response.Name)
		handler(response)
	}
}
