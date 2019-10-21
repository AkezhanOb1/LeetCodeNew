package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
}

func repeatedNTimes(A []int) int {
	N := len(A) / 2
	m := make(map[int]int)
	for _, v := range A {
		if _, ok := m[v]; ok {
			m[v]++
			val := m[v]
			if val == N {
				return v
			}
		} else {
			m[v] = 1
		}
	}
	return 0
}
