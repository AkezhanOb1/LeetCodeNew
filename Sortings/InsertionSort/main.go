package main

import "fmt"

func main() {
	arr := []int{2, 6, 9, 5, 4, 3, 2}
	insertionSort(arr)
	fmt.Println(arr)

}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
}
