package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	//at1()

}

func at1() {
	var n int32 = 100
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			atomic.AddInt32(&n, 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			atomic.AddInt32(&n, -1)
		}
	}()

	wg.Wait()
	println("-------")
	println(n)
}
