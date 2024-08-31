package main

import (
	"fmt"
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
)

func main() {
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

	for num := range toString(done, take(done, gen(done, val), 10)) {
		fmt.Printf("\n --> Num = %v", num)
	}
}
