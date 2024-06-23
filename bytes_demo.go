package main

import (
	"bytes"
	"fmt"
)

func main() {
	bs := []byte("1234asdff12321")
	println(bytes.Contains(bs, []byte("123")))
	println(bytes.Count(bs, []byte("123")))

	bs1 := make([]byte, 16)
	fmt.Printf("len(bs1):%v , cap(bs1):%v\r\n", len(bs1), cap(bs1))
	bss1 := bs1[:8]
	fmt.Printf("len(bss1):%v , cap(bss1):%v\n", len(bss1), cap(bss1))
}
