package main

import (
	"fmt"
	"l/dungeon"
	"l/minesweeper"
	"os"

	"golang.org/x/term"
)

func main() {
	fmt.Println("Choose which game you want to play: \n\r")
	fmt.Println("1. Dungeon")
	fmt.Println("2. Minesweeper")

	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		dungeon.InitGame()
	case "2":
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		minesweeper.InitGame()
	default:
		fmt.Println("Invalid choice.")

	}
}
