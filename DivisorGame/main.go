package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(divisorGame(4))
}

func divisorGame(N int) bool {
	if N == 1 {
		return false
	}
	result := false
	for N > 1 {
		move := nextMove(N)
		N = N - move
		result = !result
	}
	return result
}

func nextMove(max int) int {

	randVal := rand.Intn(max-1) + 1
	for max%randVal != 0 {
		randVal = rand.Intn(max-1) + 1
	}
	return randVal
}
