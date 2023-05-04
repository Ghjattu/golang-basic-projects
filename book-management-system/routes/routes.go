package routes

import (
	"book-management-system/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
