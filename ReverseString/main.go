package main

import "fmt"

func main() {
	arr := []byte("Hello")
	reverseString(arr)
	fmt.Println(string(arr))
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
