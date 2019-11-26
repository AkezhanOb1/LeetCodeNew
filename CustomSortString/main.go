package main

import "fmt"

func main() {
	fmt.Println(customSortString("kqep", "pekeq"))
}

func customSortString(S string, T string) string {
	sortedPointer := 0
	sorterLetterPointer := 0
	for sorterLetterPointer != len(T)-1 && sorterLetterPointer < len(S) {
		for scanPointer := 0; scanPointer < len(T); scanPointer++ {
			if T[scanPointer] == S[sorterLetterPointer] {
				b := []byte(T)
				b[sortedPointer], b[scanPointer] = b[scanPointer], b[sortedPointer]
				T = string(b)
				sortedPointer++
			}
			fmt.Println(sorterLetterPointer)
		}
		sorterLetterPointer++
	}
	return T
}
