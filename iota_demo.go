package main

import (
	"learn-go/iotademo"
)

/*
 * iota是const(...)块中的常量定义index
 * 如果单纯多行声明常量, iota永远返回0, 例如
 * const a1 = iota //0
 * const a2 = iota //0
 */
const (
	a1 = iota
	_
	a2 = iota
)

func main() {
	iotademo.Hello()

	const (
		g1 = iota
		g2 = iota
	)

	println(a1)
	println(a2)
	println(g1)
	println(g2)
}
