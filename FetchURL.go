package main

import (
	"fmt"
	// "io/ioutil"
	"io"
	"strings"
	"os"
	"net/http"
)


func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
			fmt.Println("ADDED.........PREFIX")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while get for %s, : %v", url, err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()

		// if err != nil {
		// 	fmt.Println("Error while get for %s, : %v", url, err)
		// 	os.Exit(1)
		// }
		dst := os.Stdout
		io.Copy(dst, resp.Body)

		fmt.Println("-----> %v", dst)
		fmt.Println("---->  Res status = %v", resp.Status)
	}
}