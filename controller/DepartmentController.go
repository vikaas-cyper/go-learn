package controller

import (
	"html/template"
	"net/http"
)

func DepartmentIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/department/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the template with data (if any)
	data := struct{}{}

	// Render the template with the category data
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
