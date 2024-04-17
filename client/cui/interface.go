package cui

import "github.com/Fabriciope/cli_chat/pkg/escapecode"

type CUIInterface interface {
	InitConsoleUserInterface()
	RenderLoginInterface()
	RenderChatInterface()
	CurrentInterface() ConsoleInterface
	RedrawTypingBox()
	PrintLine(*Line)
	PrintMessageInLoginInterface(string, escapecode.ColorCode)
	PrintLineForInternalError(string) // TODO: mandar os erros internos para um arquivo de log e criar um package logger para fazer isso
	PrintLineAndExit(uint8, Line)
}
