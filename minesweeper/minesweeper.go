package minesweeper

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type MineField struct {
	width      int
	height     int
	numOfBombs int
	tiles      []Tile
	numOfFlags int
}

func InitGame() {
	mf := createMineField(10, 10, 10)
	mf.generateTiles()
	mf.calibrate()

	for !mf.isGameOver() {
		fmt.Print(mf)
		command, err := readCommand()
		if err != nil {
			fmt.Printf("Invalid input: %v\n", err)
		}
		err = mf.processCommand(command)

		if err != nil {
			fmt.Println(err)
		}
	}

	if mf.isVictory() {
		fmt.Println("Congratulations! You have cleared the minefield!")
	} else {
		fmt.Println("Game over! You lost!")
	}
}

func (mf *MineField) isVictory() bool {
	if !mf.revealedBomb() {
		return true
	}

	return false
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

func (mf *MineField) processCommand(command string) error {
	err := mf.validateCommand(command)

	if err != nil {
		return err
	}

	fields := strings.Fields(command)

	verb, x, y := fields[0], atoi(fields[1]), atoi(fields[2])

	switch verb {
	case "reveal":
		mf.revealTile(x, y)
	case "flag":
		mf.flagTile(x, y)
	case "unflag":
		mf.unflagTile(x, y)
	default:
		fmt.Println("Invalid command.")
	}
	return nil
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

			if tile.val == 0 && !tile.revealed {
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

func (mf *MineField) validateCommand(command string) error {
	fields := strings.Fields(command)

	verb, x, y := fields[0], atoi(fields[1]), atoi(fields[2])

	if verb != "reveal" && verb != "flag" && verb != "unflag" {
		return fmt.Errorf("This is not a valid command")
	}

	if len(fields) < 3 {
		return fmt.Errorf("Not enough arguments")
	}

	if x < 0 || x > mf.width || y < 0 || y > mf.height {
		return fmt.Errorf("Invalid coords")
	}

	return nil
}

func readCommand() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	return command, err
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
			tiles[i] = NewTile(false, true)
		} else {
			tiles[i] = NewTile(false, false)
		}
	}

	rand.Shuffle(len(tiles), func(i, j int) { tiles[i], tiles[j] = tiles[j], tiles[i] })

	mf.tiles = tiles
}

func (mf MineField) String() string {
	res := "   "

	for i := 0; i < mf.width; i++ {
		res += fmt.Sprintf("|%d|", i)
	}

	res += "\n"

	for i := 0; i < mf.height; i++ {
		res += fmt.Sprintf("|%d|", i)
		for j := 0; j < mf.width; j++ {
			res += fmt.Sprintf("%v", mf.tiles[i*mf.width+j])
		}
		res += "\n"
	}

	return res
}

func (mf *MineField) calibrate() {
	for i := 0; i < len(mf.tiles); i++ {
		x := i % mf.width
		y := i / mf.width

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
	surroundingPoints := []Point{}

	surroundingPoints = append(surroundingPoints, Point{x - 1, y - 1})
	surroundingPoints = append(surroundingPoints, Point{x - 1, y})
	surroundingPoints = append(surroundingPoints, Point{x - 1, y + 1})
	surroundingPoints = append(surroundingPoints, Point{x, y + 1})
	surroundingPoints = append(surroundingPoints, Point{x + 1, y + 1})
	surroundingPoints = append(surroundingPoints, Point{x + 1, y})
	surroundingPoints = append(surroundingPoints, Point{x + 1, y - 1})
	surroundingPoints = append(surroundingPoints, Point{x, y - 1})

	filtered := []Point{}

	for _, point := range surroundingPoints {
		if point.x >= 0 && point.x < mf.width && point.y >= 0 && point.y < mf.height {
			filtered = append(filtered, point)
		}
	}

	return filtered
}
