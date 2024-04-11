package cui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
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
	login = iota
	chat
)

type CUI struct {
	consoleHeight          uint16
	consoleWidth           uint16
	chatBoxHeight          uint16
	typingBoxHeight        uint16
	xCoordinateToType      int16
	xCoordinateToTypeLogin int16
	sizeListener           *tsize.SizeListener

	currentInterface string
	logged           chan bool
	chatLines        []*ChatLine
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
		currentInterface:       interfaces[login],
		logged:                 make(chan bool),
		chatLines:              make([]*ChatLine, 0),
	}

	return cui, nil
}

func (cui *CUI) SetLoggedAs(logged bool) {
	cui.logged <- logged
}

func (cui *CUI) InitConsoleUserInterface() {
	cui.drawLoginInterface()
	go cui.listenToConsoleSize()

	if logged := <-cui.logged; logged {
		cui.currentInterface = interfaces[chat]
		cui.drawChatInterface()
	}
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
	case "login":
		cui.drawLoginInterface()
	case "chat":
		cui.drawChatInterface()
		cui.adaptChatLinesOnTerminal()
	}
}

func (cui *CUI) drawLoginInterface() {
	designer := newConsoleDesigner()
	designer.clearTerminal()

	var startOfCliChatText int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.setColor(escapecode.Blue).print()
	for currentLine, text := range cliChatText {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine + 2), y: startOfCliChatText})
		designer.setDrawing(text).print()
	}
	designer.resetColors()

	startOfLoginBox := startOfCliChatText + 13
	designer.
		setCursorColor(escapecode.White).
		moveCursor(coordinates{
			x: cui.xCoordinateToTypeLogin, y: startOfLoginBox + 3,
		})

	cui.currentInterface = interfaces[login]
}

func (cui *CUI) DrawLoginError(message string) {
	designer := newConsoleDesigner()
	startOfError := cui.xCoordinateToTypeLogin + 2

	// clear old error message
	linesToClear := int16(cui.consoleHeight) - startOfError
	for cLine := startOfError; cLine <= linesToClear; cLine++ {
		designer.eraseLineWithXCoordinates(cLine)
	}

	startOfCliChatText := int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.moveCursor(coordinates{x: startOfError, y: startOfCliChatText + 15})
	designer.setDrawing(message).setColor(escapecode.Red).printAndResetColors()

	startOfLoginBox := startOfCliChatText + 13
	designer.
		setCursorColor(escapecode.White).
		moveCursor(coordinates{
			x: cui.xCoordinateToTypeLogin, y: startOfLoginBox + 3,
		})
}

// TODO: exibir os comandos disponiveis
func (cui *CUI) drawChatInterface() {
	designer := newConsoleDesigner()
	designer.clearTerminal()
	defer cui.moveCursorToTypeInChat(designer)

	designer.setColor(escapecode.White)
	chatBox := cui.chatBoxToSlice()
	for currentLine, lineText := range chatBox {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine + 1), y: 1})
		designer.setDrawing(lineText).print()
	}
	designer.resetColors()

	designer.setColor(escapecode.Yellow)
	typingBox := cui.typingBoxToSlice()
	startOfDrawing := cui.consoleHeight - cui.typingBoxHeight
	for currentLine, lineText := range typingBox {
		currentLine += int(startOfDrawing)
		designer.setCursorCoordinates(coordinates{x: int16(currentLine), y: 1})
		designer.setDrawing(lineText).print()
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
		designer.setDrawing(line).print()
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

func (cui *CUI) DrawNewLineInChat(line *ChatLine) {
	cui.addChatLine(line)
	cui.drawChatLines()
}

func (cui *CUI) addChatLine(line *ChatLine) {
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

// TODO: tentar resolver problema quando a mensagem passa do limite da tela
func (cui *CUI) drawChatLines() {
	designer := newConsoleDesigner()
	defer cui.moveCursorToTypeInChat(designer)

	designer.setColor(escapecode.White)
	for key, chatLine := range cui.chatLines {
		info := newConsoleDesigner().
			setColor(chatLine.InfoColor).
			setDrawing(chatLine.Info).
			toStringWithResetColors()
		lineText := info + " " + chatLine.Text
		amountOfSpaces := (int(cui.consoleWidth) - 4) - len(chatLine.Info+" "+chatLine.Text)
		if amountOfSpaces <= 0 {
			amountOfSpaces = 0
		}

		lineStr := fmt.Sprintf(` │%s%s│ `, lineText, strings.Repeat(" ", amountOfSpaces))
		designer.setCursorCoordinates(coordinates{x: int16(key + 2), y: 1})
		designer.eraseLine()

		designer.setDrawing(lineStr).printAndResetColors()
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
	cui.drawChatLines()
}

func (cui *CUI) DrawLineAndExit(code uint8, chatLine ChatLine) {
	designer := newConsoleDesigner()
	designer.clearTerminal()

	info := newConsoleDesigner().
		setColor(chatLine.InfoColor).
		setDrawing(chatLine.Info).
		toStringWithResetColors()
	lineStr := info + " " + chatLine.Text
	designer.
		setCursorCoordinates(coordinates{
			x: 1, y: 1,
		}).
		setDrawing(lineStr + "\n").
		printAndResetColors()
	os.Exit(int(code))
}

func (cui *CUI) drawLoading(length uint, color escapecode.ColorCode) error {
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

		err = designer.setDrawing(moveLeftCode + loadingString).print()
		if err != nil {
			return err
		}

		counter++
	}

	designer.resetColors()

	return nil
}
