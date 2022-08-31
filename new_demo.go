package main

func main() {
	b := new(bool)
	println(*b)

	i := new(int)
	println(*i)

	intsP := new([]int)
	ints := *intsP
	//ints是长度为0的数组, 未初始化, 赋值会异常, 需使用make, 创建数组同时初始化, make只对slice array和map类型有效
	//println(ints)
	//ints[0] = 1
	//ints[1] = 2
	ints1 := append(ints, 1, 2, 3)
	println(ints1)
	println(&ints1)
}
