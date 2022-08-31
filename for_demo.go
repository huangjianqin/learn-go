package main

import "strconv"

func main() {
	f1()
	f2()
}

func f1() {
	m := make(map[string]string)
	m["a"] = "A"
	m["b"] = "B"
	m["c"] = "C"

	for k, v := range m {
		println(k + "-" + v)
	}
}

func f2() {
	s := "abcdqeqw"
	for i, c := range s {
		println(strconv.Itoa(i) + "-" + string(c))
	}
}
