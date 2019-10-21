package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("DDI"))
}

func diStringMatch(S string) []int {
	res := []int{}
	N := len(S)
	inc := 0
	dec := N
	for i := 0; i < N; i++ {
		if S[i] == 'I' {
			res = append(res, inc)
			inc++
		} else {
			res = append(res, dec)
			dec--
		}
	}

	return append(res, dec)
}
