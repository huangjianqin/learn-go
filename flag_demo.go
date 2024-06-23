package main

import (
	"flag"
	"log"
	"os"
)

var (
	src = flag.String("src", "", "source")
	dst = flag.String("dst", "", "destination")
)

func main() {
	flag.Parse()
	log.Printf("src: %v, dst: %v", *src, *dst)
	log.Printf("args: %v", os.Args)
}
