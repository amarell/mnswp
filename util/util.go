package util

import (
	"fmt"
)

func ArrayContains[T comparable](slice []T, element T) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}

func CleanTerminal() {
	fmt.Print("\033[H\033[2J")
}
