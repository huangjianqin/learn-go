package main

func main() {
	d1()
}

func d1() {
	//函数结束时执行defer, 并且defer按stack形式存储, 故最后定义的先执行. 有点像java的finally
	println("start")
	defer println("step1")
	defer println("step2")
	defer println("step3")
	println("end")
}
