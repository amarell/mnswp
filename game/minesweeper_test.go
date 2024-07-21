package minesweeper

import "testing"

const testMfHeight = 20
const testMfWidth = 20
const testNumOfBombs = 10

func createTestMinefield() MineField {
	return createMineField(testMfHeight, testMfWidth, testNumOfBombs)
}

func TestCreateMinefield(t *testing.T) {
	mf := createTestMinefield()

	if mf.height != testMfHeight {
		t.Errorf("Wrong height when creating minefield: Expected %d, got %d", testMfHeight, mf.height)
	}

	if mf.width != testMfWidth {
		t.Errorf("Wrong width when creating minefield: Expected %d, got %d", testMfWidth, mf.width)
	}
}

func TestGenerationOfTiles(t *testing.T) {
	mf := createTestMinefield()
	mf.generateTiles()

	if len(mf.tiles) != testMfHeight*testMfWidth {
		t.Errorf("Wrong number of tiles created: Expected %d, got %d", testMfHeight*testMfWidth, len(mf.tiles))
	}
}

func TestGenerationOfBombTiles(t *testing.T) {
	mf := createTestMinefield()
	mf.generateTiles()

	var bombCount = 0
	for i := 0; i < len(mf.tiles); i++ {
		if mf.tiles[i].isBomb {
			bombCount++
		}
	}

	if bombCount != testNumOfBombs {
		t.Errorf("Unexpected number of bombs generated: Expected %d, got %d", testNumOfBombs, bombCount)
	}
}

func TestRevealAllTiles(t *testing.T) {
	mf := createTestMinefield()
	mf.generateTiles()

	var countRevealed = 0

	for i := 0; i < len(mf.tiles); i++ {
		if mf.tiles[i].revealed {
			countRevealed++
		}
	}

	if countRevealed != 0 {
		t.Errorf("After generating the mine field, no tiles should be revealed. Expected %d, got %d", 0, countRevealed)
	}

	mf.revealAllFields()

	for i := 0; i < len(mf.tiles); i++ {
		if mf.tiles[i].revealed {
			countRevealed++
		}
	}

	if countRevealed != len(mf.tiles) {
		t.Errorf("After calling reveal all fields, all tiles should be revealed. Expected %d, got %d", len(mf.tiles), countRevealed)
	}
}

func TestGetSurroundingTiles(t *testing.T) {
	mf := createTestMinefield()
	mf.generateTiles()

	getSurroundingTilesTests := []struct {
		input          Point
		expectedOutput []Point
		name           string
	}{
		{
			input: Point{0, 0},
			expectedOutput: []Point{
				{1, 0},
				{0, 1},
				{1, 1},
			},
			name: "Test origin neighbors",
		},
		{
			input: Point{1, 1},
			expectedOutput: []Point{
				{0, 0},
				{0, 1},
				{0, 2},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
				{1, 2},
			},
			name: "Test a central tile's neighbors",
		},
		{
			input: Point{0, 3},
			expectedOutput: []Point{
				{0, 2},
				{1, 2},
				{1, 3},
				{1, 4},
				{0, 4},
			},
			name: "Test a side tile's neighbors",
		},
	}

	for _, test := range getSurroundingTilesTests {
		output := mf.getSurroundingTiles(test.input.x, test.input.y)

		tileSet := make(map[Point]bool)

		for _, tile := range test.expectedOutput {
			tileSet[tile] = true
		}

		if len(output) != len(test.expectedOutput) {
			t.Errorf("Expected output to have length of %d, got %d", len(test.expectedOutput), len(output))
		}

		for _, tile := range output {
			if tileSet[tile] == false {
				t.Errorf("Unexpected surrounding tiles received: Expected %v, got %v. Test name: %v", test.expectedOutput, output, test.name)
			}
		}
	}
}
