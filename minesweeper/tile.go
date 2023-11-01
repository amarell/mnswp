package minesweeper

import "fmt"

type Tile struct {
	// x        int
	// y        int
	flagged  bool
	revealed bool
	isBomb   bool
	val      int
}

func NewTile(revealed bool, isBomb bool) Tile {
	return Tile{
		revealed: revealed, isBomb: isBomb,
	}
}

func (t Tile) String() string {
	if t.revealed && t.isBomb {
		return fmt.Sprint("[*]")
	} else if t.flagged {
		return "[!]"
	} else if t.revealed {
		return fmt.Sprintf("[%d]", t.val)
	} else {
		return fmt.Sprint("[ ]")
	}
}
