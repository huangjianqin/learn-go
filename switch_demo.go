package main

import "math/rand"

func main() {
	num := rand.Intn(20)

	println(num)

	//值匹配
	switch num {
	case 10:
		println("10")
	case 20:
		println("20")
	case 30:
		println("30")
	case 40:
		println("40")
	default:
		println("值匹配失败")
	}

	//表达式匹配
	switch {
	case 0 < num && num <= 10:
		println("(0,10]")
		//可以执行满足条件的下一个case(不管它是否满足条件)
		fallthrough
	case 10 < num && num <= 20:
		println("(10,20]")
		fallthrough
	case 20 < num && num <= 30:
		println("(20,30]")
		break
	case 30 < num && num <= 40:
		println("(30,40]")
		fallthrough
	default:
		println(">40")
	}
}
