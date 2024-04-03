package controller

import (
	"log"
	"strings"

	"github.com/Fabriciope/cli_chat/client/clientapp/controller/handlers"
	"github.com/Fabriciope/cli_chat/shared"
)

type commandHandler func() error
type commandsHandlersMap map[string]commandHandler
type responseHandler func(shared.Response)
type responsesHandlersMap map[string]responseHandler

type Controller struct {
	commandsHandlers  commandsHandlersMap
	responsesHandlers responsesHandlersMap
	handler           *handlers.Handler
}

func NewController(handler *handlers.Handler) *Controller {
	controller := &Controller{
		commandsHandlers:  make(commandsHandlersMap),
		responsesHandlers: make(responsesHandlersMap),
		handler:           handler,
	}

	controller.setHandlerForEachCommand()
	controller.setHandlerForEachResponse()

	return controller
}

func (controller *Controller) setHandlerForEachCommand() {
	controller.commandsHandlers = commandsHandlersMap{
		// TODO: add comands and handlers
	}
}

func (controller *Controller) setHandlerForEachResponse() {
	controller.responsesHandlers = responsesHandlersMap{
		shared.NewClientNotificationName:  controller.handler.NewClientResponseHandler,
		shared.NewMessageNotificationName: controller.handler.NewMessageReceivedHandler,
		shared.SendMessageActionName:      controller.handler.SendMessageInChatResponse,
	}
}

func (controller *Controller) commandHandler(actionName string) (commandHandler, bool) {
	handler, exists := (*controller).commandsHandlers[actionName]
	return handler, exists
}

func (controller *Controller) responseHandler(actionName string) (responseHandler, bool) {
	handler, exists := (*controller).responsesHandlers[actionName]
	return handler, exists
}

func (controller *Controller) LoginHandler() func(string) error {
	return controller.handler.LoginHandler
}

func (controller *Controller) HandleInput(input string) {
	// TODO: serealizar input
	if strings.HasPrefix(input, ":") {
		controller.findHandlerAndRun(input)
		return
	}

	// TODO: tratar erro retornado
	controller.handler.SendMessageInChat(input)
}

func (controller *Controller) findHandlerAndRun(command string) {
	handler, exists := controller.commandHandler(command)
	if !exists {
		// TODO: this command does not exists - exibir em chat line
	}

	handler()
}

func (controller *Controller) HandleResponse(response shared.Response) {
	if response.Err && response.Name == "unknown" {
		// TODO: jogar todos os logs para um arquivo e tratar o erro de outra maneira
		log.Fatalf("error name: %s - msg: %s", response.Name, response.Payload)
		return
	}

	// TODO: tratar erro
	handler, _ := controller.responseHandler(response.Name)
	handler(response)
}
