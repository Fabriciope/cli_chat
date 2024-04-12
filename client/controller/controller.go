package controller

import (
	"errors"
	"strings"

	"github.com/Fabriciope/cli_chat/client/controller/handler"
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
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
		dto.NewClientNotificationName:  controller.handler.NewClientResponseHandler,
		dto.NewMessageNotificationName: controller.handler.NewMessageReceivedHandler,
		dto.SendMessageActionName:      controller.handler.SendMessageInChatResponse,
	}
}

func (controller *Controller) commandHandler(command string) (handler.CommandHandler, error) {
	handler, exists := (*controller).commandsHandlers[command]
	if exists != false {
		return nil, errors.New("handler for this command does not exist")
	}

	return handler, nil
}

func (controller *Controller) responseHandler(actionName string) (handler.ResponseHandler, error) {
	handler, exists := (*controller).responsesHandlers[actionName]
	if exists == false {
		return nil, errors.New("handler for this action name does not exists")
	}

	return handler, nil
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

	controller.handler.SendMessageInChat(input)
}

func (controller *Controller) findHandlerAndRun(command string) {
	handler, err := controller.commandHandler(command)
	if err != nil {
		controller.handler.CUI().DrawNewLineInChat(&cui.ChatLine{
			Info:      "[insert time]",
			InfoColor: escapecode.Yellow,
			Text:      "this command does not exists",
		})
	}

	handler()
}

func (controller *Controller) HandleResponse(response dto.Response) {
	// TODO: criar action name desconhecido
	if response.Err && response.Name == "unknown" { // TODO: verificar o erro da response rem o name unknow
		controller.handler.CUI().DrawNewLineForInternalError(response.Payload.(string))
		return
	}

	handler, err := controller.responseHandler(response.Name)
	if err != nil {
		controller.handler.CUI().DrawNewLineForInternalError(err.Error())
		return
	}

	handler(response)
}
