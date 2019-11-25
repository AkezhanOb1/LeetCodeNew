package main

import "fmt"

import "strings"

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}

func uncommonFromSentences(A string, B string) []string {
	diff := []string{}
	mapper := make(map[string]int)
	aArr := strings.Fields(A)
	bArr := strings.Fields(B)
	for _, el := range aArr {
		if _, ok := mapper[el]; !ok {
			mapper[el] = 1
		} else {
			mapper[el]++
		}
	}
	for _, el := range bArr {
		if _, ok := mapper[el]; !ok {
			mapper[el] = 1
		} else {
			mapper[el]++
		}
	}

	for key, val := range mapper {
		if val == 1 {
			diff = append(diff, key)
		}
	}
	return diff
}
