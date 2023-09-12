package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	tmpl, err := template.ParseFiles("hello.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template with data (if any)
	data := struct {
		Name string
	}{
		Name: vars["name"],
	}

	// Render the template with the category data
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func World(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HelloWorld)
	router.HandleFunc("/hello/{name}", HelloWorld)
	router.HandleFunc("/world", World)

	http.Handle("/", router)

	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)

	fmt.Print("Hello\n")
}
