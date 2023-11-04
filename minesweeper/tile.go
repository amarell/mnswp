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
	4: ColorCyan,
	5: ColorRed,
	6: ColorRed,
	7: ColorRed,
	8: ColorRed,
}

type Tile struct {
	// x        int
	// y        int
	flagged  bool
	revealed bool
	isBomb   bool
	val      int
}

func NewTile(isBomb bool) Tile {
	return Tile{
		isBomb: isBomb,
	}
}

func (t Tile) String() string {
	if t.revealed && t.isBomb {
		return coloredString("[*]", ColorRed)
	} else if t.flagged {
		return coloredString("[!]", ColorPurple)
	} else if t.revealed {
		return coloredString(fmt.Sprintf("[%d]", t.val), colorMap[t.val])
	} else {
		return fmt.Sprint("[ ]")
	}
}

func coloredString(str, color string) string {
	return fmt.Sprintf("%s%s%s", color, str, ColorReset)
}
