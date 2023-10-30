package dungeon

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
		p.x--
	case "a":
		p.y--
	case "s":
		p.x++
	case "d":
		p.y++
	default:
		fmt.Println("Invalid direction")
	}
}

func (p Player) String() string {
	return fmt.Sprintf("%s (%d, %d) \n", p.name, p.x, p.y)
}
