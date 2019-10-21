package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{7, 3, 1, 0, 0, 6}))
}

func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}
	return sum
}
