package main

import "log"

func main() {
	log.Println(singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for key, val := range m {
		if val == 1 {
			return key
		}
	}
	return 0
}
