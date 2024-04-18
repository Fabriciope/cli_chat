package escapecode

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	invalidCodeErr     = errors.New("invalid code")
	invalidReplacesErr = errors.New("invalid replaces")
)

func ReplaceInEscapeCode[T string | int16](c escapeCode, replaces map[string]T) (string, error) {
	code := string(c)

	regexPattern := regexp.MustCompile(`(\{ *[a-zA-Z0-9- ]* *})`)
	allSubstringsMatches := regexPattern.FindAllString(code, -1)
	if allSubstringsMatches == nil || strings.Contains(code, "{}") {
		return "", invalidCodeErr
	}
	var newCode string

	firstLoop := true
	for _, oldKey := range allSubstringsMatches {
		toReplace, ok := replaces[strings.TrimSuffix(strings.TrimPrefix(oldKey, "{"), "}")]
		if !ok {
			return "", invalidReplacesErr
		}

		if firstLoop {
			newCode = strings.Replace(code, oldKey, fmt.Sprint(toReplace), 1)
			firstLoop = false
		} else {
			newCode = strings.Replace(newCode, oldKey, fmt.Sprint(toReplace), 1)
		}
	}

	return newCode, nil
}

func TextToBold(text string) string {
	return string(Bold) + text + string(Reset)
}
