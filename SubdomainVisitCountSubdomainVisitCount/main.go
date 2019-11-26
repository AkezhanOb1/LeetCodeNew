package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

func subdomainVisits(cpdomains []string) []string {
	visitCounter := []string{}
	vsMapper := make(map[string]int)
	for _, domain := range cpdomains {
		index := strings.Index(domain, " ")
		visit, _ := strconv.Atoi(domain[:index])
		domain = domain[index+1:]
		i := 0
		for i != -1 {
			i = strings.Index(domain, ".")
			if val, ok := vsMapper[domain]; ok {
				vsMapper[domain] = val + visit
			} else {
				vsMapper[domain] = visit
			}
			domain = domain[i+1:]
		}
	}

	for key, val := range vsMapper {
		visitCounter = append(visitCounter, fmt.Sprintf("%d %s", val, key))
	}
	return visitCounter
}
