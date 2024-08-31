package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 1)

	go func() {
		defer close(intStream)

		// defer fmt.Fprintln(&stdoutBuff, "Hello..Done")
		defer fmt.Printf("...\n HELLO DONE... \n")

		for i := 0; i < 5; i++ {
			// fmt.Fprintf(&stdoutBuff, "Sending ... %d\n", i)
			// time.Sleep(2 * time.Second)

			intStream <- i
			fmt.Printf("...\n Sent... %d\n", i)
		}
	}()

	for integer := range intStream {
		// fmt.Fprintf(&stdoutBuff, "Received %d\n", integer)
		fmt.Printf("\n Received %d\n", integer)
	}
}
