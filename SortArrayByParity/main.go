package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	srtArr := []int{}
	for _, val := range A {
		if val&1 == 0 {
			srtArr = append(srtArr, val)
		}
	}
	for _, val := range A {
		if val&1 != 0 {
			srtArr = append(srtArr, val)
		}
	}
	return srtArr
	// for i, iVal := range A {
	// 	for j, jVal := range A {
	// 		if iVal&1 == 0 && jVal&1 != 0 {
	// 			A[i], A[j] = A[j], A[i]
	// 		}
	// 	}
	// }
	// return A
}
