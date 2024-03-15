package cui

import (
	"strconv"
	"strings"
)

type escapeCode string

const (
	reset escapeCode = "\033[0m"

	clearScreen   escapeCode = "\033[2J"
	moveHome      escapeCode = "\033[H"
	moveCursor    escapeCode = "\033[{lines};{columns}H"
	moveUp        escapeCode = "\033[{n}A"
	moveDown      escapeCode = "\033[{n}B"
	moveRight     escapeCode = "\033[{n}C"
	moveLeft      escapeCode = "\033[{n}D"
	moveNextLine  escapeCode = "\033[E"
	movePrevLine  escapeCode = "\033[F"
	saveCursor    escapeCode = "\033[s"
	restoreCursor escapeCode = "\033[u"
	hideCursor    escapeCode = "\033[?25l"
	showCursor    escapeCode = "\033[?25h"

	bold       escapeCode = "\033[1m"
	dim        escapeCode = "\033[2m"
	italic     escapeCode = "\033[3m"
	underline  escapeCode = "\033[4m"
	blink      escapeCode = "\033[5m"
	reverse    escapeCode = "\033[7m"
	hidden     escapeCode = "\033[8m"
	crossedOut escapeCode = "\033[9m"

	defaultColor  escapeCode = "\033[39m"
	black         escapeCode = "\033[30m"
	red           escapeCode = "\033[31m"
	green         escapeCode = "\033[32m"
	yellow        escapeCode = "\033[33m"
	blue          escapeCode = "\033[34m"
	magenta       escapeCode = "\033[35m"
	cyan          escapeCode = "\033[36m"
	lightGray     escapeCode = "\033[37m"
	darkGray      escapeCode = "\033[90m"
	brightRed     escapeCode = "\033[91m"
	brightGreen   escapeCode = "\033[92m"
	brightYellow  escapeCode = "\033[93m"
	brightBlue    escapeCode = "\033[94m"
	brightMagenta escapeCode = "\033[95m"
	brightCyan    escapeCode = "\033[96m"
	white         escapeCode = "\033[97m"

	defaultBackground escapeCode = "\033[49m"
	bgBlack           escapeCode = "\033[40m"
	bgRed             escapeCode = "\033[41m"
	bgGreen           escapeCode = "\033[42m"
	bgYellow          escapeCode = "\033[43m"
	bgBlue            escapeCode = "\033[44m"
	bgMagenta         escapeCode = "\033[45m"
	bgCyan            escapeCode = "\033[46m"
	bgLightGray       escapeCode = "\033[47m"
	bgDarkGray        escapeCode = "\033[100m"
	bgBrightRed       escapeCode = "\033[101m"
	bgBrightGreen     escapeCode = "\033[102m"
	bgBrightYellow    escapeCode = "\033[103m"
	bgBrightBlue      escapeCode = "\033[104m"
	bgBrightMagenta   escapeCode = "\033[105m"
	bgBrightCyan      escapeCode = "\033[106m"
	bgWhite           escapeCode = "\033[107m"
)

func getEscapeCode(code escapeCode, numberToReplace int16) string {
	var firstIndex int = strings.Index(string(code), "{")
	var lastIndex int = strings.LastIndex(string(code), "}")

	var number escapeCode = escapeCode(strconv.Itoa(int(numberToReplace)))

	return string(code[:firstIndex] + number + code[lastIndex+1:])
}
