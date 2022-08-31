package main

import (
	"sync"
)

var n = 100

func main() {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			lock.Lock()
			increment()
			lock.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			lock.Lock()
			decrement()
			lock.Unlock()
		}
	}()

	wg.Wait()
	println("-------")
	println(n)
}

func increment() {
	n++
}

func decrement() {
	n--
}
