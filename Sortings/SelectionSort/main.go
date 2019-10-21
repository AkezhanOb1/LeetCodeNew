package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 2, 4}
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	for i := range arr {
		minP := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minP] > arr[j] {
				minP = j
			}
		}
		arr[i], arr[minP] = arr[minP], arr[i]
	}
}
