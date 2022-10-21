package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github/gorilla/mux"
)


type Movie struct{
ID string `json:"id"`
Isbn string `json:"isbn"`
Title string `json:"title"`
Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Last string `json:"lastname"`
}


var movies [] Movie

func main() {
	r:=mux.NewRouter()
	movies = append(movies,)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
  