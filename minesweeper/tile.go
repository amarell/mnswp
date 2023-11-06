package minesweeper

import "fmt"

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

var colorMap = map[int]string{
	0: ColorBlue,
	1: ColorYellow,
	2: ColorGreen,
	3: ColorRed,
	4: ColorRed,
	5: ColorRed,
	6: ColorRed,
	7: ColorRed,
	8: ColorRed,
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
		return coloredString("[*]", ternary(t.selected, ColorCyan, ColorRed))
	} else if t.flagged {
		return coloredString("[!]", ternary(t.selected, ColorCyan, ColorPurple))
	} else if t.revealed {
		return coloredString(fmt.Sprintf("[%d]", t.val), ternary(t.selected, ColorCyan, colorMap[t.val]))
	} else {
		return coloredString(fmt.Sprint("[ ]"), ternary(t.selected, ColorCyan, ColorReset))
	}
}

func ternary(cond bool, val1, val2 string) string {
	if cond {
		return val1
	}
	return val2
}

func coloredString(str, color string) string {
	return fmt.Sprintf("%s%s%s", color, str, ColorReset)
}
