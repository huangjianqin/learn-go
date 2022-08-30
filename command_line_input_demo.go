package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age int
	var name string
	println("请输入age和name")
	fmt.Scan(&age, &name)
	println(name + "-" + strconv.Itoa(age))
}
