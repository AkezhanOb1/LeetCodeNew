package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	var str bytes.Buffer
	for _, val := range arr {
		str.WriteString(reverseWord(val))
	}
	return strings.TrimSpace(str.String())
}

func reverseWord(word string) string {
	var str bytes.Buffer
	for i := len(word) - 1; i >= 0; i-- {
		str.WriteByte(word[i])
	}
	return str.String() + " "
}
