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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // mux.Vars(r) extracts URL parameters (like {id} in the route).

	// Loop through movies to find the one with the matching ID.
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)    // Decode the JSON request body into a Movie struct.
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // Generate a random ID for the new movie.
	movies = append(movies, movie)                // Append the new movie to the movies slice.
	json.NewEncoder(w).Encode(movie)              // Return the newly created movie as a JSON response.
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // mux.Vars(r) extracts URL parameters (like {id} in the route).

	// Loop through movies to find the one with the matching ID and update it.
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Remove the old movie.

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the JSON request body into a Movie struct.
			movie.ID = params["id"]                    // Ensure the ID remains the same.
			movies = append(movies, movie)             // Add the updated movie to the slice.
			json.NewEncoder(w).Encode(movie)           // Return the updated movie as a JSON response.
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // mux.Vars(r) extracts URL parameters (like {id} in the route).

	// Loop through movies to find the one with the matching ID and delete it.
	for index, item := range movies { // Loop through the movies slice.

		// If the movie ID matches the ID from the URL parameters.
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // Return the updated list of movies.
}

func main() {
	r := mux.NewRouter()

	// Mock Data
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454555", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
	// Route Handlers / Endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

//sruct and slices as db
