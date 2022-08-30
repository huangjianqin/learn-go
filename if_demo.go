package main

import (
	"math/rand"
	"strconv"
)

func main() {
	if age, desc := rand.Intn(40), "satisfied"; age > 10 {
		println(desc + "-" + strconv.Itoa(age))
	} else {
		println("unsatisfied-" + strconv.Itoa(age))
	}
}
