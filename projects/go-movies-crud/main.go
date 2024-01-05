package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:  "id"`
	Title    string    `json: "title"`
	Isbn     string    `json: "isbn"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json : "firstname"`
	Lastname  string `json : "lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "12334", Title: "PS1", Director: &Director{Firstname: "Mani", Lastname: "Rathnam"}})
	movies = append(movies, Movie{ID: "1", Isbn: "12335", Title: "Jigarthanda", Director: &Director{Firstname: "Karthick", Lastname: "Subburaj"}})

	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting CRUD Movie web server at 8888\n")
	log.Fatal(http.ListenAndServe(":8888", r))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to movies database\n\n")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			fmt.Fprintf(w, "Deleted Boss....Movie gone\n")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Fprintf(w, "Movie Not found boss")

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			fmt.Printf("  Id found = %s", params["id"])
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	fmt.Fprintf(w, "Movie Not found boss,, what to update ??")
}
