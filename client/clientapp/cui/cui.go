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

var interfaces = [2]string{"login", "chat"}

const (
	Login = iota
	Chat
)

type CUI struct {
	consoleHeight uint16
	consoleWidth  uint16

	sizeListener *tsize.SizeListener

	currentInterface string

	logged chan bool
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
		logged:        make(chan bool),
	}

	return cui, nil
}

func (cui *CUI) SetLoggedAs(logged bool) {
	cui.logged <- logged
}

func (cui *CUI) InitApp() {
	cui.drawLoginInterface()
	go cui.listenToConsoleSize()

	if logged := <-cui.logged; logged {
		cui.currentInterface = interfaces[Chat]
		cui.drawConsoleUserInterface()
	}
}

func (cui *CUI) listenToConsoleSize() {
	for newSize := range cui.sizeListener.Change {
		cui.updateConsoleSize(newSize)
		cui.adaptNewConsoleSize()
	}
}

func (cui *CUI) updateConsoleSize(newSize tsize.Size) {
	cui.consoleWidth = uint16(newSize.Width)
	cui.consoleHeight = uint16(newSize.Height)
	fmt.Print(cui.consoleWidth, "-", cui.consoleHeight)
}

func (cui *CUI) adaptNewConsoleSize() {
	switch cui.currentInterface {
	case "login":
		cui.drawLoginInterface()
	case "chat":
		cui.drawConsoleUserInterface()
	}
}

//func (cui *CUI) adaptLoginInterfaceToNewSize() {
//
//}
//
//func (cui *CUI) adaptChatInterfaceToNewSize() {
//    // TODO: quando for adptar cortar a partir da direita
//}

func (cui *CUI) drawLoginInterface() {
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

	cui.currentInterface = interfaces[Login]
}

func (cui *CUI) DrawLoginError(message string) {
	designer := newConsoleDesigner()
	var start int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.moveCursor(coordinates{x: 11, y: start + 15})
	designer.setDrawing(message).setColor(Red).print()
	designer.resetColors()

	// TODO: limpar login anterior

	designer.resetColors()
	designer.
		setCursorColor(White).
		moveCursor(coordinates{
			x: 9, y: start + 16,
		})
}

func (cui *CUI) drawConsoleUserInterface() {
	// TODO: do this after login
	fmt.Print("start chat interface")
}

// TODO: testar
func (cui *CUI) drawLoading(length uint, color escapeCode) error {
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
