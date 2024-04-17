package escapecode

import (
	"errors"
	"fmt"
	"strings"
)

func ReplaceInEscapeCode[T string | int16](code escapeCode, replaces map[string]T) (string, error) {
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

func TextToBold(text string) string {
	return string(Bold) + text + string(Reset)
}
