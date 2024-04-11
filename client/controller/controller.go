package controller

import (
	"log"
	"strings"

	"github.com/Fabriciope/cli_chat/client/controller/handler"
	"github.com/Fabriciope/cli_chat/pkg/shared"
)

type Controller struct {
	commandsHandlers  handler.CommandsHandlersMap
	responsesHandlers handler.ResponsesHandlersMap
	handler           *handler.Handler
}

func NewController(h *handler.Handler) *Controller {
	controller := &Controller{
		commandsHandlers:  make(handler.CommandsHandlersMap),
		responsesHandlers: make(handler.ResponsesHandlersMap),
		handler:           h,
	}

	controller.setHandlerForEachCommand()
	controller.setHandlerForEachResponse()

	return controller
}

func (controller *Controller) setHandlerForEachCommand() {
	controller.commandsHandlers = handler.CommandsHandlersMap{
		// TODO: add comands and handler
	}
}

func (controller *Controller) setHandlerForEachResponse() {
	controller.responsesHandlers = handler.ResponsesHandlersMap{
		shared.NewClientNotificationName:  controller.handler.NewClientResponseHandler,
		shared.NewMessageNotificationName: controller.handler.NewMessageReceivedHandler,
		shared.SendMessageActionName:      controller.handler.SendMessageInChatResponse,
	}
}

func (controller *Controller) commandHandler(actionName string) (handler.CommandHandler, bool) {
	handler, exists := (*controller).commandsHandlers[actionName]
	return handler, exists
}

func (controller *Controller) responseHandler(actionName string) (handler.ResponseHandler, bool) {
	handler, exists := (*controller).responsesHandlers[actionName]
	return handler, exists
}

func (controller *Controller) LoginHandler() func(string) error {
	return controller.handler.LoginHandler
}

func (controller *Controller) HandleInput(input string) {
	input = strings.Trim(input, " ")

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
		log.Fatalf("error name: %s - msg: %s", response.Name, response.Payload)
		return
	}

	// TODO: tratar erro
	handler, _ := controller.responseHandler(response.Name)
	handler(response)
}
