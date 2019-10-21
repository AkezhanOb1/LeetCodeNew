package main

import "fmt"

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

func uniqueMorseRepresentations(words []string) int {
	morseList := []string{}
	for _, word := range words {
		morseWord := ""
		for _, letter := range word {
			morseWord = morseWord + getMorseSign(letter)
		}
		morseList = append(morseList, morseWord)
	}
	counter := make(map[string]int)
	for _, row := range morseList {
		counter[row]++
	}
	distinctStrings := make([]string, len(counter))
	i := 0
	for k := range counter {
		distinctStrings[i] = k
		i++
	}
	return i
}

func getMorseSign(a rune) string {
	return morse[a-97]
}

/*
	func uniqueMorseRepresentations(words []string) int {
		m := make(map[string]bool)
		for _, word := range words {
			morseWord := ""
			for _, letter := range word {
				morseWord = morseWord + getMorseSign(letter)
			}

			if _, ok := m[morseWord]; !ok {
				m[morseWord] = true
			}
		}

		return len(m)
	}
*/
