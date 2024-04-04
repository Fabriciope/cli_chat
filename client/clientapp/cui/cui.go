package cui

import (
	"fmt"
	"strings"
	"time"

	"github.com/Fabriciope/cli_chat/shared"
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
	consoleHeight     uint16
	consoleWidth      uint16
	chatBoxHeight     uint16
	typingBoxHeight   uint16
	xCoordinateToType int16
	sizeListener      *tsize.SizeListener

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
		consoleHeight:     uint16(size.Height),
		consoleWidth:      uint16(size.Width),
		chatBoxHeight:     uint16(size.Height) - typingBoxHeight,
		xCoordinateToType: int16(uint16(size.Height) - 9),
		typingBoxHeight:   typingBoxHeight,
		sizeListener:      listener,
		logged:            make(chan bool),
		chatLines:         make([]*ChatLine, 0),
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
	designer.setColor(shared.Blue).print()
	for currentLine, text := range cliChatText {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine + 2), y: startOfCliChatText})
		designer.setDrawing(text).print()
	}
	designer.resetColors()

	startOfLoginBox := startOfCliChatText + 13
	designer.
		setCursorColor(shared.White).
		moveCursor(coordinates{
			x: 10, y: startOfLoginBox + 3,
		})

	cui.currentInterface = interfaces[Login]
}

func (cui *CUI) DrawLoginError(message string) {
	designer := newConsoleDesigner()
	var startOfCliChatText int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.moveCursor(coordinates{x: 12, y: startOfCliChatText + 15})
	designer.setDrawing(message).setColor(shared.Red).print()
	designer.resetColors()

	// TODO: limpar login e erro anterior
	startOfLoginBox := startOfCliChatText + 13
	designer.
		setCursorColor(shared.White).
		moveCursor(coordinates{
			x: 10, y: startOfLoginBox + 3,
		})
}

// TODO: exibir os comandos disponiveis
func (cui *CUI) drawChatInterface() {
	designer := newConsoleDesigner()
	designer.clearTerminal()
	defer cui.moveCursorToTypeInChat(designer)

	designer.setColor(shared.White)
	chatBox := cui.chatBoxToSlice()
	for currentLine, lineText := range chatBox {
		designer.setCursorCoordinates(coordinates{x: int16(currentLine + 1), y: 1})
		designer.setDrawing(lineText).print()
	}
	designer.resetColors()

	designer.setColor(shared.Yellow)
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

	designer.setColor(shared.Yellow)
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

	for key, chatLine := range cui.chatLines {
		info := newConsoleDesigner().
			setColor(chatLine.InfoColor).
			setDrawing(chatLine.Info).
			toStringWithResetColors()
		lineText := info + " " + chatLine.Text
		amountOfSpaces := (int(cui.consoleWidth) - 4) - len(chatLine.Info+" "+chatLine.Text)
		lineStr := fmt.Sprintf(` │%s%s│ `, lineText, strings.Repeat(" ", amountOfSpaces))

		designer.setCursorCoordinates(coordinates{x: int16(key + 2), y: 1})
		designer.eraseLine()

		designer.setColor(shared.White).setDrawing(lineStr).print()
		designer.resetColors()
	}
}

func (cui *CUI) moveCursorToTypeInChat(designer *consoleDesigner) {
	designer.resetColors()
	designer.
		setCursorColor(shared.White).
		moveCursor(coordinates{
			x: cui.xCoordinateToType,
			y: 5,
		})
}

func (cui *CUI) adaptChatLinesOnTerminal() {
	cui.checkIfChatLinesExceededTheLimit()
	cui.drawChatLines()
}

func (cui *CUI) drawLoading(length uint, color shared.ColorCode) error {
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
