package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf(" Starting Our Simple Webserver...")
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hellooo", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf(" Starting server at port 9999\n ")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found path dude...", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported dude.....", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found path", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error : %v", err)
	}

	fmt.Fprintf(w, "POST request success......")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "\nName = %s, \n Address = %s", name, address)
}
