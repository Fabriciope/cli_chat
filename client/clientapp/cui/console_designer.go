package cui

import (
	"errors"
	"fmt"
)

// TODO: criar variavies de erro

type consoleDesigner struct {
	drawingColor escapeCode
	cursorColor  escapeCode

	cursorCoordinates coordinates

	drawing string
}

func newConsoleDesigner() *consoleDesigner {
	return &consoleDesigner{
		drawingColor:      DefaultColor,
		cursorColor:       DefaultColor,
		cursorCoordinates: coordinates{x: 0, y: 0},
	}
}

func (designer *consoleDesigner) setColor(drawingColor escapeCode) *consoleDesigner {
	designer.drawingColor = drawingColor

	return designer
}

func (designer *consoleDesigner) setCursorColor(cursorColor escapeCode) *consoleDesigner {
	designer.cursorColor = cursorColor

	return designer
}

func (designer *consoleDesigner) setCursorCoordinates(coordinates coordinates) *consoleDesigner {
	designer.cursorCoordinates = coordinates
	return designer
}

func (designer *consoleDesigner) moveCursor(coordinates coordinates) {
	replaces := map[string]int16{
		"lines":   coordinates.x,
		"columns": coordinates.y,
	}
	fmt.Print(designer.cursorColor)
	code, _ := replaceInEscapeCode(moveCursor, replaces)
	fmt.Print(code)
}

func (designer *consoleDesigner) setDrawing(drawing string) *consoleDesigner {
	designer.drawing = drawing

	return designer
}

// TODO: depois de prinntar o desenho no console resetar as cores
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
		code, _ := replaceInEscapeCode(moveCursor, replaces)
		fmt.Print(code)
	}

	fmt.Print(designer.drawingColor, designer.drawing)
	return nil
}

func (designer *consoleDesigner) printAndResetColors() error {
	return nil
}

func (designer *consoleDesigner) toString() string {
	return ""
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
		moveCursorCode, _ = replaceInEscapeCode(moveCursor, replaces)
	}

	return fmt.Sprintf(
		"%s%s%s%s",
		designer.drawingColor,
		designer.drawing,
		moveCursorCode,
		reset,
	)
}

func (designer *consoleDesigner) clearTerminal() {
	fmt.Print(clearScreen)
}

func (designer *consoleDesigner) resetColors() {
	designer.drawingColor = DefaultColor
	designer.cursorColor = DefaultColor
	fmt.Print(reset)
}
