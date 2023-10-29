package main

import "fmt"

type Player struct {
	name string
	x    int
	y    int
}

func createPlayer(name string) Player {
	return Player{
		name: name,
		x:    0,
		y:    0,
	}
}

func (p *Player) move(direction string) {
	switch direction {
	case "w":
		p.x = p.x - 1
	case "a":
		p.y = p.y - 1
	case "s":
		p.x = p.x + 1
	case "d":
		p.y = p.y + 1
	default:
		fmt.Println("Invalid direction")
	}
}

func (p Player) String() string {
	return fmt.Sprintf("%s (%d, %d) \n", p.name, p.x, p.y)
}
