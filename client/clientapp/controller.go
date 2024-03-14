package clientapp

import (
	"log"

	"github.com/Fabriciope/cli_chat/shared"
)

type inputsHandlersMap map[string]func(string)
type inputHandler func(string)
type responsesHandlersMap map[string]func(shared.Response)
type responseHandler func(shared.Response)

type controller struct {
	inputsHandlers    inputsHandlersMap
	responsesHandlers responsesHandlersMap

	client *Client
}

func newController(client *Client) *controller {
	var controller = &controller{
		inputsHandlers:    make(inputsHandlersMap),
		responsesHandlers: make(responsesHandlersMap),
		client:            client,
	}

	var handlers *handler = newHandlers(client)
	controller.setHandlerForEachInput(handlers)
	controller.setHandlerForEachResponse(handlers)

	return controller
}

func (controller *controller) setHandlerForEachInput(handlers *handler) {
	controller.inputsHandlers = inputsHandlersMap{
		shared.LoginActionName: handlers.loginInputHandler,
	}
}

func (controller *controller) setHandlerForEachResponse(handlers *handler) {
	controller.responsesHandlers = responsesHandlersMap{
		shared.LoginActionName: handlers.loginResponseHandler,
	}
}

func (controller *controller) inputHandler(actionName string) inputHandler {
	return (*controller).inputsHandlers[actionName]
}

func (controller *controller) responseHandler(actionName string) responseHandler {
	return (*controller).responsesHandlers[actionName]
}

// TODO: quando estiver dentro do chat mudar o comeco do input para type message:

func (controller *controller) handleInput(input string) error {
	switch input {
	case "logout":
		return nil
	default:
		return nil
	}
}

func (controller *controller) catchResponsesAndHandle() {
	for response := range controller.client.responsesFromServer {
		if response.Err {
			log.Fatalf("error name: %s - msg: %s", response.Name, response.Payload)
			return
		}

		controller.responseHandler(response.Name)(response)
			
	}
}
