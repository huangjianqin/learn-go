package interfacedemo

type Usb interface {
	Read()
	Write()
}

type Person struct {
	name string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) Read() {
	println(p.name + " read...")
}

func (p Person) Write() {
	println(p.name + " write...")
}
