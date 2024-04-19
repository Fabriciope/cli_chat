package cui

import "github.com/Fabriciope/cli_chat/pkg/escapecode"

type CUIInterface interface {
	InitConsoleUserInterface()
	RenderLoginInterface()
	RenderChatInterface()
	CurrentInterface() ConsoleInterface
	RedrawTypingBox()
	PrintLine(*Line)
	DrawLoginInterfaceWithMessage(string, escapecode.ColorCode)
	PrintLineForInternalError(string)
	PrintLineAndExit(uint8, Line)
}
