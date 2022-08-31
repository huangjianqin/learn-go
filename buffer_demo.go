package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//b1()

	//b2()

	//b3()
}

func b1() {
	file, err := os.OpenFile("data/file.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		defer file.Close()
		br := bufio.NewReader(file)
		s, err := br.ReadString('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%v\n", s)
		}
	}
}

func b2() {
	file, err := os.OpenFile("data/file.txt", os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		defer file.Close()
		wr := bufio.NewWriterSize(file, 2)
		wr.WriteString("12321312312")
		wr.Flush()
	}
}

func b3() {
	sr := strings.NewReader("ABC12AXVZXC12ASDRASD12ASDFASF12")
	scanner := bufio.NewScanner(sr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("%v\n", string(scanner.Bytes()))
	}
}
