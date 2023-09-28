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

type Movies struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movie []Movies

func main() {
	r := mux.NewRouter()

	movie = append(movie, Movies{ID: "1", Isbn: "420985", Title: "Eze Goes To School", Director: &Director{Firstname: "Peter", Lastname: "Obi"}})

	movie = append(movie, Movies{ID: "2", Isbn: "813703", Title: "Fela And The Kalakutas", Director: &Director{Firstname: "Jazzy", Lastname: "Don"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movie {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var moviee Movies
	_ = json.NewDecoder(r.Body).Decode(&moviee)
	moviee.ID = strconv.Itoa(rand.Intn(1000000))
	movie = append(movie, moviee)
	json.NewEncoder(w).Encode(moviee)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramss := mux.Vars(r)
	for index, item := range movie {
		if item.ID == paramss["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			var moviie Movies
			_ = json.NewDecoder(r.Body).Decode(&moviie)
			moviie.ID = paramss["id"]
			movie = append(movie, moviie)
			json.NewEncoder(w).Encode(moviie)
			return
		}
	}

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parmas := mux.Vars(r)
	for index, item := range movie {

		if item.ID == parmas["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}
