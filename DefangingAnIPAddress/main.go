package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("222.222.22.22"))
}

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
