package main

import (
	"fmt"
	"time"
)

var or func(channels ...<-chan interface{}) <-chan interface{}

func main() {

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		fmt.Printf("\n OR Started ... %v", len(channels))
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})

		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}

			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()

	<-or(sig(1*time.Minute),
		sig(1*time.Hour),
		sig(3*time.Minute),
		sig(2*time.Minute),
		sig(10*time.Second))

	fmt.Printf("\n Done after %v", time.Since(start))
}
