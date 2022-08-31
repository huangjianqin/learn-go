package main

import "fmt"

func main() {
	println("1-------")
	m := map[string]string{"a": "A", "b": "B", "c": "C"}
	fmt.Printf("%v\n", m)

	println("2-------")
	dk := "d"
	dv, dr := m[dk]
	fmt.Printf("%v\n", dv)
	fmt.Printf("%v\n", dr)

	println("3-------")
	m1 := map[string]string{}
	fmt.Printf("%v\n", m1)
	m1["E"] = "e"
	m1["G"] = "g"
	m1["K"] = "k"
	m1["I"] = "i"
	fmt.Printf("%v\n", m1)
}
