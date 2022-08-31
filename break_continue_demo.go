package main

import "strconv"

func main() {
	//f3()
	f4()
}

func f3() {
	//相当于简单直接使用continue
	for i := 0; i < 5; i++ {
	l:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue l
				//continue
			}

			println(strconv.Itoa(i) + "-" + strconv.Itoa(j))
		}
	}
}

func f4() {
l:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				break l
			}

			println(strconv.Itoa(i) + "-" + strconv.Itoa(j))
		}
	}
}
