package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"
)

type Movie struct {
	Id string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`
}

var movies []Movie

type Director struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...)
			break 
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)

	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		log.Printf("err: Movie encoding not successful....")
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			err := json.NewDecoder(r.Body).Decode(&movie)
			if err != nil {
				log.Fatal(err)
			}
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie {Id: "1", Isbn: "435-234-098", Title: "The Good Doctor", Director: &Director{Firstname: "Steve", Lastname: "Lee"}})
	movies = append(movies, Movie {Id: "2", Isbn: "477-274-098", Title: "One of Us is Lying", Director: &Director{Firstname: "Robert", Lastname: "Coulson"}})
	
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at 8000......")
	log.Fatal(http.ListenAndServe(":8000", r))



}



