package interfacedemo

type Flyer interface {
	/*
		同理, 接口内部定义的函数, 大写即表示其他package可以访问
	*/
	Fly()
}

type Swimmer interface {
	Swim()
}

//组合接口
type FLyFish interface {
	Flyer
	Swimmer
}
