package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}

func sortArrayByParityII(A []int) []int {
	arr := make([]int, len(A))
	odd := 1
	even := 0
	for _, val := range A {
		if val&1 == 0 {
			arr[even] = val
			even += 2
		} else {
			arr[odd] = val
			odd += 2
		}
	}
	return arr
}
