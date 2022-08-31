package main

import "fmt"

func main() {
	//int指针
	println("1---------")
	var ip *int
	println(ip)
	i := 100
	//取地址
	ip = &i
	println(ip)
	//取值
	println(*ip)
	j := 100
	//两次100并不是同一实例
	println(ip == &j)

	println("2---------")
	ints := [3]int{1, 2, 3}
	//数组指针数组
	var pa [3]*[3]int
	pa[0] = &ints
	pa[1] = &ints
	pa[2] = &ints
	for _, v := range pa {
		fmt.Printf("%v\n", *v)
	}
}
