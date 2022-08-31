package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	io1()
}

func io1() {
	sr := strings.NewReader("abcs1231\r\n")
	lr := io.LimitReader(sr, 4)
	file, err := os.OpenFile("data/file.txt", os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("%v/n", err)
		return
	}

	//共享指针, 所以lr copy完, file指针在file末尾, 故无法复制file的内容
	//lr = io.MultiReader(lr, file)
	//copy有缓冲区, 所以不会无限复制file内容
	lr = io.MultiReader(file, lr)
	n, err := io.Copy(file, lr)
	if err != nil {
		fmt.Printf("%v/n", err)
		return
	} else {
		println("copy " + strconv.FormatInt(n, 10))
	}

	defer file.Close()
}
