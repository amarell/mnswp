package main

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

func initGame() Dungeon {
	var username string
	fmt.Println("Enter your username: ")
	fmt.Scanln(&username)
	player := createPlayer(username)

	d := createDungeon(10, 10, 3, player, 10)
	d.generateVampires()

	return d
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
	var command string
	for d.numOfMoves > 0 {
		if len(d.vampires) == 0 {
			fmt.Println("Congratulations, you have won!")
			return
		}
		d.printSituation()
		fmt.Scanln(&command)
		d.processCommand(command)
		d.numOfMoves = d.numOfMoves - 1
		d.moveVamps()
	}

	fmt.Println("Game over!")
}

func (d *Dungeon) printSituation() {
	fmt.Printf("There are %d vampires left and you have %d moves\n", len(d.vampires), d.numOfMoves)
	fmt.Printf("%s (%d, %d) \n", d.player.name, d.player.x, d.player.y)

	for _, vamp := range d.vampires {
		fmt.Printf("v: (%d, %d) \n", vamp.x, vamp.y)
	}

	fmt.Println(strings.Repeat("=", d.height*3))
	for i := 0; i < d.height; i++ {
		for j := 0; j < d.width; j++ {
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
	switch direction {
	case "w":
		if d.player.x == 0 {
			return false
		}
	case "s":
		if d.player.x == d.width-1 {
			return false
		}
	case "a":
		if d.player.y == 0 {
			return false
		}
	case "d":
		if d.player.y == d.height-1 {
			return false
		}
	}

	return true

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
