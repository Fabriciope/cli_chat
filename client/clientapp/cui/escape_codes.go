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
