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
	consoleHeight    uint16
	consoleWidth     uint16
	chatBoxHeight    uint16
	typingBoxHeight  uint16
	xCoodinateToType int16
	sizeListener     *tsize.SizeListener

	currentInterface string
	logged           chan bool
	visibleLines     []*ChatLine
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
		consoleHeight:    uint16(size.Height),
		consoleWidth:     uint16(size.Width),
		chatBoxHeight:    uint16(size.Height) - typingBoxHeight,
		xCoodinateToType: int16(uint16(size.Height) - 8),
		typingBoxHeight:  typingBoxHeight,
		sizeListener:     listener,
		logged:           make(chan bool),
		visibleLines:     make([]*ChatLine, 0),
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
		cui.updateSizes(newSize)
		cui.adaptNewConsoleSize()
	}
}

func (cui *CUI) updateSizes(newSize tsize.Size) {
	cui.consoleWidth = uint16(newSize.Width)
	cui.consoleHeight = uint16(newSize.Height)
	cui.chatBoxHeight = uint16(newSize.Height) - cui.typingBoxHeight
}

func (cui *CUI) adaptNewConsoleSize() {
	switch cui.currentInterface {
	case "login":
		cui.drawLoginInterface()
	case "chat":
		cui.drawChatInterface()
	}
}

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
	var startOfCliChatText int16 = int16(cui.consoleWidth/2) - int16(cliChatTextWidth/2)
	designer.moveCursor(coordinates{x: 11, y: startOfCliChatText + 15})
	designer.setDrawing(message).setColor(Red).print()
	designer.resetColors()

	// TODO: limpar login e erro anterior
	startOfLoginBox := startOfCliChatText + 13
	designer.resetColors()
	designer.
		setCursorColor(White).
		moveCursor(coordinates{
			x: 9, y: startOfLoginBox + 3,
		})
}

// TODO: exibir os comandos disponiveis
func (cui *CUI) drawChatInterface() {
	typingBox := make([]string, cui.typingBoxHeight, cui.typingBoxHeight)
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

		typingBox = append(typingBox, lineStr)
	}

	chatBox := make([]string, cui.chatBoxHeight, cui.chatBoxHeight)
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

	designer := newConsoleDesigner()
	designer.clearTerminal()
	defer func() {
		designer.
			setCursorColor(White).
			moveCursor(coordinates{
				x: cui.xCoodinateToType,
				y: 5,
			})
	}()

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

	cui.adaptVisibleLinesInChatBox()
}

func (cui *CUI) DrawNewLineInChat(line *ChatLine) {
	cui.addChatLine(line)
	cui.drawVisibleLinesInChat()
}

func (cui *CUI) addChatLine(line *ChatLine) {
	cui.visibleLines = append(cui.visibleLines, line)
	cui.updateVisibleLines()
}

func (cui *CUI) adaptVisibleLinesInChatBox() {
	cui.updateVisibleLines()
	cui.drawVisibleLinesInChat()
}

func (cui *CUI) updateVisibleLines() {
	numberOfVisibleLines := uint16(len(cui.visibleLines))
	if numberOfVisibleLines > cui.chatBoxHeight {
		diff := numberOfVisibleLines - cui.chatBoxHeight
		cui.visibleLines = cui.visibleLines[diff:]
	}
}

func (cui *CUI) drawVisibleLinesInChat() {
	designer := newConsoleDesigner()
	defer func() {
		designer.resetColors()
		designer.
			setCursorColor(White).
			moveCursor(coordinates{
				x: cui.xCoodinateToType,
				y: 5,
			})
	}()

	for key, chatLine := range cui.visibleLines {
		designer.setCursorCoordinates(coordinates{x: int16(key + 2), y: 3})

		info := newConsoleDesigner().
			setColor(chatLine.InfoColor).
			setDrawing(chatLine.Info).
			toStringWithResetColors()
		designer.setColor(White).setDrawing(info + " " + chatLine.Text).print()
		designer.resetColors()
	}
}

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
