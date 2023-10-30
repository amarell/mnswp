package util

import (
	"math/rand"
)

func Bogo(arr []int) []int {
	for !is_sorted(arr) {
		shuffle(arr)
	}

	return arr
}

func is_sorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

func shuffle(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-2; i++ {
		j := rand.Intn(n-i) + i
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

