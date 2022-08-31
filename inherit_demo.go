package main

type Animal struct {
	kind string
}

type Dog struct {
	Animal
	name string
}

func (an Animal) eat() {
	println(an.kind + " eating...")
}

func (d Dog) eat() {
	println(d.name + " eating...")
}

func main() {
	dog := Dog{Animal{"Dog"}, "ddd"}
	//不一样的方法, 不存在方法重写
	dog.eat()
	dog.Animal.eat()
}
