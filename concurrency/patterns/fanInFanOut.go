package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	genFn = func(done <-chan bool, fn func() interface{}) <-chan interface{} {

		numStream := make(chan interface{})

		go func() {
			defer close(numStream)

			for {
				select {
				case <-done:
					return

				case numStream <- fn():
					// fmt.Printf("\n Written %v\n", num)
				}
			}
		}()

		return numStream
	}

	gen = func(done <-chan bool, values []interface{}) <-chan interface{} {

		numStream := make(chan interface{})

		go func() {
			defer close(numStream)
			for {
				for _, val := range values {
					select {
					case <-done:
						return

					case numStream <- val:
						// fmt.Printf("\n Written %v\n", num)
					}
				}
			}
		}()

		return numStream
	}

	take = func(done <-chan bool, valStream <-chan interface{}, num int) <-chan interface{} {

		values := make(chan interface{})

		go func() {
			defer close(values)

			for i := 0; i < num; i++ {
				select {
				case <-done:
					return

				case values <- <-valStream:
				}
			}
		}()

		return values
	}

	toString = func(done <-chan bool, values <-chan interface{}) <-chan string {

		strs := make(chan string)

		go func() {
			defer close(strs)

			for v := range values {
				select {
				case <-done:
					return

				case strs <- v.(string):
				}
			}
		}()

		return strs
	}

	longFn = func(done <-chan bool, valStream <-chan interface{}) <-chan interface{} {

		vals := make(chan interface{})

		go func() {
			defer close(vals)

			for val := range valStream {
				select {
				case <-done:
					return
				case vals <- val:
					time.Sleep(2 * time.Second)
				}
			}
		}()

		return vals

	}

	fanIn = func(done <-chan bool, channels ...<-chan interface{}) <-chan interface{} {

		multiplexedOpStream := make(chan interface{})

		var wg sync.WaitGroup

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()

			for ch := range c {
				select {
				case <-done:
					return

				case multiplexedOpStream <- ch:
				}
			}
		}
		wg.Add(len(channels))

		for _, c := range channels {
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multiplexedOpStream)
		}()

		return multiplexedOpStream
	}
)

func main() {

	start := time.Now()
	// randd := func() interface{} {
	// 	return rand.Float64()
	// 	// return ret
	// }

	done := make(chan bool)

	// for num := range take(done, genFn(done, randd), 1) {
	// 	fmt.Printf("\n--> Num = %v\n", num)
	// }

	strs := []string{"a", "basdf", "asdfasdf", "cccc", "dddd"}
	// strs := []int{1, 2, 3, 43, 4, 5, 5, 4, 43, 3, 3, 3}
	val := make([]interface{}, len(strs))

	for i, s := range strs {
		val[i] = s
	}

	numFn := runtime.NumCPU()
	fns := make([]<-chan interface{}, numFn)

	for i := 0; i < numFn; i++ {
		fns[i] = longFn(done, gen(done, val))
	}

	fmt.Printf("\n -------- NUM FNS = %v", numFn)

	for num := range toString(done, take(done, fanIn(done, fns...), 10)) {
		fmt.Printf("\n --> Num = %v", num)
	}

	fmt.Printf("\n\n END . Time taken = %v", time.Since(start))
}
