package main

import (
	"fmt"
	"learn-go/interfacedemo"
)

func main() {
	println("1-----------")
	p := interfacedemo.Person{}
	p.SetName("A")
	fmt.Printf("%v\n", p)

	println("2-----------")
	fish := Fish{"飞鱼"}
	fmt.Printf("%v\n", fish)
	fish.Fly()
	fish.Swim()
}

type Fish struct {
	kind string
}

func (f Fish) Fly() {
	println(f.kind + " fly")
}

func (f Fish) Swim() {
	println(f.kind + " swim")
}
