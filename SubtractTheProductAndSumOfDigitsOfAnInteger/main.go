package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	s := strconv.Itoa(n)
	sum := 0
	product := 1
	for _, digit := range s {
		digitI, _ := strconv.Atoi(string(digit))
		sum += digitI
		product *= digitI
	}
	return product - sum
}
