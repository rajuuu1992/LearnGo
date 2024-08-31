package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subs := func(c *sync.Cond, fn func()) {
		var goRoute sync.WaitGroup

		goRoute.Add(1)

		go func() {
			goRoute.Done()

			c.L.Lock()
			defer c.L.Unlock()

			c.Wait()
			fn()
		}()

		goRoute.Wait()
	}

	var clickReg sync.WaitGroup
	clickReg.Add(3)

	subs(button.Clicked, func() {
		fmt.Printf("\n Subscibe 1 called")
		clickReg.Done()
	})

	subs(button.Clicked, func() {
		fmt.Printf("\n Subscibe 2 called")
		clickReg.Done()
	})

	subs(button.Clicked, func() {
		fmt.Printf("\n Subscibe 3 called")
		clickReg.Done()
	})

	button.Clicked.Broadcast()

	clickReg.Wait()

	// sync once deadlock eg

	var onceA, onceB sync.Once

	var initB func()

	initA := func() { onceB.Do(initB) }

	initB = func() { onceA.Do(initA) }

	onceA.Do(initB)
}
