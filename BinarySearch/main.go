package main

import "fmt"

func main() {
	fmt.Println(biSearch([]int{1, 2, 3, 5, 6, 7, 8, 9, 10}, 110))
}

func biSearch(arr []int, n int) int {
	lPointer := 0
	rPointer := len(arr) - 1
	mPointer := 0
	for lPointer <= rPointer {
		mPointer = (lPointer + rPointer) / 2
		if arr[mPointer] == n {
			return mPointer
		}
		if arr[mPointer] > n {
			rPointer = mPointer - 1
		} else {
			lPointer = mPointer + 1
		}
	}

	return -1
}
