package cui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/acarl005/stripansi"
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

type ConsoleInterface string

var Interfaces = [2]ConsoleInterface{"login", "chat"}

const (
	Login = iota
	Chat
)

type CUI struct {
	consoleHeight          uint16
	consoleWidth           uint16
	chatBoxHeight          uint16
	typingBoxHeight        uint16
	xCoordinateToType      int16
	xCoordinateToTypeLogin int16
	sizeListener           *tsize.SizeListener

	currentInterface ConsoleInterface
	chatLines        []*Line
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

	typingBoxHeight := uint16(10)
	var cui = &CUI{
		consoleHeight:          uint16(size.Height),
		consoleWidth:           uint16(size.Width),
		chatBoxHeight:          uint16(size.Height) - typingBoxHeight,
		xCoordinateToType:      int16(uint16(size.Height) - 9),
		xCoordinateToTypeLogin: 10,
		typingBoxHeight:        typingBoxHeight,
		sizeListener:           listener,
		currentInterface:       Interfaces[Login],
		chatLines:              make([]*Line, 0),
	}

	return cui, nil
}

func (cui *CUI) setCurrentInterface(currentInterface ConsoleInterface) {
	cui.currentInterface = currentInterface
}

func (cui *CUI) CurrentInterface() ConsoleInterface {
	return cui.currentInterface
}

func (cui *CUI) InitConsoleUserInterface() {
	cui.RenderLoginInterface()
	go cui.listenToConsoleSize()
}

func (cui *CUI) listenToConsoleSize() {
	for newSize := range cui.sizeListener.Change {
		cui.updateMeasures(newSize)
		cui.adaptNewConsoleSize()
	}
}

func (cui *CUI) updateMeasures(newSize tsize.Size) {
	consoleHeight := uint16(newSize.Height)
	cui.consoleWidth = uint16(newSize.Width)
	cui.consoleHeight = consoleHeight
	cui.chatBoxHeight = consoleHeight - cui.typingBoxHeight
	cui.xCoordinateToType = int16(consoleHeight - 9)
}

func (cui *CUI) adaptNewConsoleSize() {
	switch cui.currentInterface {
	case Interfaces[Login]:
		cui.RenderLoginInterface()
	case Interfaces[Chat]:
		cui.RenderChatInterface()
		cui.adaptChatLinesOnTerminal()
	}
}

func (cui *CUI) RenderLoginInterface() {
	defer cui.setCurrentInterface(Interfaces[Login])

	designer := newConsoleDesigner()
	designer.clearTerminal()

	var startOfCliChatText int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.setColor(escapecode.Blue).print()
	for currentLine, text := range cliChatText {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine + 2), y: startOfCliChatText})
		designer.setText(text).print()
	}
	designer.resetColors()

	startOfLoginBox := startOfCliChatText + 13
	designer.
		setCursorColor(escapecode.White).
		moveCursor(coordinates{
			x: cui.xCoordinateToTypeLogin, y: startOfLoginBox + 3,
		})
}

func (cui *CUI) DrawLoginInterfaceWithMessage(message string, color escapecode.ColorCode) {
	defer cui.setCurrentInterface(Interfaces[Login])

	cui.RenderLoginInterface()

	designer := newConsoleDesigner()
	startOfError := cui.xCoordinateToTypeLogin + 2

	startOfCliChatText := int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.moveCursor(coordinates{x: startOfError, y: startOfCliChatText + 15})
	designer.setText(message).setColor(color).printAndResetColors()

	startOfLoginBox := startOfCliChatText + 13
	designer.
		setCursorColor(escapecode.White).
		moveCursor(coordinates{
			x: cui.xCoordinateToTypeLogin, y: startOfLoginBox + 3,
		})
}

// TODO: exibir os comandos disponiveis
func (cui *CUI) RenderChatInterface() {
	defer cui.setCurrentInterface(Interfaces[Chat])

	designer := newConsoleDesigner()
	designer.clearTerminal()
	defer cui.moveCursorToTypeInChat(designer)

	designer.setColor(escapecode.White)
	chatBox := cui.chatBoxToSlice()
	for currentLine, lineText := range chatBox {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine + 1), y: 1})
		designer.setText(lineText).print()
	}
	designer.resetColors()

	designer.setColor(escapecode.Yellow)
	typingBox := cui.typingBoxToSlice()
	startOfDrawing := cui.consoleHeight - cui.typingBoxHeight
	for currentLine, lineText := range typingBox {
		currentLine += int(startOfDrawing)
		designer.setCursorCoordinates(coordinates{x: int16(currentLine), y: 1})
		designer.setText(lineText).print()
	}
}

func (cui *CUI) RedrawTypingBox() {
	designer := newConsoleDesigner()
	defer cui.moveCursorToTypeInChat(designer)

	designer.setColor(escapecode.Yellow)
	typingBox := cui.typingBoxToSlice()
	startOfDrawing := cui.consoleHeight - cui.typingBoxHeight
	for currentLine, line := range typingBox {
		currentLine += int(startOfDrawing)
		designer.setCursorCoordinates(coordinates{x: int16(currentLine), y: 1})
		designer.setText(line).print()
	}
}

func (cui *CUI) chatBoxToSlice() (chatBox []string) {
	chatBox = make([]string, cui.chatBoxHeight, cui.chatBoxHeight)
	for lineNumber := range cui.chatBoxHeight {
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

	return
}

func (cui *CUI) typingBoxToSlice() (typingBox []string) {
	typingBox = make([]string, cui.typingBoxHeight, cui.typingBoxHeight)
	for lineNumber := range cui.typingBoxHeight {
		var lineStr string
		switch lineNumber {
		case 0:
			amountOfSpaces := int(cui.consoleWidth) - 10
			lineStr = fmt.Sprintf(` ┌──TYPE%s┐ `, strings.Repeat(`─`, amountOfSpaces))
		case 1:
			amountOfSpaces := int(cui.consoleWidth) - 5
			lineStr = fmt.Sprintf(` │>%s│ `, strings.Repeat(` `, amountOfSpaces))
		case cui.typingBoxHeight - 1:
			amountOfSpaces := int(cui.consoleWidth) - 4
			lineStr = fmt.Sprintf(` └%s┘ `, strings.Repeat(`─`, amountOfSpaces))
		default:
			amountOfSpaces := int(cui.consoleWidth) - 4
			lineStr = fmt.Sprintf(` │%s│ `, strings.Repeat(` `, amountOfSpaces))
		}

		typingBox[lineNumber] = lineStr
	}

	return
}

func (cui *CUI) PrintLine(line *Line) {
	line = addDataToLine(line)

	switch cui.CurrentInterface() {
	case Interfaces[Login]:
		cui.DrawLoginInterfaceWithMessage((*line).Text, (*line).InfoColor)
	case Interfaces[Chat]:
		cui.addChatLine(line)
		cui.printChatLines()
	}
}

func (cui *CUI) PrintLineForInternalError(message string) {
	cui.PrintLine(&Line{
		Info:      "internal error:",
		InfoColor: escapecode.Red,
		Text:      message,
		TextColor: escapecode.Yellow,
	})
}

func (cui *CUI) addChatLine(line *Line) {
	cui.chatLines = append(cui.chatLines, line)
	cui.checkIfChatLinesExceededTheLimit()
}

func (cui *CUI) checkIfChatLinesExceededTheLimit() {
	numberOfVisibleLines := uint16(len(cui.chatLines))
	if numberOfVisibleLines > (cui.chatBoxHeight - 3) {
		diff := numberOfVisibleLines - (cui.chatBoxHeight - 3)
		cui.chatLines = cui.chatLines[diff:]
	}
}

func (cui *CUI) printChatLines() {
	designer := newConsoleDesigner()
	defer cui.moveCursorToTypeInChat(designer)

	designer.setColor(escapecode.White)
	for key, line := range cui.chatLines {
		info := newConsoleDesigner().
			setColor(line.InfoColor).
			setText(line.Info).
			toStringWithResetColors()

		text := newConsoleDesigner().
			setColor(line.TextColor).
			setText(line.Text).
			toStringWithResetColors()

		lineText := info + " " + text
		infoAndText := stripansi.Strip(line.Info + " " + line.Text)
		amountOfSpaces := (int(cui.consoleWidth) - 4) - len(infoAndText)
		if amountOfSpaces <= 0 {
			amountOfSpaces = 0
		}

		lineStr := fmt.Sprintf(` │%s%s│ `, lineText, strings.Repeat(" ", amountOfSpaces))
		designer.setCursorCoordinates(coordinates{x: int16(key + 2), y: 1})
		designer.eraseLine()
		designer.setText(lineStr).printAndResetColors()
	}
}

func (cui *CUI) moveCursorToTypeInChat(designer *consoleDesigner) {
	designer.resetColors()
	designer.
		setCursorColor(escapecode.White).
		moveCursor(coordinates{
			x: cui.xCoordinateToType,
			y: 5,
		})
}

func (cui *CUI) adaptChatLinesOnTerminal() {
	cui.checkIfChatLinesExceededTheLimit()
	cui.printChatLines()
}

func (cui *CUI) PrintLineAndExit(code uint8, l Line) {
	line := addDataToLine(&l)

	designer := newConsoleDesigner()
	designer.clearTerminal()

	info := newConsoleDesigner().
		setColor(line.InfoColor).
		setText(line.Info).
		toStringWithResetColors()

	text := newConsoleDesigner().
		setColor(line.TextColor).
		setText(line.Text).
		toStringWithResetColors()

	lineStr := info + " " + text
	designer.
		setCursorCoordinates(coordinates{
			x: 1, y: 1,
		}).
		setText(lineStr + "\n").
		printAndResetColors()
	os.Exit(int(code))
}

func (cui *CUI) DrawLoading(length uint, color escapecode.ColorCode) error {
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

		moveLeftCode, err := escapecode.ReplaceInEscapeCode(escapecode.MoveLeft, map[string]int16{"n": 1000})
		if err != nil {
			return err
		}

		err = designer.setText(moveLeftCode + loadingString).print()
		if err != nil {
			return err
		}

		counter++
	}

	designer.resetColors()

	return nil
}
