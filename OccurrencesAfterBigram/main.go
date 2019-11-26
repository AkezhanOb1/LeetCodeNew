package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOcurrences("ypkk lnlqhmaohv lnlqhmaohv lnlqhmaohv ypkk ypkk ypkk ypkk ypkk ypkk lnlqhmaohv lnlqhmaohv lnlqhmaohv lnlqhmaohv ypkk ypkk ypkk lnlqhmaohv lnlqhmaohv ypkk", "lnlqhmaohv", "ypkk"))
}

func findOcurrences(text string, first string, second string) []string {
	res := []string{}
	arr := strings.Fields(text)
	for i := 0; i < len(arr); i++ {
		if arr[i] == second && i != 0 {
			if arr[i-1] == first {
				if i+1 != len(arr) {
					res = append(res, arr[i+1])
				}
			}
		}
	}
	return res
}
