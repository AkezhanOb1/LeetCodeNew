package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 3, 4}))
}

func distributeCandies(candies []int) int {
	candyKinds := make(map[int]int)
	for _, candy := range candies {
		candyKinds[candy]++
	}
	counter := 0
	for range candyKinds {
		counter++
	}
	half := len(candies) / 2
	if counter > half {
		counter = counter - (counter - half)
	}

	return counter
}
