package main

import (
	"fmt"
	"l/dungeon"
	"l/minesweeper"
)

func main() {
	fmt.Println("Choose which game you want to play: ")
	fmt.Println("1. Dungeon")
	fmt.Println("2. Minesweeper")

	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		dungeon.InitGame()
	case "2":
		minesweeper.InitGame()
	default:
		fmt.Println("Invalid choice.")

	}
}
