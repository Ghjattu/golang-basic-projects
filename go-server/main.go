package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./ui/form.html")
		if err != nil {
			log.Fatalln(err)
		}
		t.Execute(w, nil)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln("ParseForm() ", err)
		}
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "name: %s\n", name)
		fmt.Fprintf(w, "address: %s\n", address)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting the server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
