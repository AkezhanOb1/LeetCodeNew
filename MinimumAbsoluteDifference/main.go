package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
}

func minimumAbsDifference(arr []int) [][]int {
	res := [][]int{}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	minDif := math.MaxInt64
	for i := 0; i < len(arr)-1; i++ {
		dif := arr[i+1] - arr[i]
		if dif < minDif && dif != 0 {
			minDif = dif
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		dif := arr[i+1] - arr[i]
		if dif == minDif {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}
