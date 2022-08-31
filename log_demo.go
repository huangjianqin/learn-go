package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	//指定输出内容
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//指定前缀
	log.SetPrefix("myLog ")
	logF, err := os.Create("data/log.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		log.SetOutput(logF)
	}
}

func main() {
	//l1()

	//l2()

	l3()
}

func l1() {
	defer println("defer end..")
	log.Println("my log")
	log.Panic(io.EOF)
	fmt.Println("end..")
}

func l2() {
	defer println("defer end..")
	log.Println("my log")
	//os.exit(1)
	log.Fatal("fatal")
	fmt.Println("end..")
}

func l3() {
	logF, err := os.Create("data/custom_log.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	logger := log.New(logF, "custom: ", log.LstdFlags|log.Llongfile)
	logger.Println("custom log")
}
