package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented!")
}

func FindMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented!")
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented!")
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented!")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented!")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", AllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", FindMovies).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovie).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
