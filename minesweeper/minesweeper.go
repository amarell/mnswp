package minesweeper

import (
	"fmt"
	"l/util"
	"math/rand"
	"os"
	"strconv"
)

type Input int

const (
	EXIT    = 0
	UP      = 1
	DOWN    = 2
	LEFT    = 3
	RIGHT   = 4
	FLAG    = 5
	UNFLAG  = 6
	REVEAL  = 7
	UNKNOWN = 8
)

type MineField struct {
	width      int
	height     int
	numOfBombs int
	tiles      []Tile
	numOfFlags int
}

func InitGame() {
	util.CleanTerminal()
	mf := createMineField(10, 10, 10)
	mf.generateTiles()
	mf.calibrate()

	for !mf.isGameOver() {
		fmt.Print(mf)
		input := input()
		exit := mf.processInput(input)

		if exit {
			break
		}

		util.CleanTerminal()
	}

	if mf.isVictory() {
		fmt.Println("Congratulations! You have cleared the minefield!\n\r")
	} else {
		fmt.Println("Game over! You lost!\n\r")
	}

	mf.revealAllFields()
	fmt.Println(mf)
}

func (mf *MineField) processInput(input Input) bool {

	selectedTileIndex := mf.getSelectedTile()

	if selectedTileIndex == -1 {
		fmt.Println("Something went wrong!")
	}

	x, y := selectedTileIndex%mf.width, selectedTileIndex/mf.width

	switch input {
	case EXIT:
		return true
	case UNKNOWN:
		// nothing
	case DOWN, UP, RIGHT, LEFT:
		mf.move(input)
	case REVEAL:
		mf.revealTile(x, y)
	case FLAG:
		mf.flagTile(x, y)
	case UNFLAG:
		mf.unflagTile(x, y)
	default:
		// nothing
	}
	return false
}

func (mf *MineField) move(input Input) {
	selectedTileIndex := mf.getSelectedTile()

	if selectedTileIndex == -1 {
		fmt.Println("Something went wrong\n\r")
		return
	}

	mf.tiles[selectedTileIndex].selected = false

	switch input {
	case DOWN:
		mf.tiles[(selectedTileIndex+mf.width)%len(mf.tiles)].selected = true
	case RIGHT:
		mf.tiles[(selectedTileIndex+1)%len(mf.tiles)].selected = true
	case LEFT:
		mf.tiles[(len(mf.tiles)+(selectedTileIndex-1))%len(mf.tiles)].selected = true
	case UP:
		mf.tiles[(len(mf.tiles)+(selectedTileIndex-mf.width))%len(mf.tiles)].selected = true
	}
}

func (mf *MineField) getSelectedTile() int {
	for index, tile := range mf.tiles {
		if tile.selected {
			return index
		}
	}
	// should never reach this
	return -1
}

func input() Input {
	for {
		b := make([]byte, 3)
		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			return EXIT
		}

		switch b[0] {
		// 0x03 - code for control + C
		case 0x03, 'q':
			return EXIT
		case 'd':
			return RIGHT
		case 'w':
			return UP
		case 's':
			return DOWN
		case 'a':
			return LEFT
		case 'f':
			return FLAG
		case 'u':
			return UNFLAG
		case 'r':
			return REVEAL
		case 0x1b:
			if b[1] == '[' {
				switch b[2] {
				case 'A':
					return UP
				case 'B':
					return DOWN
				case 'C':
					return RIGHT
				case 'D':
					return LEFT
				default:
					return UNKNOWN
				}
			}
			return UNKNOWN
		default:
			return UNKNOWN
		}

	}
}

func (mf *MineField) isVictory() bool {
	return !mf.revealedBomb()
}

func (mf *MineField) isGameOver() bool {
	if mf.numOfFlags == 0 && mf.isValidSolution() {
		return true
	}

	if mf.numOfBombs == 0 || mf.revealedBomb() {
		return true
	}

	return false
}

func (mf *MineField) isValidSolution() bool {
	for _, tile := range mf.tiles {
		if tile.flagged && !tile.isBomb {
			return false
		}
	}

	return true
}

func (mf *MineField) revealedBomb() bool {
	for i := 0; i < len(mf.tiles); i++ {
		if mf.tiles[i].isBomb && mf.tiles[i].revealed {
			return true
		}
	}

	return false
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return num
}

func (mf *MineField) revealTile(x, y int) {
	tile := &mf.tiles[y*mf.width+x]
	tile.revealed = true

	if tile.val == 0 {
		surroundingPoints := mf.getSurroundingTiles(x, y)

		for _, point := range surroundingPoints {
			tile := mf.tiles[point.y*mf.width+point.x]

			if !tile.revealed {
				mf.revealTile(point.x, point.y)
			}
		}
	}

}

func (mf *MineField) flagTile(x, y int) {
	mf.tiles[y*mf.width+x].flagged = true
	mf.numOfFlags--
}

func (mf *MineField) unflagTile(x, y int) {
	mf.tiles[y*mf.width+x].flagged = false
	mf.numOfFlags++
}

func createMineField(width, height, numOfBombs int) MineField {
	return MineField{
		width: width, height: height, numOfBombs: numOfBombs, numOfFlags: numOfBombs,
	}
}

func (mf *MineField) generateTiles() {
	tiles := make([]Tile, mf.height*mf.width)

	for i := 0; i < len(tiles); i++ {
		if i < mf.numOfBombs {
			tiles[i] = NewTile(true) // bombs
		} else {
			tiles[i] = NewTile(false)
		}
	}

	rand.Shuffle(len(tiles), func(i, j int) { tiles[i], tiles[j] = tiles[j], tiles[i] })

	tiles[0].selected = true

	mf.tiles = tiles
}

func (mf MineField) String() string {
	res := "   "

	for i := 0; i < mf.width; i++ {
		res += fmt.Sprintf("|%d|", i)
	}

	res += "\n\r"

	for i := 0; i < mf.height; i++ {
		res += fmt.Sprintf("|%d|", i)
		for j := 0; j < mf.width; j++ {
			res += fmt.Sprintf("%v", mf.tiles[i*mf.width+j])
		}
		res += "\n\r"
	}

	return res
}

func (mf *MineField) calibrate() {
	for i := 0; i < len(mf.tiles); i++ {
		x, y := i%mf.width, i/mf.width

		mf.tiles[y*mf.width+x].val = mf.getNumberOfSurroundingBombs(x, y)
	}
}

func (mf *MineField) getNumberOfSurroundingBombs(x, y int) int {
	surroundingTiles := mf.getSurroundingTiles(x, y)
	surroundingBombs := 0

	for _, tile := range surroundingTiles {
		if mf.tiles[tile.y*mf.width+tile.x].isBomb {
			surroundingBombs++
		}
	}

	return surroundingBombs
}

type Point struct {
	x, y int
}

func (mf *MineField) getSurroundingTiles(x, y int) []Point {
	surroundingPoints := [8]Point{
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
		{x + 1, y},
		{x + 1, y - 1},
		{x, y - 1},
	}
	filtered := []Point{}

	for _, point := range surroundingPoints {
		if point.x >= 0 && point.x < mf.width && point.y >= 0 && point.y < mf.height {
			filtered = append(filtered, point)
		}
	}

	return filtered
}

func (mf *MineField) revealAllFields() {
	for i := 0; i < len(mf.tiles); i++ {
		mf.tiles[i].revealed = true
	}
}
