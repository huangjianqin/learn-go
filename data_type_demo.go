package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unsafe"
)

func empty() string {
	return "empty"
}

func main() {
	name := "abc"
	a := 1
	b := false

	//整形
	println("整形----")
	fmt.Printf("%T\n", name)
	//输出值
	fmt.Printf("%v\n", name)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)

	//多少位int
	fmt.Printf("%v\n", unsafe.Sizeof(a))
	fmt.Printf("%v\n", math.MaxInt16)
	fmt.Printf("%v\n", math.MinInt16)

	//输出二进制
	fmt.Printf("%b\n", 4)
	fmt.Printf("%o\n", 4)
	fmt.Printf("%x\n", 4)
	//cast
	println(int(8.0) % 3)

	//浮点
	println("浮点----")
	fmt.Printf("%.2f\n", 3.14159)

	//10个循环后, 变得不精准
	ratio := 1.0 / 10.0
	for range []int{10: 0} {
		ratio += 1.0 / 10.0
	}
	fmt.Printf("%.60f\n", ratio)

	//指针
	println("指针----")
	ap := &a
	fmt.Printf("%T\n", ap)
	fmt.Printf("%v\n", ap)

	//切片
	println("切片----")
	ints := []int{1, 2, 3}
	fmt.Printf("%T\n", ints)
	fmt.Printf("%v\n", ints)

	//func
	println("func----")
	fmt.Printf("%T\n", empty)
	fmt.Printf("%v\n", empty())

	//string
	println("string----")
	println(`a
			b
			c
			d
			f`)
	println(fmt.Sprintln("ab", "dc", "ef"))
	println(strings.Join([]string{"ab", "dc"}, ","))
	var buffer bytes.Buffer
	buffer.WriteString("ab")
	buffer.WriteString("1")
	buffer.WriteString("cd")
	println(buffer.String())
	s := "abc 1111 2222 "
	fmt.Printf("%c\n", s[2])
	//[)
	fmt.Printf("%v\n", s[:5])
	fmt.Printf("%v\n", s[5:])
	fmt.Printf("%v\n", s[1:3])

}
