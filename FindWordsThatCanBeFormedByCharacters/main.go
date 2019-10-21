package main

import (
	"fmt"
)

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}

func countCharacters(words []string, chars string) int {
	count := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		count[chars[i]-'a']++
	}

	ans := 0
	for _, w := range words {
		if canBeFormed(w, count) {
			ans += len(w)
		}
	}
	return ans
}
func canBeFormed(w string, c []int) bool {
	count := make([]int, 26)
	for i := 0; i < len(w); i++ {
		count[w[i]-'a']++
		if count[w[i]-'a'] > c[w[i]-'a'] {
			return false
		}
	}
	return true
}
