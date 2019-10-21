package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 4}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	for i := len(arr); i > 0; i-- {
		swaps := true
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps = false
			}
		}
		if swaps == true {
			break
		}
	}
}
