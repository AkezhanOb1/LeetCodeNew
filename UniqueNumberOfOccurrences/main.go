package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 1, 2, 2, 2, 3}))
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, i := range arr {
		if _, ok := m[i]; ok {
			m[i]++
		} else {
			m[i] = 1
		}
	}

	c := make(map[int]bool)
	for _, val := range m {
		if _, ok := c[val]; ok {
			return false
		}
		c[val] = true
	}
	return true
}

/*
	func uniqueOccurrences(arr []int) bool {
		m := make(map[int]int)
		for _, i := range arr {
			if _, ok := m[i]; ok {
				m[i]++
			} else {
				m[i] = 1
			}
		}
		a := []int{}
		for _, val := range m {
			a = append(a, val)
		}
		sort.Ints(a)

		for i := 0; i < len(a)-1; i++ {
			if a[i] == a[i+1] {
				return false
			}
		}
		return true
	}
*/
