package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}))
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]interface{})
	for _, email := range emails {
		str := strings.Split(email, "@")
		lName := str[0]
		dName := str[1]
		lName = strings.Replace(strings.Split(lName, "+")[0], ".", "", -1)
		fmt.Println(lName)
		email = lName + "@" + dName
		if _, ok := m[email]; !ok {
			m[email] = true
		}
	}
	return len(m)
}
