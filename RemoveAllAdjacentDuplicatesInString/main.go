package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {
	for i, j := 0, 1; i < len(S)-1; i, j = i+1, j+1 {
		if S[i] == S[j] {
			S = S[:i] + S[j+1:]
			i = -1
			j = 0
		}
	}
	return S
}
