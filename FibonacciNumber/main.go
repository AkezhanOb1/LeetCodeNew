package main

import "fmt"

func main() {
	fmt.Println(fibRec(4))
}

func fibRec(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fibRec(N-1) + fibRec(N-2)
}

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	counter := 1
	last := 0
	for i := 2; i <= N; i++ {
		temp := counter
		counter = counter + last
		last = temp
	}

	return counter
}
