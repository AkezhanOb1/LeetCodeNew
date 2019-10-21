package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
}

func dayOfTheWeek(day int, month int, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t.Weekday().String()
}
