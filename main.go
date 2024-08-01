package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string
	Author string
	Price  string
}

var books = make(map[string]Book)

func AddBook(book Book) {
	books[book.ID] = book
}
func DeleteBook(id string) {
	delete(books, id)
}
func GetBookByID(id string) (Book, bool) {
	book, exists := books[id]
	return book, exists
}

func UpdateBook(id string, newBook Book) {
	books[id] = newBook
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, exists := GetBookByID(id)
	if !exists {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Unable to Add", http.StatusBadRequest)
		return
	}
	AddBook(book)
	w.WriteHeader(http.StatusCreated)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", GetHandler).Methods("GET")
	r.HandleFunc("/books/add", AddHandler).Methods("POST")
	fmt.Println("80808081")
	http.ListenAndServe(":8080", r)
}
