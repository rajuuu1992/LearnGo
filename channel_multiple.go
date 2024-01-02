package main

import (
	"log"
	"net"
	"os"
	"strings"
	"time"

	"tailscale.com/net/interfaces"
)

func main() {

	defaultRoutes, _ := interfaces.DefaultRoute()

	pInterface, _ := net.InterfaceByIndex(defaultRoutes.InterfaceIndex)
	pAddrs, _ := pInterface.Addrs()
	log.Printf("Iface = %v, paddrs = %v ", pInterface, pAddrs)
	host, _ := os.Hostname()
	log.Printf(" Default routes = ", defaultRoutes)
	log.Println(" Host name = ", host)

	for _, addr := range pAddrs {
		addrStr := addr.String()

		var addrSuffix []string
		if strings.Contains(addrStr, "/") {
			addrSuffix, _ = net.LookupAddr(strings.Split(addrStr, "/")[0])

		} else {
			addrSuffix, _ = net.LookupAddr(addrStr)
		}
		log.Printf("Addrstr = ", addrStr)
		log.Print(" Iface = %v,  Suffix = %s", addr, addrSuffix)

	}
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

	iface, err := interfaces.DefaultRouteInterface()
	if err != nil {
		panic(err)
	}
	log.Println("Default iface = ", iface)

	log.Println(" Time = ", time.Now().UTC().Format(time.RFC3339))
}
