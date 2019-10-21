package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	a := int64(x ^ y)
	return strings.Count(strconv.FormatInt(a, 2), "1")
}
