package main

import (
	"fmt"
	"math/rand"
)

type Vampire struct {
	x int
	y int
}

var moves = []string{"w", "a", "s", "d"}

func (v *Vampire) move() {
	switch randomMove := moves[rand.Intn(len(moves))]; randomMove {
	case "w":
		v.y = v.y - 1
	case "a":
		v.x = v.x - 1
	case "s":
		v.y = v.y + 1
	case "d":
		v.x = v.x + 1
	default:
		fmt.Println("Something went seriously wrong")
	}
}
