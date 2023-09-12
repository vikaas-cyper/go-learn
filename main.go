package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("hello.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template with data (if any)
	data := struct{}{} // You can pass data to your template if needed
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
	router.HandleFunc("/hello", HelloWorld)
	router.HandleFunc("/world", World)

	http.Handle("/", router)

	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)

	fmt.Print("Hello\n")
}
