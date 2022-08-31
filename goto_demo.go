package main

import "strconv"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				goto END
			}

			println(strconv.Itoa(i) + "-" + strconv.Itoa(j))
		}
	}

END:
	println("goto end")
}
