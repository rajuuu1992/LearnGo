package main

import "time"

func main() {

	msgs := make(chan string)
	uppers := make(chan string)

	messages := []string{"hello", "hi", "hw", "are", "you"}

	go func() {
		for m := range messages {
			time.Sleep(1 * time.Second)
			msgs <- m
		}
	}()

}
