package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 7
	}()
	go func() {
		ch <- 8
	}()
	time.Sleep(time.Minute)
	fmt.Println(<-ch)
}
