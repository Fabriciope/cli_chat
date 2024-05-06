package controller

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/Fabriciope/cli_chat/client/controller/inputhandler"
	"github.com/Fabriciope/cli_chat/client/controller/responsehandler"
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

type CommandHandler func(string)
type CommandsHandlersMap map[string]CommandHandler
type ResponseHandler func(dto.Response)
type ResponsesHandlersMap map[string]ResponseHandler

type Controller struct {
	commandsHandlers  CommandsHandlersMap
	responsesHandlers ResponsesHandlersMap
	inputHandler      *inputhandler.InputHandler
	responseHandler   *responsehandler.ResponseHandler
	cui               cui.CUIInterface
	userLoggedIn      *bool
}

func NewController(conn *net.TCPConn, cui cui.CUIInterface, loggedIn *bool) *Controller {
	inputHandler := inputhandler.NewInputHandler(conn, cui, loggedIn)
	responseHandler := responsehandler.NewResponseHandler(cui, loggedIn)
	controller := &Controller{
		commandsHandlers:  make(CommandsHandlersMap),
		responsesHandlers: make(ResponsesHandlersMap),
		inputHandler:      inputHandler,
		responseHandler:   responseHandler,
		cui:               cui,
		userLoggedIn:      loggedIn,
	}

	controller.setHandlerForEachCommand()
	controller.setHandlerForEachResponse()

	return controller
}

func (controller *Controller) setHandlerForEachCommand() {
	availableCommand := func(command int) string {
		return inputhandler.AvailableCommands[command]
	}

	controller.commandsHandlers = CommandsHandlersMap{
		":login":                                     controller.inputHandler.Login,
		availableCommand(inputhandler.Logout):        controller.inputHandler.Logout,
		availableCommand(inputhandler.Users):         controller.inputHandler.GetUsers,
		availableCommand(inputhandler.NumberOfUsers): controller.inputHandler.GetNumberOfUsers,
		availableCommand(inputhandler.Commands):      controller.inputHandler.GetAvailableCommands,
		availableCommand(inputhandler.H):             controller.inputHandler.GetAvailableCommands,
		availableCommand(inputhandler.Help):          controller.inputHandler.GetAvailableCommands,
		availableCommand(inputhandler.Q):             controller.inputHandler.Logout,
		availableCommand(inputhandler.Quit):          controller.inputHandler.Logout,
	}
}

func (controller *Controller) setHandlerForEachResponse() {
	controller.responsesHandlers = ResponsesHandlersMap{
		dto.LoginActionName:              controller.responseHandler.Login,
		dto.NewClientNotificationName:    controller.responseHandler.NewClient,
		dto.SendMessageActionName:        controller.responseHandler.SendMessageInChat,
		dto.NewMessageNotificationName:   controller.responseHandler.NewMessageReceived,
		dto.LogoutActionName:             controller.responseHandler.Logout,
		dto.ClientDisconnectedActionName: controller.responseHandler.UserDisconnected,
		dto.GetUsersActionName:           controller.responseHandler.GetUsers,
		dto.GetUsersCountActionName:      controller.responseHandler.GetUsersCount,
	}
}

func (controller *Controller) getCommandHandler(command string) (CommandHandler, error) {
	handler, exists := (*controller).commandsHandlers[command]
	if exists == false {
		return nil, errors.New("input handler for this command does not exist")
	}

	return handler, nil
}

func (controller *Controller) getResponseHandler(actionName string) (ResponseHandler, error) {
	handler, exists := (*controller).responsesHandlers[actionName]
	if exists == false {
		return nil, errors.New("response handler for this action name does not exists")
	}

	return handler, nil
}

func (controller *Controller) HandleInput(input string) {
	input = strings.Trim(input, " ")

	if strings.HasPrefix(input, ":") {
		inputSplitted := strings.Split(input, " ")
		handler, err := controller.getCommandHandler(inputSplitted[0])
		if err != nil {
			controller.cui.PrintLine(&cui.Line{
				Info:      "error:",
				Text:      fmt.Sprintf("%s command does not exist", inputSplitted[0]),
				TextColor: escapecode.Yellow,
			})

			return
		}

		handler(strings.Join(inputSplitted[1:], " ")) //ou usar o trimPrefix com inputSplitted[0]
		return
	}

	controller.inputHandler.SendMessageInChat(input)
}

func (controller *Controller) HandleResponse(response dto.Response) {
	if response.Err && response.Name == dto.UnknownActionName {
		controller.cui.PrintLineForInternalError(response.Payload.(string))
		return
	}

	handler, err := controller.getResponseHandler(response.Name)
	if err != nil {
		controller.cui.PrintLineForInternalError(err.Error())
		return
	}

	handler(response)
}
