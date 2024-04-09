package escapecode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type replaceInt16 map[string]int16
type replaceString map[string]string

func TestReplaeInEscapeCode(t *testing.T) {
	type myTests[T string | int16] struct {
		code           escapeCode
		replaces       map[string]T
		expectedReturn string
	}

	tests1 := [...]myTests[int16]{
		{
			code:           MoveCursor,
			replaces:       replaceInt16{"lines": 6, "columns": 12},
			expectedReturn: "\x1b[6;12H",
		},
		{
			code:           MoveLeft,
			replaces:       replaceInt16{"n": 9},
			expectedReturn: "\x1b[9D",
		},
		{
			code:           MoveRight,
			replaces:       replaceInt16{"n": 3},
			expectedReturn: "\x1b[3C",
		},
	}

	tests2 := [...]myTests[string]{
		{
			code:           escapeCode("{name}|{lastname}"),
			replaces:       replaceString{"name": "Fabricio", "lastname": "Pereira Alves"},
			expectedReturn: "Fabricio|Pereira Alves",
		},
		{
			code:           escapeCode("{dev} is {type}-end"),
			replaces:       replaceString{"dev": "Fabricio", "type": "back"},
			expectedReturn: "Fabricio is back-end",
		},
		{
			code:           escapeCode("{word1} {word2} {word3}"),
			replaces:       replaceString{"word1": "i'm", "word2": "the", "word3": "best"},
			expectedReturn: "i'm the best",
		},
	}

	for _, data := range tests1 {
		code, err := ReplaceInEscapeCode(data.code, data.replaces)
		assert.Nil(t, err)
		assert.Equal(t, data.expectedReturn, code)
	}

	for _, data := range tests2 {
		code, err := ReplaceInEscapeCode(data.code, data.replaces)
		assert.Nil(t, err)
		assert.Equal(t, data.expectedReturn, code)
	}
}

// TODO: fazer o teste com dados errados
