package dungeon

import (
	"fmt"
	"math/rand"
)

type Vampire struct {
	x int
	y int
}

var moves = []string{"w", "a", "s", "d"}

func (v *Vampire) move(upperBound int) *Vampire {
	direction := moves[rand.Intn(len(moves))]

	for !v.isValidVampMove(direction, upperBound) {
		direction = moves[rand.Intn(len(moves))]
	}

	switch direction {
	case "w":
		v.x--
	case "a":
		v.y--
	case "s":
		v.x++
	case "d":
		v.y++
	default:
		fmt.Println("Something went seriously wrong")
	}

	return v
}

func (v *Vampire) isValidVampMove(direction string, upperBound int) bool {
	newX, newY := v.x, v.y

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

	return newX >= 0 && newX < upperBound && newY >= 0 && newY < upperBound
}

func (v Vampire) String() string {
	return fmt.Sprintf("v: (%d, %d) \n", v.x, v.y)
}
