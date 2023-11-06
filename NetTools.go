package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println(" Usage ./NetTools.exe 192.168.1.1")
		return
	}

	addr, err := net.LookupHost(os.Args[1])
	fmt.Println(addr, err)
}
