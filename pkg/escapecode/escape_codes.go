// TODO: pensar em colocar o escapecode dentro da pasta pkg no client
package escapecode

type ColorCode string

const (
	DefaultColor  ColorCode = "\x1b[39m"
	Black         ColorCode = "\x1b[30m"
	Red           ColorCode = "\x1b[31m"
	Green         ColorCode = "\x1b[32m"
	Yellow        ColorCode = "\x1b[33m"
	Blue          ColorCode = "\x1b[34m"
	Magenta       ColorCode = "\x1b[35m"
	Cyan          ColorCode = "\x1b[36m"
	LightGray     ColorCode = "\x1b[37m"
	DarkGray      ColorCode = "\x1b[90m"
	BrightRed     ColorCode = "\x1b[91m"
	BrightGreen   ColorCode = "\x1b[92m"
	BrightYellow  ColorCode = "\x1b[93m"
	BrightBlue    ColorCode = "\x1b[94m"
	BrightMagenta ColorCode = "\x1b[95m"
	BrightCyan    ColorCode = "\x1b[96m"
	White         ColorCode = "\x1b[97m"

	DefaultBackground ColorCode = "\x1b[49m"
	BgBlack           ColorCode = "\x1b[40m"
	BgRed             ColorCode = "\x1b[41m"
	BgGreen           ColorCode = "\x1b[42m"
	BgYellow          ColorCode = "\x1b[43m"
	BgBlue            ColorCode = "\x1b[44m"
	BgMagenta         ColorCode = "\x1b[45m"
	BgCyan            ColorCode = "\x1b[46m"
	BgLightGray       ColorCode = "\x1b[47m"
	BgDarkGray        ColorCode = "\x1b[100m"
	BgBrightRed       ColorCode = "\x1b[101m"
	BgBrightGreen     ColorCode = "\x1b[102m"
	BgBrightYellow    ColorCode = "\x1b[103m"
	BgBrightBlue      ColorCode = "\x1b[104m"
	BgBrightMagenta   ColorCode = "\x1b[105m"
	BgBrightCyan      ColorCode = "\x1b[106m"
	BgWhite           ColorCode = "\x1b[107m"
)

type escapeCode string

const (
	Reset escapeCode = "\x1b[0m"

	ClearScreen   escapeCode = "\x1b[2J"
	MoveHome      escapeCode = "\x1b[H"
	MoveCursor    escapeCode = "\x1b[{lines};{columns}H"
	MoveUp        escapeCode = "\x1b[{n}A"
	MoveDown      escapeCode = "\x1b[{n}B"
	MoveRight     escapeCode = "\x1b[{n}C"
	MoveLeft      escapeCode = "\x1b[{n}D"
	MoveNextLine  escapeCode = "\x1b[E"
	MovePrevLine  escapeCode = "\x1b[F"
	SaveCursor    escapeCode = "\x1b[s"
	RestoreCursor escapeCode = "\x1b[u"
	HideCursor    escapeCode = "\x1b[?25l"
	ShowCursor    escapeCode = "\x1b[?25h"
	EraseLine     escapeCode = "\x1b[2K"

	Bold       escapeCode = "\x1b[1m"
	Dim        escapeCode = "\x1b[2m"
	Italic     escapeCode = "\x1b[3m"
	Underline  escapeCode = "\x1b[4m"
	Blink      escapeCode = "\x1b[5m"
	Reverse    escapeCode = "\x1b[7m"
	Hidden     escapeCode = "\x1b[8m"
	CrossedOut escapeCode = "\x1b[9m"
)
