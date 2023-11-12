package minesweeper

import "fmt"

var (
	// Normal colors
	nBlack   = []byte{'\033', '[', '3', '0', 'm'}
	nRed     = []byte{'\033', '[', '3', '1', 'm'}
	nGreen   = []byte{'\033', '[', '3', '2', 'm'}
	nYellow  = []byte{'\033', '[', '3', '3', 'm'}
	nBlue    = []byte{'\033', '[', '3', '4', 'm'}
	nMagenta = []byte{'\033', '[', '3', '5', 'm'}
	nCyan    = []byte{'\033', '[', '3', '6', 'm'}
	nWhite   = []byte{'\033', '[', '3', '7', 'm'}
	// Bright colors
	bBlack   = []byte{'\033', '[', '3', '0', ';', '1', 'm'}
	bRed     = []byte{'\033', '[', '3', '1', ';', '1', 'm'}
	bGreen   = []byte{'\033', '[', '3', '2', ';', '1', 'm'}
	bYellow  = []byte{'\033', '[', '3', '3', ';', '1', 'm'}
	bBlue    = []byte{'\033', '[', '3', '4', ';', '1', 'm'}
	bMagenta = []byte{'\033', '[', '3', '5', ';', '1', 'm'}
	bCyan    = []byte{'\033', '[', '3', '6', ';', '1', 'm'}
	bWhite   = []byte{'\033', '[', '3', '7', ';', '1', 'm'}

	reset = []byte{'\033', '[', '0', 'm'}
)

var colorMap = map[int][]byte{
	0: nBlue,
	1: bGreen,
	2: nGreen,
	3: nYellow,
	4: bRed,
	5: nRed,
	6: nRed,
	7: nRed,
	8: nRed,
}

type Tile struct {
	flagged  bool
	revealed bool
	isBomb   bool
	selected bool
	val      int
}

func NewTile(isBomb bool) Tile {
	return Tile{
		isBomb: isBomb,
	}
}

func (t Tile) String() string {
	if t.revealed && t.isBomb {
		return coloredString("[*]", ternary(t.selected, string(nCyan), string(bRed)))
	} else if t.flagged {
		return coloredString("[!]", ternary(t.selected, string(nCyan), string(nMagenta)))
	} else if t.revealed {
		return coloredString(fmt.Sprintf("[%d]", t.val), ternary(t.selected, string(nCyan), string(colorMap[t.val])))
	} else {
		return coloredString(fmt.Sprint("[ ]"), ternary(t.selected, string(nCyan), string(reset)))
	}
}

func ternary(cond bool, val1, val2 string) string {
	if cond {
		return val1
	}
	return val2
}

func coloredString(str, color string) string {
	return fmt.Sprintf("%s%s%s", color, str, reset)
}
