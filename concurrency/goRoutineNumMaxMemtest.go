package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numGoRout = 1e5
	wg.Add(numGoRout)
	before := memConsumed()

	for i := numGoRout; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	after := memConsumed()

	fmt.Printf("%.3fkb ", float64(after-before)/numGoRout/1000)

	// syncgroup eg

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("\n... HELLO %v\n", id)
	}

	const num = 5
	var wwg sync.WaitGroup
	wwg.Add(num)
	for i := 0; i < num; i++ {
		go hello(&wwg, i+1)
	}
	wwg.Wait()

}
