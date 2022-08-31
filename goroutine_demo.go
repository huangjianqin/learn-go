package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	println("cpu: " + strconv.Itoa(runtime.NumCPU()))

	//go showMsg("hello")
	//go showMsg("world")
	//time.Sleep(time.Millisecond * 1300)

	//g1()

	//g2()

	g3()

	//g4()

	//g5()

	//限制并发数
	//runtime.GOMAXPROCS(1)
	//g6()

}

func showMsg(msg string) {
	for i := 0; i < 10; i++ {
		println(time.Now().String(), msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func g1() {
	c := make(chan bool)
	defer close(c)

	go g11(c)
	<-c
	println("g1 out")
}
func g11(c chan bool) {
	time.Sleep(time.Second * 2)
	println("g11 out")
	c <- true
}

func g2() {
	c := make(chan int)
	q := make(chan bool)
	defer close(c)
	defer close(q)

	go g21(c, q)
	for {
		if j := <-c; j%2 == 0 {
			break
		}
	}
	//等待g21发送消息退出
	<-q
	println("g2 out")
}

/**
q定义只能发送
*/
func g21(c chan int, q chan<- bool) {
	for i := 0; i < 50; i++ {
		j := rand.Intn(10)
		println(j)
		c <- j
		if j%2 == 0 {
			break
		}
	}
	//通知g2退出
	q <- true
	println("g21 out")
}

func g3() {
	c := make(chan int)
	q := make(chan bool)

	go g31(c, q)
	for {
		r := rand.Intn(int(time.Second.Milliseconds()))
		time.Sleep(time.Duration(r))
		//如果receiver或sender还没准备好(closed也算是没准备好), 都会阻塞
		j, ok := <-c
		if j%2 == 0 || !ok {
			//通知g31退出
			q <- true
			close(q)
			break
		}
	}
	println("g3 out")
	time.Sleep(time.Millisecond * 500)
}

func g31(c chan int, q chan bool) {
out:
	for {
		j := rand.Intn(10)
		//伪随机的选择一个case处理
		//如果default case存在的情况下, 如果没有default case, 则select语句会阻塞, 直到某个case需要处理
		select {
		case c <- j:
			println(j)
		case <-q:
			//等待g31发送消息退出
			break out
		case <-time.After(time.Millisecond * 500):
			//500ms后超时
			println("time out")
			//结束g3
			c <- 2
			close(c)
		}
	}
	println("g31 out")
}

func g4() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()
	time.Sleep(time.Second * 3)
}

func g5() {
	var wg sync.WaitGroup
	num := 5
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(index int) {
			r := rand.Intn(5)
			time.Sleep(time.Duration(int(time.Second.Seconds()) * r))
			fmt.Println(strconv.Itoa(index) + "-finish")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("g5 out")
}

func g6() {
	go func() {
		for i := 0; i < 3; i++ {
			println("sub-" + strconv.Itoa(i))
			//本协程让出CPU
			//runtime.Gosched()

			if i == 1 {
				runtime.Goexit()
			}
		}
		println("sub finish")
	}()

	for i := 0; i < 3; i++ {
		println("main-" + strconv.Itoa(i))
		//本协程让出CPU
		runtime.Gosched()
	}

	println("main finish")

	time.Sleep(time.Second)
}
