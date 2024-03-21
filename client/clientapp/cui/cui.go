package cui

import (
	"fmt"
	"strings"
	"time"

	tsize "github.com/kopoli/go-terminal-size"
)

var cliChatText = [10]string{
	`░█████╗░██╗░░░░░██╗  ░█████╗░██╗░░██╗░█████╗░████████╗`,
	`██╔══██╗██║░░░░░██║  ██╔══██╗██║░░██║██╔══██╗╚══██╔══╝`,
	`██║░░╚═╝██║░░░░░██║  ██║░░╚═╝███████║███████║░░░██║░░░`,
	`██║░░██╗██║░░░░░██║  ██║░░██╗██╔══██║██╔══██║░░░██║░░░`,
	`╚█████╔╝███████╗██║  ╚█████╔╝██║░░██║██║░░██║░░░██║░░░`,
	`░╚════╝░╚══════╝╚═╝  ░╚════╝░╚═╝░░╚═╝╚═╝░░╚═╝░░░╚═╝░░░`,
	`                                                      `,
	`             ┌───────────login───────────┐            `,
	`             │>                          │            `,
	`             └───────────────────────────┘            `,
}

const cliChatTextWidth = 54

type CUI struct {
	consoleHeight uint16
	consoleWidth  uint16

	sizeListener *tsize.SizeListener
}

func NewCUI() (*CUI, error) {
	var size, err = tsize.GetSize()
	if err != nil {
		return nil, tsize.ErrNotATerminal
	}

	listener, err := tsize.NewSizeListener()
	if err != nil {
		return nil, err
	}

	var cui = &CUI{
		consoleHeight: uint16(size.Height),
		consoleWidth:  uint16(size.Width),
		sizeListener:  listener,
	}

	go cui.listenToConsoleSize()

	return cui, nil
}

func (cui *CUI) DrawLoginInterface() {
	designer := newConsoleDesigner()
	designer.clearTerminal()

	var x int16 = 1
	var y int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.setColor(Blue).print()
	for _, text := range cliChatText {
		designer.setCursorCoordinates(coordinates{x: x, y: y})
		designer.setDrawing(text).print()
		x++
	}
	designer.resetColors()

	designer.
		setCursorColor(White).
		moveCursor(coordinates{
			x: 9, y: y + 16,
		})
}

func (cui *CUI) DrawLoginError(message string) {
	designer := newConsoleDesigner()
	var start int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.moveCursor(coordinates{x: 11, y: start + 15})
	designer.setDrawing(message).setColor(Red).print()
	designer.resetColors()

	designer.resetColors()
	designer.
		setCursorColor(White).
		moveCursor(coordinates{
			x: 9, y: start + 16,
		})
}

// func (cui *CUI) InitApp() {
// 	cui.drawConsoleUserInterface()
// }

func (cui *CUI) DrawConsoleUserInterface() {
	// TODO: do this after login
}

// TODO: testar
func (cui *CUI) DrawLoading(length uint, color escapeCode) error {
	designer := newConsoleDesigner()
	designer.setColor(color)

	l := int(length)
	divider := 6
	var counter uint
	for counter < length {
		time.Sleep(16 * time.Millisecond)

		currentWidth := counter / uint(divider)
		loadingString := fmt.Sprintf(
			"[%s%s]",
			strings.Repeat("=", int(currentWidth)),
			strings.Repeat(" ", (l/divider)-int(currentWidth)),
		)
		loadingString += fmt.Sprintf(" %d%%", counter)

		moveLeftCode, err := replaceInEscapeCode(moveLeft, map[string]int16{"n": 1000})
		if err != nil {
			return err
		}

		err = designer.setDrawing(moveLeftCode + loadingString).print()
		if err != nil {
			return err
		}

		counter++
	}

	designer.resetColors()

	return nil
}

func (cui *CUI) listenToConsoleSize() {
	for size := range cui.sizeListener.Change {
		cui.adaptNewConsoleSize(size)
	}
}

func (cui *CUI) adaptNewConsoleSize(size tsize.Size) {
	// TODO: do this
}
