package main

import "fmt"

func main() {
	type ai = int
	var i ai = 10
	println(i)
	//虽然用了别名, 但是打印出来时一样的类型
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", 1)
}
