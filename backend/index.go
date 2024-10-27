package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/home" {
// 		http.Error(w, "404 Not found", http.StatusNotFound)
// 		return
// 	}l
// 	fmt.Fprint(w, "HEllo world")
// }
// func bookHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "books")
// }

// func main() {
// 	fileServerHandler := http.FileServer(http.Dir("./static"))
// 	// Returns the handler
// 	fmt.Println(fileServerHandler)
// 	http.Handle("/", fileServerHandler)
// 	http.HandleFunc("/home", homeHandler)
// 	http.HandleFunc("/book", bookHandler)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

type Movie struct {
	Id       int      `json:id`
	Title    string   `json:title`
	Director Director `json:director`
}

type Director struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

// slice

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
	w.WriteHeader(http.StatusOK)
}

var movies []Movie

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//length of a movie iterate and then select a id
	vars := mux.Vars(r)
	id := vars["id"]
	ID, _ := strconv.Atoi(id)
	for i := 0; i < len(movies); i++ {
		if movies[i].Id == ID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(movies[i])
		}
	}
	// fmt.Fprint(w, http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Parse int
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])
	// Remove it from the slice we have
	movies = append(movies[:Id], movies[Id+1:]...)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	// get the data from the body and append to the movie slice
	var new_movie *Movie
	// create a transformation pipeline by decoding
	json.NewDecoder(r.Body).Decode(new_movie)
	fmt.Println(new_movie)
	movies = append(movies, *new_movie)

	w.Header().Set("Content-type", "application/json")
	//Encode GO values into json
	//1. Create a transformation pipeline by encoding and streaming to w response
	json.NewEncoder(w).Encode(movies)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	router := mux.NewRouter()
	movies = append(movies, Movie{Id: 1, Title: "Interstellar", Director: Director{Firstname: "Chris", Lastname: "Nolan"}})
	movies = append(movies, Movie{Id: 2, Title: "Inception", Director: Director{Firstname: "Chris", Lastname: "Nolan"}})
	movies = append(movies, Movie{Id: 3, Title: "Avengers", Director: Director{Firstname: "Russo", Lastname: "Brothers"}})
	router.HandleFunc("/movies", AddMovie).Methods("POST")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
