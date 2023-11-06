package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	done := make(chan int)

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Copy done in go routine\n")
		done <- 0 // Signal main
	}()

	mustCopy(conn, os.Stdin) // to send stdin to server
	conn.Close()
	log.Println("wait for go routine to finish")

	<-done // wait for go routine to finish
	log.Println("Done wait, exiting")
}

func mustCopy(dst io.Writer, src io.Reader) {
	log.Printf("Must Copy called\n")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	log.Printf(" Copied \n")
}
