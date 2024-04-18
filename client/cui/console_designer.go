package cui

import (
	"errors"
	"fmt"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

type consoleDesigner struct {
	textColor   escapecode.ColorCode
	cursorColor escapecode.ColorCode

	cursorCoordinates coordinates

	text string
}

func newConsoleDesigner() *consoleDesigner {
	return &consoleDesigner{
		textColor:         escapecode.DefaultColor,
		cursorColor:       escapecode.DefaultColor,
		cursorCoordinates: coordinates{x: 0, y: 0},
	}
}

func (designer *consoleDesigner) setColor(drawingColor escapecode.ColorCode) *consoleDesigner {
	if drawingColor != "" {
		designer.textColor = drawingColor
	}

	return designer
}

func (designer *consoleDesigner) setCursorColor(cursorColor escapecode.ColorCode) *consoleDesigner {
	if cursorColor != "" {
		designer.cursorColor = cursorColor
	}

	return designer
}

func (designer *consoleDesigner) setCursorCoordinates(coordinates coordinates) *consoleDesigner {
	designer.cursorCoordinates = coordinates
	return designer
}

func (designer *consoleDesigner) moveCursor(coordinates coordinates) *consoleDesigner {
	replaces := map[string]int16{
		"lines":   coordinates.x,
		"columns": coordinates.y,
	}
	fmt.Print(designer.cursorColor)
	code, _ := escapecode.ReplaceInEscapeCode(escapecode.MoveCursor, replaces)
	fmt.Print(code)

	return designer
}

func (designer *consoleDesigner) setText(drawing string) *consoleDesigner {
	designer.text = drawing

	return designer
}

func (designer *consoleDesigner) print() error {
	if designer.text == "" {
		return errors.New("make a drawing first")
	}

	x := designer.cursorCoordinates.x
	y := designer.cursorCoordinates.y
	if x != 0 && y != 0 {
		replaces := map[string]int16{
			"lines":   x,
			"columns": y,
		}
		code, _ := escapecode.ReplaceInEscapeCode(escapecode.MoveCursor, replaces)
		fmt.Print(code)
	}

	fmt.Print(designer.textColor, designer.text)
	return nil
}

func (designer *consoleDesigner) printAndResetColors() error {
	if designer.text == "" {
		return errors.New("make a drawing first")
	}

	x := designer.cursorCoordinates.x
	y := designer.cursorCoordinates.y
	if x != 0 && y != 0 {
		replaces := map[string]int16{
			"lines":   x,
			"columns": y,
		}
		code, _ := escapecode.ReplaceInEscapeCode(escapecode.MoveCursor, replaces)
		fmt.Print(code)
	}

	fmt.Print(designer.textColor, designer.text, escapecode.Reset)
	return nil
}

func (designer *consoleDesigner) toString() string {
	var moveCursorCode string

	x := designer.cursorCoordinates.x
	y := designer.cursorCoordinates.y
	if x != 0 && y != 0 {
		replaces := map[string]int16{
			"lines":   x,
			"columns": y,
		}
		moveCursorCode, _ = escapecode.ReplaceInEscapeCode(escapecode.MoveCursor, replaces)
	}

	return fmt.Sprintf(
		"%s%s%s",
		designer.textColor,
		designer.text,
		moveCursorCode,
	)
}

func (designer *consoleDesigner) toStringWithResetColors() string {
	var moveCursorCode string

	x := designer.cursorCoordinates.x
	y := designer.cursorCoordinates.y
	if x != 0 && y != 0 {
		replaces := map[string]int16{
			"lines":   x,
			"columns": y,
		}
		moveCursorCode, _ = escapecode.ReplaceInEscapeCode(escapecode.MoveCursor, replaces)
	}

	return fmt.Sprintf(
		"%s%s%s%s",
		moveCursorCode,
		designer.textColor,
		designer.text,
		escapecode.Reset,
	)
}

func (designer *consoleDesigner) clearTerminal() {
	fmt.Print(escapecode.ClearScreen)
}

func (designer *consoleDesigner) resetColors() {
	designer.textColor = escapecode.DefaultColor
	designer.cursorColor = escapecode.DefaultColor
	fmt.Print(escapecode.Reset)
}

func (designer *consoleDesigner) eraseLine() {
	var moveCursorCode string

	if x := designer.cursorCoordinates.x; x != 0 {
		replaces := map[string]int16{
			"lines":   x,
			"columns": 1,
		}
		moveCursorCode, _ = escapecode.ReplaceInEscapeCode(escapecode.MoveCursor, replaces)
	}

	fmt.Print(moveCursorCode, escapecode.EraseLine)
}

func (designer *consoleDesigner) eraseLineWithXCoordinates(x int16) {
	var moveCursorCode string

	replaces := map[string]int16{
		"lines":   x,
		"columns": 1,
	}
	moveCursorCode, _ = escapecode.ReplaceInEscapeCode(escapecode.MoveCursor, replaces)

	fmt.Print(moveCursorCode, escapecode.EraseLine)
}
