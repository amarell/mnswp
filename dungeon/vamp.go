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
	randomMove := moves[rand.Intn(len(moves))]

	for !v.isValidVampMove(randomMove, upperBound) {
		randomMove = moves[rand.Intn(len(moves))]
	}

	switch randomMove {
	case "w":
		v.x = v.x - 1
	case "a":
		v.y = v.y - 1
	case "s":
		v.x = v.x + 1
	case "d":
		v.y = v.y + 1
	default:
		fmt.Println("Something went seriously wrong")
	}

	return v
}

func (v *Vampire) isValidVampMove(direction string, upperBound int) bool {
	switch direction {
	case "w":
		if v.x == 0 {
			return false
		}
	case "s":
		if v.x == upperBound-1 {
			return false
		}
	case "a":
		if v.y == 0 {
			return false
		}
	case "d":
		if v.y == upperBound-1 {
			return false
		}
	}

	return true
}

func (v Vampire) String() string {
	return fmt.Sprintf("v: (%d, %d) \n", v.x, v.y)
}
