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
		cui.drawChatInterface()
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
}

func (cui *CUI) adaptNewConsoleSize() {
	switch cui.currentInterface {
	case "login":
		cui.drawLoginInterface()
	case "chat":
		cui.drawChatInterface()
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

	var y int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.setColor(Blue).print()
	for currentLine, text := range cliChatText {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine + 2), y: y})
		designer.setDrawing(text).print()
	}
	designer.resetColors()

	designer.
		setCursorColor(White).
		moveCursor(coordinates{
			x: 10, y: y + 16,
		})

	cui.currentInterface = interfaces[Login]
}

func (cui *CUI) DrawLoginError(message string) {
	designer := newConsoleDesigner()
	var start int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.moveCursor(coordinates{x: 11, y: start + 15})
	designer.setDrawing(message).setColor(Red).print()
	designer.resetColors()

	// TODO: limpar login e erro anterior

	designer.resetColors()
	designer.
		setCursorColor(White).
		moveCursor(coordinates{
			x: 9, y: start + 16,
		})
}

func (cui *CUI) drawChatInterface() {
	typingBoxHeight := 10
	typingBox := make([]string, typingBoxHeight, typingBoxHeight)
	for lineNumber := range typingBoxHeight {
		var lineStr string
		switch lineNumber {
		case 0:
			amountOfSpaces := int(cui.consoleWidth) - 10
			lineStr = fmt.Sprintf(` ┌──type%s┐ `, strings.Repeat(`─`, amountOfSpaces))
		case 1:
			amountOfSpaces := int(cui.consoleWidth) - 5
			lineStr = fmt.Sprintf(` └%s┘ `, strings.Repeat(`─`, amountOfSpaces))
			lineStr = fmt.Sprintf(` │>%s│ `, strings.Repeat(` `, amountOfSpaces))
		case typingBoxHeight - 1:
			amountOfSpaces := int(cui.consoleWidth) - 4
			lineStr = fmt.Sprintf(` └%s┘ `, strings.Repeat(`─`, amountOfSpaces))
		default:
			amountOfSpaces := int(cui.consoleWidth) - 4
			lineStr = fmt.Sprintf(` │%s│ `, strings.Repeat(` `, amountOfSpaces))
		}

		typingBox = append(typingBox, lineStr)
	}

	chatBoxHeight := int(cui.consoleHeight - uint16(typingBoxHeight))
	chatBox := make([]string, chatBoxHeight, chatBoxHeight)
	for lineNumber := range chatBoxHeight {
		var lineStr string

		if lineNumber == 0 {
			amountOfSpaces := int(cui.consoleWidth) - 10
			lineStr = fmt.Sprintf(` ┌──CHAT%s┐ `, strings.Repeat(`─`, amountOfSpaces))
			chatBox[lineNumber] = lineStr
			continue
		}

		amountOfSpaces := int(cui.consoleWidth) - 4
		lineStr = fmt.Sprintf(` │%s│ `, strings.Repeat(` `, amountOfSpaces))
		chatBox[lineNumber] = lineStr
	}

	designer := newConsoleDesigner()
	designer.clearTerminal()

	designer.setColor(White)
	for currentLine, lineText := range chatBox {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine + 1), y: 1})
		designer.setDrawing(lineText).print()
	}
	designer.resetColors()

	designer.setColor(Yellow)
	for currentLine, line := range typingBox {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine), y: 0})
		designer.setDrawing(line).print()
	}
	designer.resetColors()

	designer.
		setCursorColor(White).
		moveCursor(coordinates{
			x: int16(cui.consoleHeight - 8),
			y: 5,
		})
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
