package main

import (
	"os"
	"net/http"
	"io"
	"io/ioutil"
	"time"
	"fmt"
)

func main() {
	start := time.Now()

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Println("\nProgram ended in ", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	start_time := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("--> Error(%v) while getting %s", err, url)
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("--> Error(%v) while copy output(%s)", err, url)
		return
	}

	secs := time.Since(start_time).Seconds()
	ch <- fmt.Sprintf("Success! Time = %v,  Status = %v,  Bytes = %v. URL = %v", secs, resp.Status, nbytes, url)
}