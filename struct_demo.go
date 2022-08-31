package main

import "fmt"

type Pet struct {
	name string
	kind string
}

type Person struct {
	name string
	age  int
	pet  Pet
}

func main() {
	println("1-----------")
	p := Person{"aaa", 11, Pet{"ddd", "dog"}}
	//p := Person{name: "aaa", age: 11}
	fmt.Printf("%v\n", p)
	fmt.Printf("%T\n", p)

	println("2-----------")
	as := struct {
		a int
		b int
	}{a: 1, b: 2}
	fmt.Printf("%v\n", as)
	fmt.Printf("%T\n", as)

	println("3-----------")
	var pp *Person = &p
	fmt.Printf("%v\n", pp)
	fmt.Printf("%v\n", *pp)
	fmt.Printf("%T\n", *pp)
	//可省略*
	fmt.Printf("%v\n", pp.name)

	println("4-----------")
	p1 := Person{"aaa", 11, Pet{"ddd", "dog"}}
	s1(p1)
	fmt.Printf("%v\n", p1)
	s2(&p1)
	fmt.Printf("%v\n", p1)

	println("5-----------")
	p.pet.call()
}

//值传递
func s1(person Person) {
	person.age += 10
	person.pet.name = "ddda"
	fmt.Printf("%v\n", person)
}

//指针传递
func s2(person *Person) {
	person.age += 10
	person.pet.name = "ddda"
	fmt.Printf("%v\n", *person)
}

//Pet所属方法
func (p Pet) call() {
	println(p.name + " wang wang")
}
