package main

import "learn-go/initdemo"

/*
1. init比main函数要先执行
2. 同一package(多个go源文件)可以有多个init函数, 执行顺序依赖于build顺序, 故同一package下的init函数逻辑不能相互依赖. 同一go源文件下的init函数按定义顺序执行

*/
func init() {
	println("init")
}

func main() {
	println("main")
	initdemo.Hello()
}
