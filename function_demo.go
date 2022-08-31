package main

import (
	"fmt"
	"math"
)

func main() {
	println("1---------")
	println(f5())

	println("2---------")
	ints := []int{1, 2, 3, 4, 5}
	f6(ints)
	fmt.Printf("%v\n", ints)

	println("3---------")
	f7(5, 6, 7, 8, 9)

	println("4---------")
	type fints func(...int)
	var ff fints = f7
	ff(ints...)

	println("5---------")
	f8(f7, ints...)

	println("6---------")
	f9(ints...)(ints...)

	println("7---------")
	max := func(a int, b int) int {
		return int(math.Max(float64(a), float64(b)))
	}
	println(max(2, 3))
	afr := func(a int, b int) int {
		return int(math.Max(float64(a), float64(b)))
	}(2, 3)
	println(afr)

	println("8---------")
	println(f10(5)(9))
}

func f5() (name string, age int) {
	return "abc", 10
}

func f6(ints []int) {
	ints[0] = 10
}

func f7(ints ...int) {
	for _, v := range ints {
		fmt.Printf("%v,", v)
	}
	println()
}

func f8(intsf func(...int), ints ...int) {
	intsf(ints...)
}

func f9(ints ...int) func(...int) {
	f7(ints...)
	return f7
}

func f10(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
