package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup

	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}

	send := func() {
		defer wg.Done()

		<-begin

		for i := 0; i < b.N; i++ {
			c <- token
		}
	}

	recv := func() {
		defer wg.Done()

		<-begin

		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)

	go send()
	go recv()

	b.StartTimer()
	close(begin)

	wg.Wait()

}
