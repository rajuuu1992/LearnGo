package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Accept loop")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// go handleConn(conn)
		go handleConnEcho(conn)
	}
}

func handleConn(c net.Conn) {
	log.Printf("handle Conn")
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func handleConnEcho(c net.Conn) {
	log.Printf("handle Conn")
	defer c.Close()
	input := bufio.NewScanner(c)
	log.Printf("Enter input to send to client...")
	for input.Scan() {
		go echo(c, input.Text(), 2*time.Second)
	}
}

func echo(c net.Conn, str string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(str))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", str)

	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(str))
}
