package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {

	defer printStack()
	fmt.Println(" Main ", sum(1, 2, 3, 4, 5))

	slowOp()

	readFiles("hello.txt", "new1.txt", "new2.txt")

	panicTest(7)
	log.Printf("\n\n......................Safe Exit..........")
}

func sum(elems ...int) int {
	sum := 0
	for s := range elems {
		sum += s
	}
	return sum
}

func slowOp() {

	time.Sleep(2 * time.Second)
	defer trace("slowOp")()
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("--- %s", msg)

	return func() {
		log.Printf("Exit %s (%s)", msg, time.Since(start))
	}
}

func readFiles(files ...string) {

	for _, file := range files {
		readFile(file)
	}
}

func readFile(file string) {

	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer func() {
		log.Printf(" Closing file %s", file)
		f.Close()
	}()
}

func panicTest(num int) {
	_ = 10 / num
	log.Printf(" Panic test %d", num)
	defer func() {
		log.Printf("Defer %d\n and recover", num)
		if p := recover(); p != nil {
			log.Printf("Recover successfull at %d: Recovery = %v", num, p)
		}

	}()
	panicTest(num - 1)
}

func printStack() {
	var buf [4096]byte

	n := runtime.Stack(buf[:], false)
	log.Printf("%v", string(buf[:n]))
	// os.Stdout.Write(buf[:n])
}
