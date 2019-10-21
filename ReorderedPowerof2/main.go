package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reorderedPowerOf2(10))
}

func reorderedPowerOf2(N int) bool {
	strN := strconv.Itoa(N)
	permArr := permutations(strN)
	fmt.Println(permArr)
	for _, val := range permArr {
		x, _ := strconv.Atoi(val)
		if val[0] == '0' {
			continue
		}
		if ok := isPowerOf2(x); ok {
			return true
		}
	}
	return false

}

func isPowerOf2(x int) bool {
	return (x & (x - 1)) == 0
}

func join(ins []rune, c rune) []string {
	result := []string{}
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return result
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		}
		result := []string{}
		for _, e := range p {
			result = append(result, join([]rune(e), testStr[0])...)
		}
		return n(testStr[1:], result)

	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}
