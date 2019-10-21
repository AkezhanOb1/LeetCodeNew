package main

import (
	"fmt"
)

func main() {
	fmt.Println(selfDividingNumbers(2, 22))
	fmt.Println(devider(21))
}

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		if ok := devider(i); ok {
			res = append(res, i)
		}
	}
	return res
}

func devider(a int) bool {
	num := a
	for num > 0 {
		digit := num % 10
		if digit == 0 || a%digit != 0 {
			return false
		}
		num = num / 10
	}
	return true
}
