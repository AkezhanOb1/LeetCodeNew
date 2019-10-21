package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(A []string) []string {
	outterCounter := [26]int{}
	for i := range outterCounter {
		outterCounter[i] = math.MaxInt8
	}
	innerCounter := [26]int{}
	for _, word := range A {
		for _, char := range []byte(word) {
			innerCounter[char-'a']++
		}

		for i := 0; i < 26; i++ {
			if innerCounter[i] < outterCounter[i] {
				outterCounter[i] = innerCounter[i]
			}
		}
		for i := range innerCounter {
			innerCounter[i] = 0
		}
	}
	res := []string{}
	for i := range outterCounter {
		for j := 0; j < outterCounter[i]; j++ {
			res = append(res, string(i+'a'))
		}
	}

	return res
}
