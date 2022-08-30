package main

import "strconv"

func main() {
	if age, name := 20, "aaa"; age > 10 {
		println(name + "-" + strconv.Itoa(age+10))
	} else {
		println("unsatisfied")
	}
}
