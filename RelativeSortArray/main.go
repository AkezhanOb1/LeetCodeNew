package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	pointer := 0
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr1); j++ {
			if arr1[j] == arr2[i] {
				arr1[pointer], arr1[j] = arr2[i], arr1[pointer]
				pointer++
			}
		}
	}

	sortRest(arr1[pointer:])
	return arr1
}

func sortRest(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}
