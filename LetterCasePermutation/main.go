package main

import "fmt"

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
func letterCasePermutation(S string) []string {
	res := []string{S}
	for i, letter := range S {
		newLetter := checkLetter(letter)
		if newLetter != letter {
			res = append(res, S[:i]+string(newLetter)+S[i+1:])
		}
	}
	return res
}

func checkLetter(letter rune) rune {
	if letter > 64 && letter < 91 {
		return letter + 32
	} else if letter > 96 && letter < 123 {
		return letter - 32
	}
	return letter
}
