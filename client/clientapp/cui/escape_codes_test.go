package cui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEscapeCodeByReplacingTheNParameter(t *testing.T) {
	var tests = []struct {
		scapeCode       escapeCode
		numberToReplace int16
		expectedReturn  string
	}{
		{moveLeft, 5, "\033[5D"},
		{moveRight, 3, "\033[3C"},
		{moveUp, 8, "\033[8A"},
		{moveDown, 2, "\033[2B"},
	}

	for _, data := range tests {
		code := getEscapeCode(data.scapeCode, data.numberToReplace)
		assert.Equal(t, data.expectedReturn, code)
	}
}
