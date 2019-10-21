package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := range A {
		f, s := 0, len(A[i])-1
		for ; f < s; f, s = f+1, s-1 {
			A[i][f] = A[i][f] ^ 1
			A[i][s] = A[i][s] ^ 1
			A[i][f], A[i][s] = A[i][s], A[i][f]
		}

		if len(A[i])%2 != 0 {
			A[i][s] = A[i][s] ^ 1
		}
	}
	return A
}
