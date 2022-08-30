package iotademo

const (
	a1 = iota
	_
	a2 = iota
)

func Hello() {
	println(a1)
	println(a2)
}
