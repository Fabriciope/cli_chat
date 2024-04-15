package cui

// TODO: tirar o metodo de desenhar DrawLoginError e DrawNewLineInChat e fazer so uma e deixar que o cui entenda se o usuario esta logado ou nao
// TODO: trocar os nomes que printam mensagem na tela de draw para print
type CUIInterface interface {
	InitConsoleUserInterface()
	SetLoggedAs(bool) // TODO: trocar nome do metodo para setUserLoggedInAs
	DrawNewLineInChat(*ChatLine)
	RedrawTypingBox()
	DrawLoginError(string)
	DrawLineAndExit(uint8, ChatLine)
	// TODO: mandar os erros internos para um arquivo de log e criar um package logger para fazer isso
	// TODO: pensar em trocar o campo logged para *bool
	DrawNewLineForInternalError(string)
}
