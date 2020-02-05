package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct - Model
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init Books Variables
var books []Book

// Getbooks Function
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Getbook Function
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through books
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func main() {
	//Init Router
	var r = mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Title: "First Book", Author: &Author{FirstName: "Test First Name", Lastname: "Test Last Name"}})
	books = append(books, Book{ID: "2", Title: "Second Book", Author: &Author{FirstName: "Test First Name", Lastname: "Test Last Name"}})

	// Route Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
