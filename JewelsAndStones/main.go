package main

import (
	"fmt"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(J string, S string) int {
	counter := 0
	m := make(map[string]int)
	for _, i := range S {
		letter := string(i)
		if _, ok := m[letter]; ok {
			m[letter]++
		} else {
			m[letter] = 1
		}
	}
	for _, i := range J {
		letter := string(i)
		if val, ok := m[letter]; ok {
			counter = counter + val
		}
	}
	return counter
}

/*
	func numJewelsInStones(J string, S string) int {
		counter := 0
		for _, j := range J {
			amount := strings.Count(S, string(j))
			counter = counter + amount
		}
		return counter
	}
*/
