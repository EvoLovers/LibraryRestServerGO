package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author *Author   `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

/*
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}*/

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
} /*
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var err error
	for _, item := range books {
		item.ID, err = uuid.Parse(params["id"])
		if err == nil {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
*/

func getBook(c *gin.Context) {
	id := c.Param("ID")
	var err error
	for _, a := range books {
		a.ID, err = uuid.Parse(id)
		if err == nil {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

/*
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = uuid.New()
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
*/

func createBook(c *gin.Context) {

	var newBook Book
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new album to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var err error
	for index, item := range books {
		item.ID, err = uuid.Parse(params["id"])
		if err == nil {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID, err = uuid.Parse(params["id"])
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

/*
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var err error
	for index, item := range books {
		item.ID, err = uuid.Parse(params["id"])
		if err == nil {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
*/

func deleteBook(c *gin.Context) {
	id := c.Param("ID")

	var err error
	var books1 []Book
	for index, item := range books {
		item.ID, err = uuid.Parse(id)
		fmt.Print(item.ID)
		if err == nil {

			if id != item.ID.String() {

				books = append(books1[:index], books1[index+1:]...)
			}
			break
		} else {
			fmt.Println(err)
		}

	}
	books = books1
	c.IndentedJSON(http.StatusOK, books)

}

func main() {
	r := gin.Default()
	books = append(books, Book{ID: uuid.New(), Title: "Денискины рассказы", Author: &Author{Firstname: "Виктор", Lastname: "Драгунский"}})
	books = append(books, Book{ID: uuid.New(), Title: "Маленький принц", Author: &Author{Firstname: "Антуан", Lastname: "де Сент-Экзюпери"}})
	/*r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE") */
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBook)
	r.POST("/newBook", createBook)
	r.DELETE("/deleteBook/:id", deleteBook)
	//	fs := http.FileServer(http.Dir("./files"))

	//	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", fs))
	//http.Handle("/files/", fs)
	log.Fatal(http.ListenAndServe(":8000", r))
}
