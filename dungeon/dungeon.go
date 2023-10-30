package dungeon

import (
	"fmt"
	"math/rand"
	"strings"
)

/**
 *
 * @author amarell
 */

type Dungeon struct {
	width         int
	height        int
	numOfVampires int
	vampires      []Vampire
	player        Player
	numOfMoves    int
}

func InitGame() {
	var username string
	fmt.Println("Enter your username: ")
	fmt.Scanln(&username)
	player := createPlayer(username)

	d := createDungeon(10, 10, 3, player, 10)
	d.generateVampires()
	d.run()
}

func createDungeon(width int, height int, numOfVampires int, player Player, numOfMoves int) Dungeon {
	return Dungeon{
		width:         width,
		height:        height,
		numOfVampires: numOfVampires,
		vampires:      []Vampire{},
		player:        player,
		numOfMoves:    numOfMoves,
	}
}

func (d *Dungeon) generateVampires() {
	for i := 0; i < d.numOfVampires; i++ {
		randX := rand.Intn(d.width-1) + 1
		randY := rand.Intn(d.height-1) + 1
		d.vampires = append(d.vampires, Vampire{randX, randY})
	}
}

func (d *Dungeon) run() {
	for !d.isGameOver() {
		d.printSituation()
		command, err := d.readCommand()
		if err != nil {
			fmt.Printf("Invalid input: %v\n", err)
			continue
		}
		d.processCommand(command)
		d.numOfMoves--
		d.moveVamps()
	}

	if d.isVictory() {
		fmt.Println("Congratulations, you have won!")
	} else {
		fmt.Println("Game over!")
	}
}

func (d *Dungeon) isGameOver() bool {
	return len(d.vampires) == 0 || d.numOfMoves == 0
}

func (d *Dungeon) isVictory() bool {
	return len(d.vampires) == 0
}

func (d *Dungeon) readCommand() (string, error) {
	var command string
	_, err := fmt.Scanln(&command)
	return command, err
}

func (d *Dungeon) printSituation() {
	fmt.Printf("There are %d vampires left and you have %d moves\n", len(d.vampires), d.numOfMoves)
	d.printEntities()
	d.printMap()
}

func (d *Dungeon) printEntities() {
	fmt.Print(d.player)
	for _, vamp := range d.vampires {
		fmt.Print(vamp)
	}
}

func (d *Dungeon) printMap() {
	fmt.Println(strings.Repeat("=", d.height*3))
	for i := 0; i < d.width; i++ {
		for j := 0; j < d.height; j++ {
			if d.player.x == i && d.player.y == j {
				fmt.Printf("[@]")
			} else if d.isVampireLocatedAtCoords(i, j) {
				fmt.Printf("[v]")
			} else {
				fmt.Printf("[ ]")
			}
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("=", d.height*3))
}

func (d *Dungeon) isVampireLocatedAtCoords(x, y int) bool {
	for _, vamp := range d.vampires {
		if vamp.x == x && vamp.y == y {
			return true
		}
	}

	return false
}

func (d *Dungeon) processCommand(command string) {
	for i := 0; i < len(command); i++ {
		direction := string(command[i])
		if d.isValidMove(direction) {
			d.player.move(direction)
		}
		overlapExists, index := d.overlapExists()
		if overlapExists {
			d.killVamp(index)
		}
	}
}

func (d *Dungeon) isValidMove(direction string) bool {
	newX, newY := d.player.x, d.player.y

	switch direction {
	case "w":
		newX--
	case "s":
		newX++
	case "a":
		newY--
	case "d":
		newY++
	}

	return newX >= 0 && newX < d.width && newY >= 0 && newY < d.height
}

func (d *Dungeon) killVamp(index int) {
	d.vampires = append(d.vampires[:index], d.vampires[index+1:]...)
}

func (d *Dungeon) overlapExists() (bool, int) {
	for i, vamp := range d.vampires {
		if vamp.x == d.player.x && vamp.y == d.player.y {
			return true, i
		}
	}

	return false, -1
}

func (d *Dungeon) moveVamps() {
	for i, vamp := range d.vampires {
		d.vampires[i] = *vamp.move(d.height)
	}
}
