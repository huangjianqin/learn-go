package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%v\n", now.Year())
	fmt.Printf("%v\n", now.YearDay())
	fmt.Printf("%v\n", now.Month())
	fmt.Printf("%v\n", now.Day())
	fmt.Printf("%v\n", now.Weekday())
	fmt.Printf("%v\n", now.Hour())
	fmt.Printf("%v\n", now.Minute())
	fmt.Printf("%v\n", now.Second())
	fmt.Printf("%v\n", now.Unix())
	fmt.Printf("%v\n", now.UnixMilli())
	fmt.Printf("%v\n", now.UTC())

	nowAdd := now.AddDate(1, 1, 1)
	fmt.Printf("%v\n", nowAdd.UTC())

	fmt.Printf("%v\n", nowAdd.After(now))
	fmt.Printf("%v\n", nowAdd.Sub(now))
}
