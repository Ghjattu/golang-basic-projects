package controllers

import (
	"book-management-system/model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := model.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.ParseInt(mux.Vars(r)["bookID"], 0, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	book, _, err := model.GetBookByID(bookID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Record Not Found"))
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	createdBook, err := model.CreateBook(&book)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(createdBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook model.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)
	bookID, err := strconv.ParseInt(mux.Vars(r)["bookID"], 0, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	currentBook, db, err := model.GetBookByID(bookID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Record Not Found"))
		return
	}
	if updatedBook.Name != "" {
		currentBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		currentBook.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		currentBook.Publication = updatedBook.Publication
	}
	db.Save(currentBook)
	res, _ := json.Marshal(currentBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.ParseInt(mux.Vars(r)["bookID"], 0, 0)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	deletedBook := model.DeleteBook(bookID)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
