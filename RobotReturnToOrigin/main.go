package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("LL"))
}

func judgeCircle(moves string) bool {
	m := make(map[string]int)
	for _, letter := range moves {
		val := string(letter)
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}

	return (m["U"] == m["D"]) && (m["L"] == m["R"])
}
