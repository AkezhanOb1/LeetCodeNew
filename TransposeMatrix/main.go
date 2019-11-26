package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return nil
	}
	result := make([][]int, len(A[0]), len(A[0]))
	for _, row := range A {
		for j, cell := range row {
			result[j] = append(result[j], cell)
		}
	}

	return result
}
