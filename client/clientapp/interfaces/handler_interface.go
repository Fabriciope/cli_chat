package interfaces

import "github.com/Fabriciope/cli_chat/shared"

type CommandHandler func() error
type CommandsHandlersMap map[string]CommandHandler
type ResponseHandler func(shared.Response)
type ResponsesHandlersMap map[string]ResponseHandler

//type Handler interface {
//	CommandHandler(string) CommandHandler
//	ResponseHandler(string) ResponseHandler
//}
