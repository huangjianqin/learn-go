package main

import "bytes"

func main() {
	bs := []byte("1234asdff12321")
	println(bytes.Contains(bs, []byte("123")))
	println(bytes.Count(bs, []byte("123")))
}
