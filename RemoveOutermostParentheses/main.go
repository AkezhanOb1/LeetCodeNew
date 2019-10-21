package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}

func removeOuterParentheses(S string) string {
	var counting bool
	var after int

	var sb strings.Builder

	for i := range S {
		if S[i] == '(' {
			if counting {
				after++
				sb.WriteByte(S[i])
				continue
			} else {
				counting = true
				continue
			}
		}
		if counting && S[i] == ')' {
			if after == 0 {
				counting = false
				continue
			} else {
				after--
				sb.WriteByte(S[i])
				continue
			}
		}
	}

	return sb.String()
}
