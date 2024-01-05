package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rajuuu1992/LearnGo/bookstore/pkg/routes"
)

func main() {
	fmt.Printf("Bookstore opened")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("\nStarting web server\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
