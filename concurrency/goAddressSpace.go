package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for _, txt := range []string{"Hello", "bye", "PRINT"} {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			log.Printf("\n%v", txt)
		}()
	}

	wg.Wait()
}
