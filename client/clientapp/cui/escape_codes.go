package cui

import (
	"errors"
	"fmt"
	"strings"
)

type escapeCode string

const (
	reset escapeCode = "\x1b[0m"

	clearScreen   escapeCode = "\x1b[2J"
	moveHome      escapeCode = "\x1b[H"
	moveCursor    escapeCode = "\x1b[{lines};{columns}H"
	moveUp        escapeCode = "\x1b[{n}A"
	moveDown      escapeCode = "\x1b[{n}B"
	moveRight     escapeCode = "\x1b[{n}C"
	moveLeft      escapeCode = "\x1b[{n}D"
	moveNextLine  escapeCode = "\x1b[E"
	movePrevLine  escapeCode = "\x1b[F"
	saveCursor    escapeCode = "\x1b[s"
	restoreCursor escapeCode = "\x1b[u"
	hideCursor    escapeCode = "\x1b[?25l"
	showCursor    escapeCode = "\x1b[?25h"

	bold       escapeCode = "\x1b[1m"
	dim        escapeCode = "\x1b[2m"
	italic     escapeCode = "\x1b[3m"
	underline  escapeCode = "\x1b[4m"
	blink      escapeCode = "\x1b[5m"
	reverse    escapeCode = "\x1b[7m"
	hidden     escapeCode = "\x1b[8m"
	crossedOut escapeCode = "\x1b[9m"

	DefaultColor  escapeCode = "\x1b[39m"
	Black         escapeCode = "\x1b[30m"
	Red           escapeCode = "\x1b[31m"
	Green         escapeCode = "\x1b[32m"
	Yellow        escapeCode = "\x1b[33m"
	Blue          escapeCode = "\x1b[34m"
	Magenta       escapeCode = "\x1b[35m"
	Cyan          escapeCode = "\x1b[36m"
	LightGray     escapeCode = "\x1b[37m"
	DarkGray      escapeCode = "\x1b[90m"
	BrightRed     escapeCode = "\x1b[91m"
	BrightGreen   escapeCode = "\x1b[92m"
	BrightYellow  escapeCode = "\x1b[93m"
	BrightBlue    escapeCode = "\x1b[94m"
	BrightMagenta escapeCode = "\x1b[95m"
	BrightCyan    escapeCode = "\x1b[96m"
	White         escapeCode = "\x1b[97m"

	DefaultBackground escapeCode = "\x1b[49m"
	BgBlack           escapeCode = "\x1b[40m"
	BgRed             escapeCode = "\x1b[41m"
	BgGreen           escapeCode = "\x1b[42m"
	BgYellow          escapeCode = "\x1b[43m"
	BgBlue            escapeCode = "\x1b[44m"
	BgMagenta         escapeCode = "\x1b[45m"
	BgCyan            escapeCode = "\x1b[46m"
	BgLightGray       escapeCode = "\x1b[47m"
	BgDarkGray        escapeCode = "\x1b[100m"
	BgBrightRed       escapeCode = "\x1b[101m"
	BgBrightGreen     escapeCode = "\x1b[102m"
	BgBrightYellow    escapeCode = "\x1b[103m"
	BgBrightBlue      escapeCode = "\x1b[104m"
	BgBrightMagenta   escapeCode = "\x1b[105m"
	BgBrightCyan      escapeCode = "\x1b[106m"
	BgWhite           escapeCode = "\x1b[107m"
)

func replaceInEscapeCode[T string | int16](code escapeCode, replaces map[string]T) (string, error) {
	c := string(code)
	if !strings.Contains(c, "{") || !strings.Contains(c, "}") {
		return "", errors.New("invalid code")
	}

	var newCode string

	first := true
	for key, replace := range replaces {
		key = fmt.Sprintf("{%s}", key)
		if !strings.Contains(c, key) {
			return "", errors.New("invalid replaces")
		}

		if first {
			newCode = strings.Replace(c, key, fmt.Sprint(replace), 1)
			first = false
		} else {
			newCode = strings.Replace(newCode, key, fmt.Sprint(replace), 1)
		}
	}

	return newCode, nil
}
