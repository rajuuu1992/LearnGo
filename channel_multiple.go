package main

import (
	"log"
	"strings"
	"time"
)

func main() {

	msgs := make(chan string, 10)
	uppers := make(chan string, 10)

	messages := []string{"hello", "hi", "hw", "are", "you"}

	go func() {
		for _, m := range messages {
			time.Sleep(1 * time.Second)
			msgs <- m
		}
		close(msgs)
	}()

	go func() {
		for {
			msg, ok := <-msgs
			if !ok {
				log.Println(" 2: Channel recv error")
				break
			}
			uppers <- strings.ToUpper(msg)
		}
		close(uppers) // close already closed channel causes panic
	}()

	for up := range uppers {
		log.Println(up)
	}
	log.Println("Bye bye")
}
