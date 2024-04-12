package interfaces

import "github.com/Fabriciope/cli_chat/client/cui"

type CUI interface {
	InitConsoleUserInterface()
	SetLoggedAs(bool)
	DrawNewLineInChat(*cui.ChatLine)
	RedrawTypingBox()
	DrawLoginError(string)
	DrawLineAndExit(uint8, cui.ChatLine)
	// TODO: mandar os erros internos para um arquivo de log e criar um package logger para fazer isso
	DrawNewLineForInternalError(string)
}
