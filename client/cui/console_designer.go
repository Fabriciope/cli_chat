package cui

import (
	"errors"
	"fmt"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

// TODO: criar variavies de erro

type consoleDesigner struct {
	drawingColor escapecode.ColorCode
	cursorColor  escapecode.ColorCode

	cursorCoordinates coordinates

	drawing string
}

func newConsoleDesigner() *consoleDesigner {
	return &consoleDesigner{
		drawingColor:      escapecode.DefaultColor,
		cursorColor:       escapecode.DefaultColor,
		cursorCoordinates: coordinates{x: 0, y: 0},
	}
}

func (designer *consoleDesigner) setColor(drawingColor escapecode.ColorCode) *consoleDesigner {
	designer.drawingColor = drawingColor

	return designer
}

func (designer *consoleDesigner) setCursorColor(cursorColor escapecode.ColorCode) *consoleDesigner {
	designer.cursorColor = cursorColor

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

func (designer *consoleDesigner) setDrawing(drawing string) *consoleDesigner {
	designer.drawing = drawing

	return designer
}

func (designer *consoleDesigner) print() error {
	if designer.drawing == "" {
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

	fmt.Print(designer.drawingColor, designer.drawing)
	return nil
}

func (designer *consoleDesigner) printAndResetColors() error {
	if designer.drawing == "" {
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

	fmt.Print(designer.drawingColor, designer.drawing, escapecode.Reset)
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
		designer.drawingColor,
		designer.drawing,
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
		designer.drawingColor,
		designer.drawing,
		escapecode.Reset,
	)
}

func (designer *consoleDesigner) clearTerminal() {
	fmt.Print(escapecode.ClearScreen)
}

func (designer *consoleDesigner) resetColors() {
	designer.drawingColor = escapecode.DefaultColor
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

// TODO: eraseCurrentLine()
