package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//iu1()

	//iu2()
}

func iu1() {
	file, err := os.OpenFile("data/file.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%v\n", string(bytes))
		}
	}
}

func iu2() {
	//遍历指定目录下的文件或目录
	fis, err := ioutil.ReadDir("data")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for _, fi := range fis {
			println(fi.Name())
		}
	}
}
