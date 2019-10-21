package main

import "log"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 4, 9, 16, 25}
	log.Println(same(a, b))
}

func same(a []int, b []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(b); i++ {
		m[b[i]] = 1
	}

	for i := 0; i < len(a); i++ {
		val := a[i] * a[i]
		if _, ok := m[val]; !ok {
			return false
		}
	}
	return true
}
