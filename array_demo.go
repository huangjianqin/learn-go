package main

import "fmt"

func main() {
	println("1---------")
	var ints []int
	//有无...都是一样的效果, 自动推断数组大小
	ints = []int{1: 1, 3: 3, 10: 10}
	fmt.Printf("%v\n", ints)
	println(len(ints))

	//固定长度, 会报错
	//ints = make([]int, 5)
	//ints[20] = 20
	//fmt.Printf("%v\n", ints)
	//println(len(ints))

	//切片
	println("2---------")
	var sints []int
	sints = ints[:]
	fmt.Printf("%v\n", sints)
	println(cap(sints))

	println("3---------")
	sints = ints[2:8]
	fmt.Printf("%v\n", sints)
	println(cap(sints))

	//切片与原数组共享同一份内存
	println("4---------")
	sints[3] = 4
	fmt.Printf("%v\n", sints)
	fmt.Printf("%v\n", ints)
	println(cap(sints))

	//切片增删, 相当于复制一份出来并修改
	println("5---------")
	ints1 := append(ints, 11)
	fmt.Printf("%v\n", ints1)
	fmt.Printf("%v\n", ints)

	println("6---------")
	var ints2 []int
	ints2 = append(ints2, ints1[:3]...)
	ints2 = append(ints2, ints1[4:]...)
	fmt.Printf("%v\n", ints2)
	fmt.Printf("%v\n", ints1)

	//复制
	println("7---------")
	var ints3 = make([]int, 5)
	//依据dst长度复制
	println(copy(ints3, ints1))
	ints3[2] = 2
	fmt.Printf("%v\n", ints3)
	fmt.Printf("%v\n", ints1)
}
